package mod_installer

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/turbot/steampipe/steampipeconfig/version_map"

	"github.com/Masterminds/semver"
	git "github.com/go-git/go-git/v5"
	"github.com/turbot/steampipe/constants"
	"github.com/turbot/steampipe/steampipeconfig/modconfig"
	"github.com/turbot/steampipe/steampipeconfig/parse"
	"github.com/turbot/steampipe/utils"
)

/*
mog get

A user may install a mod with steampipe mod get modname[@version]

version may be:

- Not specified: steampipe mod get aws-cis
	The latest version (highest version tag) will be installed.
	A dependency is added to the requires block specifying the version that was downloaded
- A major version: steampipe mod get aws-cis@3
	The latest release (highest version tag) of the specified major version will be installed.
	A dependency is added to the requires block specifying the version that was downloaded
- A monotonic version tag: steampipe mod get aws-cis@v2.21
	The specified version is downloaded and added as requires dependency.
- A branch name: steampipe mod get aws-cis@staging
	The current version of the specified branch is downloaded.
	The branch dependency is added to the requires list. Note that a branch is considered a distinct "major" release, it is not cached in the registry, and has no minor version.
	Branch versions do not auto-update - you have to run steampipe mod update to get a newer version.
	Branch versioning is meant to simplify development and testing - published mods should ONLY include version tag dependencies, NOT branch dependencies.
- A local file path: steampipe mod get "file:~/my_path/aws-core"
	The mod from the local filesystem is added to the namespace, but nothing is downloaded.
	The local dependency is added to the requires list. Note that a local mod is considered a distinct "major" release, it is not cached in the registry, and has no minor version.
	Local versioning is meant to simplify development and testing - published mods should ONLY include version tag dependencies, NOT local dependencies.


Steampipe Version Dependency
If the installed version of Steampipe does not meet the dependency criteria, the user will be warned and the mod will NOT be installed.

Plugin Dependency
If the mod specifies plugin versions that are not installed, or have no connections, the user will be warned but the mod will be installed. The user should be warned at installation time, and also when starting Steampipe in the workspace.


Detecting conflicts
mod 1 require a@1.0
mod 2 require a@file:/foo

-> how do we detect if the file version satisfied constrainst of a - this is for dev purposes so always pass?

mod 1 require a@1.0
mod 2 require a@<branch>

-> how do we detect if the file version satisfied constraints of a - check branch?

*/

type ModInstaller struct {
	workspaceMod  *modconfig.Mod
	modsPath      string
	workspacePath string

	installData *InstallData

	// should we update dependencies to newer versions if they exist
	updating bool
	// are dependencies being added to the workspace
	AddMods version_map.VersionConstraintMap
	// have specific mods been specified to update
	UpdateMods map[string]bool
}

func NewModInstaller(opts *InstallOpts) (*ModInstaller, error) {
	i := &ModInstaller{
		workspacePath: opts.WorkspacePath,
		modsPath:      constants.WorkspaceModPath(opts.WorkspacePath),
		updating:      opts.Updating,
		AddMods:       make(version_map.VersionConstraintMap),
	}

	// load lock file - ignore errors
	workspaceLock, _ := version_map.LoadWorkspaceLock(i.workspacePath)

	// set the NewModVersions - exclude anything already in the lock
	i.setAddMods(opts.AddMods)

	// create install data
	i.installData = NewInstallData(workspaceLock)

	// now load workspace mod, creating a default if needed
	mod, err := i.loadModfile(i.workspacePath, true)
	if err != nil {
		return nil, err
	}
	i.workspaceMod = mod

	// if we are updating ensure we have a non empty lock file, and that all mods requested to update are in it
	if i.updating {
		if err := i.verifyUpdates(opts.UpdateMods); err != nil {
			return nil, err
		}
	}

	return i, nil
}

