package constants

const PostgresqlConfContent = `
# -----------------------------
# PostgreSQL configuration file
# -----------------------------
#
# DO NOT EDIT THIS FILE!
# It is overwritten each time Steampipe starts.
#
# In the rare case that postgres.conf customization is required, modifications
# or additions should be placed in the 'postgres.conf.d' folder as a config
# include file. For example: 'postgres.conf.d/01-custom-settings.conf'.
# See https://www.postgresql.org/docs/current/config-setting.html#CONFIG-INCLUDES
#

# First, use Steampipe's default settings for Postgres.
include = 'steampipe.conf'

# Second, allow users to customize Postgres settings with custom '.conf' files
# created in the 'postgresql.conf.d' directory. Use with care, these settings
# overwrite any 'steampipe.conf' settings above.
include_dir = 'postgresql.conf.d'
`

const SteampipeConfContent = `
# ------------------------------------------
# Steampipe's default Postgres configuration
# ------------------------------------------
#
# DO NOT EDIT THIS FILE!
# It is overwritten each time Steampipe starts.
#
# In the rare case that postgres.conf customization is required, modifications
# or additions should be placed in the 'postgres.conf.d' folder as a config
# include file. For example: 'postgres.conf.d/01-custom-settings.conf'.
# See https://www.postgresql.org/docs/current/config-setting.html#CONFIG-INCLUDES
#

# Steampipe is run in many different systems and regions, so use UTC for all
# timestamps by default - both in SQL responses and log entries.
timezone=UTC
log_timezone=UTC

# Make the database log consistent with our plugin logs in both name and daily
# rotation frequency. These will appear in '~/.steampipe/logs' and are cleared
# after 7 days by the Steampipe CLI.
log_filename='database-%Y-%m-%d.log'

# Postgres log messages sent to stderr should be redirected to the log file.
logging_collector=on

# Connection logging is fast, low volume and helpful to troubleshoot issues
# around plugin startup or failure.
log_connections=on
log_disconnections=on

# Logging of slow queries (> 5 secs) is helpful when reviewing environments or
# troubleshooting with users.
log_min_duration_statement=5000

# Increasing the locks per transaction helps PostgreSQL to not
# run out of available memory when working with large plugins
# or aggregators with a large number of sub connections (or both)
max_locks_per_transaction = 2048 

shared_buffers=256MB
max_connections=1000
`
