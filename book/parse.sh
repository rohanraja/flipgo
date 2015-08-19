#!/bin/sh

cd ..
go build
./flipgo -queues="bookcrawl" -concurrency=60 -interval=1
