#!/bin/sh

set -e

host="$1"
port="$2"
shift 2
cmd="$@"

until PGPASSWORD=$DB_PASSWORD psql -h "$host" -p "$port" -U "postgres" -c '\q'; do
  >&2 echo "PostgreSQL is unavailable - sleeping"
  sleep 1
done

>&2 echo "PostgreSQL is ready - executing command"
exec $cmd