#!/usr/bin/env bash

#doesnt work, idk why
# protoc -I internal/messages *.proto --go_out=plugins=grpc:pkg/messages

protoc -I configs/protobuf item_present.proto --go_out=plugins=grpc:./