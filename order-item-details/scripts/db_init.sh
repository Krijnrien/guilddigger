#!/usr/bin/env bash


docker run -p 9042:9042 --name cassandra cassandra
docker exec -it cassandra bash
cqlsh
CREATE KEYSPACE IF NOT EXISTS itemdetails WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'};