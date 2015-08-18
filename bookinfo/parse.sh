#!/bin/sh

cd ..
go build
./flipgo -queues="bookinfocrawl" -concurrency=1 -interval=1
