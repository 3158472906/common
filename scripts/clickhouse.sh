#!/usr/bin/env bash



case $1 in
"inck")
  for i in clickhouse-common-static-20.5.3.27-2.x86_64.rpm clickhouse-client-20.5.3.27-2.noarch.rpm clickhouse-server-20.5.3.27-2.noarch.rpm ; do
    rpm -ivh $i --nodeps
  done
;;
"unck")
  for i in clickhouse-client clickhouse-server clickhouse-common-static; do
      rpm -e $i
  done
;;
esac