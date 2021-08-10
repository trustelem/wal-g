#!/bin/bash

# this script is a wrapper, it calls 'main/pg/wal-g' while setting a known environment

# this script expects you to create a pgenv.sh file which contains environment values to connect
# to an active pgsql instance
if [ ! -e "pgenv.sh" ]; then
	cat >&2 <<EOMESSAGE
*** no file pgenv.sh file found
    please create a 'pgenv.sh' file, which contains values for the following environment variables:
export PGHOST=
export PGPORT=
export PGUSER=
export PGPASSWORD=
export PGDATABASE=
EOMESSAGE
	exit 1
fi

if [ ! -e "main/pg/wal-g" ]; then
	echo "*** 'main/pg/wal-g' doesn't exist" >&2
	echo "    cd to the root of the project, and run 'make pg_build'" >&2
	exit 1
fi

source pgenv.sh

mkdir -p tmp/backups
export WALG_FILE_PREFIX=tmp/backups

main/pg/wal-g "$@"
