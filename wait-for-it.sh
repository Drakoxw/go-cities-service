#!/bin/sh
# wait-for-it.sh

host="$1"
shift
cmd="$@"

until nc -z -v -w30 $host 3306
do
  echo "Waiting for MySQL connection..."
  sleep 1
done

>&2 echo "MySQL is up - executing command"
exec $cmd