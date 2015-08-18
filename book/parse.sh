#!/bin/sh

cd ..
go build
./flipgo -queues="bookcrawl" -concurrency=1 -interval=1
