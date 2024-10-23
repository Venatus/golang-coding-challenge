# Venatus Go Coding Challenge

## Prerequisites
* Docker (tested on v27.2.0)
* Go (tested on go1.22.2 linux/amd64)

## Run clickhouse server in docker
`docker run -d -p 9000:9000 -p 8123:8123 --name venatus-clickhouse-server --ulimit nofile=262144:262144 clickhouse/clickhouse-server`

## Task
* Design a command line tool, that will apply pending schema migrations and store the schema state on the database server.
* On first launch, the program should create a table in clickhouse that stores the migration state.
* Traverse the migrations directory, check SQL files that have not been applied already in alphanumeric order and apply each sequentially.
* The program should not rely on filename alone and should check if the files contents have changed since last applied.
* The tool should support reversion of schema with a command line parameter such as `-rollback 002_alter_column.sql` (see contents of 002_alter_column.sql for rollback). 
* Unit tests designed to verify:
    * Migration applied
    * No operation needed
    * Successfully rolled back

## Tools

Dbeaver [https://dbeaver.io/] is a useful UI to connect to clickhouse. This application will use a native connection (port 9000), but dbeaver will need to use http (port 8123)