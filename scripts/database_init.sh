#!/usr/bin/env bash

#NOTE, this isnt meant to run in bash but rather copy pasted since the SHELL wont be instated through a script but only through cmd

echo "SSH into roach1, starting sql shell"
docker exec -it roach1 ./cockroach sql --insecure
CREATE DATABASE IF NOT EXISTS ordersdb;
use ordersdb;
