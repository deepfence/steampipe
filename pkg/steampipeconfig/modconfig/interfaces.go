package modconfig

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
)

type (
	// MappableResource must be implemented by resources which can be created
	// directly from a content file (e.g. sql, markdown)
	MappableResource interface {
		// InitialiseFromFile creates a mappable resource from a file path
		// It returns the resource, and the raw file data
		InitialiseFromFile(modPath, filePath string) (MappableResource, []byte, error)
		Name() string
		GetUnqualifiedName() string
		GetMetadata() *ResourceMetadata
		SetMetadata(*ResourceMetadata)
		GetDeclRange() *hcl.Range
	}

	// HclResource must be implemented by resources defined in HCL
	HclResource interface {
		Name() string
		GetUnqualifiedName() string
		CtyValue() (cty.Value, error)
		OnDecoded(*hcl.Block, ModResourcesProvider) hcl.Diagnostics
		AddReference(ref *ResourceReference)
		GetReferences() []*ResourceReference
		GetDeclRange() *hcl.Range
	}
	// ModTreeItem must be implemented by elements of the mod resource hierarchy
	// i.e. Control, Benchmark, Dashboard
	ModTreeItem interface {
		Name() string
		GetUnqualifiedName() string
		AddParent(ModTreeItem) error
		GetParents() []ModTreeItem
		GetChildren() []ModTreeItem
		GetTitle() string
		GetDescription() string
		GetTags() map[string]string
		// GetPaths returns an array resource paths
		GetPaths() []NodePath
		SetPaths()
		GetMod() *Mod
	}

	// ResourceWithMetadata must be implemented by resources which supports reflection metadata
	ResourceWithMetadata interface {
		Name() string
		GetMetadata() *ResourceMetadata
		SetMetadata(metadata *ResourceMetadata)
		SetAnonymous(block *hcl.Block)
		IsAnonymous() bool
	}
	// QueryProvider must be implemented by resources which supports prepared statements, i.e. Control and Query
	QueryProvider interface {
		HclResource
		GetArgs() *QueryArgs
		GetParams() []*ParamDef
		GetSQL() *string
		GetQuery() *Query
		SetArgs(*QueryArgs)
		SetParams([]*ParamDef)
		GetPreparedStatementName() string
		GetPreparedStatementExecuteSQL(*QueryArgs) (*ResolvedQuery, error)
		// implemented by QueryProviderBase
		AddRuntimeDependencies([]*RuntimeDependency)
		GetRuntimeDependencies() map[string]*RuntimeDependency
		RequiresExecution(QueryProvider) bool
		VerifyQuery(QueryProvider) error
		MergeParentArgs(QueryProvider, QueryProvider) (diags hcl.Diagnostics)
	}
	// DashboardLeafNode must be implemented by resources may be a leaf node in the dashboard execution tree
	DashboardLeafNode interface {
		HclResource
		GetTitle() string
		GetDisplay() string
		GetDescription() string
		GetDocumentation() string
		GetType() string
		GetTags() map[string]string
		GetWidth() int
		GetPaths() []NodePath
		GetMetadata() *ResourceMetadata
		GetChildren() []ModTreeItem
	}
	ModResourcesProvider interface {
		GetResourceMaps() *ModResources
	}
	// EdgeAndNodeProvider must be implemented by any dashboard leaf node which supports edges and nodes
	// (DashboardGraph, DashboardFlow, DashboardHierarchy)
	EdgeAndNodeProvider interface {
		QueryProvider
		GetEdges() DashboardEdgeList
		SetEdges(DashboardEdgeList)
		GetNodes() DashboardNodeList
		SetNodes(DashboardNodeList)
	}
)
