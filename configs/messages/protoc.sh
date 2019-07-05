#!/usr/bin/env bash

#doesnt work, idk why
# protoc -I internal/messages *.proto --go_out=plugins=grpc:pkg/messages

protoc -I shared/messages/protobuf item.proto --go_out=plugins=grpc:shared/messages
protoc -I shared/messages/protobuf eventstore.proto --go_out=plugins=grpc:shared/messages
protoc -I shared/messages/protobuf commerce_listings.proto --go_out=plugins=grpc:shared/messages
protoc -I shared/messages/protobuf commerce_prices.proto --go_out=plugins=grpc:shared/messages