// InstallWorkspaceDependencies installs all dependencies of the workspace mod
func (i *ModInstaller) InstallWorkspaceDependencies() error {
	workspaceMod := i.workspaceMod

	// first check our Steampipe version is sufficient
	if err := workspaceMod.Requires.ValidateSteampipeVersion(workspaceMod.Name()); err != nil {
		return err
	}

	// if we are running 'mod get', add the required mods to the workspace mod requires
	if len(i.AddMods) > 0 {
		workspaceMod.AddModDependencies(i.AddMods)
	}

	// if there are no dependencies, we have nothing to do
	if !workspaceMod.HasDependentMods() {
		// there are no dependencies - delete the cache
		i.installData.Lock.Delete(i.workspacePath)
		return nil
	}

	if err := i.installModDependenciesRecursively(workspaceMod.Requires.Mods, workspaceMod); err != nil {
		return err
	}

	// write lock file
	if err := i.installData.Lock.Save(i.workspacePath); err != nil {
		return err
	}

	// if we are running 'mod get', we are now safe to save the mod file
	if len(i.AddMods) > 0 {
		if err := i.workspaceMod.Save(); err != nil {
			return err
		}
	}
	return nil
}

func (i *ModInstaller) installModDependenciesRecursively(mods []*modconfig.ModVersionConstraint, parent *modconfig.Mod) error {
	var errors []error

	for _, requiredModVersion := range mods {
		// get or create the installation data for this mod, adding in this mod version constraint
		availableVersions, err := i.installData.getAvailableModVersions(requiredModVersion.Name)
		if err != nil {
			errors = append(errors, err)
			continue
		}

		// check whether there is already a version which satisfies this mod version
		installedMod, err := i.getInstalledVersionForConstraint(requiredModVersion, availableVersions)
		if err != nil {
			return err
		}

		if installedMod != nil {
			// so we found an existing mod which will satisfy this requirement
			// update the workspace lock
			i.installData.addExisting(requiredModVersion.Name, installedMod.Version, parent)
			log.Printf("[TRACE] not installing %s with version constraint %s as version %s is already installed", requiredModVersion.Name, requiredModVersion.Constraint.Original, installedMod.Version)
		} else {
			// so we ARE installing

			// get a resolved mod ref that satisfies the version constraints
			resolvedRef, err := i.getModRefSatisfyingConstraints(requiredModVersion, availableVersions)
			if err != nil {
				errors = append(errors, err)
				continue
			}
			// install the mod
			installedMod, err = i.install(resolvedRef, parent)
			if err != nil {
				errors = append(errors, err)
			}
			if installedMod == nil {
				// this is unexpected but just ignore
				log.Printf("[TRACE] dependency %s does not define a mod definition - so there are no child dependencies to install", resolvedRef.Name)
				continue
			}
		}

		// to get here we have the dependency mod - either we installed it or it was already installed
		// recursively install its dependencies
		if err := i.installModDependenciesRecursively(installedMod.Requires.Mods, installedMod); err != nil {
			return err
		}

	}

	return utils.CombineErrorsWithPrefix(fmt.Sprintf("%d %s failed to install", len(errors), utils.Pluralize("dependency", len(errors))), errors...)
}

// check whether there is a mod version installed that satisfies the version constraint (and update requirements)
func (i *ModInstaller) getInstalledVersionForConstraint(requiredModVersion *modconfig.ModVersionConstraint, availableVersions []*semver.Version) (*modconfig.Mod, error) {
	// TODO
	//log.Printf("[TRACE] getInstalledVersionForConstraint required version %v", requiredModVersion)
	//
	//installedVersion := i.installData.AllInstalled.GetVersionSatisfyingRequirement(requiredModVersion)
	//if installedVersion == nil {
	//	log.Printf("[TRACE] no version of %s installed which satisfies version constrain %s", requiredModVersion.Name, requiredModVersion.Constraint.Original)
	//	return nil, nil
	//}
	//log.Printf("[TRACE] found installed version %s@%s", requiredModVersion.Name, installedVersion.Original())
	//
	//// so there IS a version installed which satisfies the constraint.
	//// if we are updating, see if there is a newer verison
	//if i.shouldUpdate(requiredModVersion.Name) {
	//	// so we should update if there is a newer version - check if there is
	//	newerModVersionFound, err := i.newerModVersionFound(requiredModVersion, installedVersion, availableVersions)
	//	if err != nil {
	//		return nil, err
	//	}
	//	if newerModVersionFound {
	//		// there is a newer version so we will NOT use the installed version - return nil
	//		return nil, nil
	//	}
	//}
	//
	//// so we resolved an installed version which will satisfy
	//// load the mod
	//modPath := filepath.Join(i.modsPath, modVersionFullName(requiredModVersion.Name, installedVersion))
	//installedMod, err := i.loadModfile(modPath, false)
	//if err != nil {
	//	return nil, err
	//}

	return nil, nil
}

