#!/bin/bash
echo "## Databases statistics rating"
echo ""
./ghstat -r mongodb/mongo,apache/cassandra,antirez/redis,apache/couchdb,apache/hbase,mysql/mysql-server,postgres/postgres,MariaDB/server,pouchdb/pouchdb,dgraph-io/dgraph,basho/riak -f stats/databases.csv
echo "[Detailed databases statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/c_frameworks.csv)"
echo ""