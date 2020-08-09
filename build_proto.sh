#!/bin/bash

protoc \
--proto_path=. \
--go-grpc_out=. \
--go_out=plugins:. \
--go_opt=paths=source_relative ./sdi.proto