// determine whether there is a newer mod version avoilable which satisfies the dependency version constraint
func (i *ModInstaller) newerModVersionFound(requiredVersion *modconfig.ModVersionConstraint, currentVersion *semver.Version, availableVersions []*semver.Version) (bool, error) {
	latestVersion, err := i.getModRefSatisfyingConstraints(requiredVersion, availableVersions)
	if err != nil {
		return false, err
	}
	if latestVersion.Version.GreaterThan(currentVersion) {
		return true, nil
	}
	return false, nil
}

// install a mod
func (i *ModInstaller) install(dependency *ResolvedModRef, parent *modconfig.Mod) (_ *modconfig.Mod, err error) {
	defer func() {
		if err == nil {
			i.installData.onModInstalled(dependency, parent)
		}
	}()

	// TODO HANDLE FILE SOURCE
	destPath := i.getDependencyDestPath(dependency.FullName())

	// if the target path exists, this is a bug - we should never try to install over an existing directory
	if _, err := os.Stat(destPath); !os.IsNotExist(err) {
		return nil, fmt.Errorf("mod %s is already installed", dependency.FullName())
	}

	if err := i.installFromGit(dependency, destPath); err != nil {
		return nil, err
	}

	// now load the installed mod and return it
	return i.loadModfile(destPath, false)

}

func (i *ModInstaller) getDependencyDestPath(dependencyFullName string) string {
	return filepath.Join(i.modsPath, dependencyFullName)
}

func (i *ModInstaller) installFromGit(dependency *ResolvedModRef, installPath string) error {
	// ensure mod directory exists - create if necessary
	if err := os.MkdirAll(i.modsPath, os.ModePerm); err != nil {
		return err
	}

	// NOTE: we need to check existing installed mods

	// get the mod from git
	gitUrl := getGitUrl(dependency.Name)
	_, err := git.PlainClone(installPath,
		false,
		&git.CloneOptions{
			URL: gitUrl,
			//Progress:      os.Stdout,
			ReferenceName: dependency.GitReference,
			Depth:         1,
			SingleBranch:  true,
		})

	return err
}

func (i *ModInstaller) loadModfile(modPath string, createDefault bool) (*modconfig.Mod, error) {
	if !parse.ModfileExists(modPath) {
		if createDefault {
			return modconfig.CreateDefaultMod(i.workspacePath), nil
		}
		return nil, nil
	}
	return parse.ParseModDefinition(modPath)
}

// get the most recent available mod version which satisfies the version constraint
func (i *ModInstaller) getModRefSatisfyingConstraints(modVersion *modconfig.ModVersionConstraint, availableVersions []*semver.Version) (*ResolvedModRef, error) {
	// find a version which satisfies the version constraint
	var version = getVersionSatisfyingConstraint(modVersion.Constraint, availableVersions)
	if version == nil {
		return nil, fmt.Errorf("no version of %s found satisfying verison constraint: %s", modVersion.Name, modVersion.Constraint.Original)
	}

	return NewResolvedModRef(modVersion, version)
}

func (i *ModInstaller) setAddMods(getMods version_map.VersionConstraintMap) {
	// TODO
	//for name, contraint := range getMods{
	//	!if i.installData.Lock.ContainsModConstraint()
	//}
}
