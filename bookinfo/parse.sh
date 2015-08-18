#!/bin/sh

cd ..
go build
./flipgo -queues="bookinfocrawl" -concurrency=100 -interval=1
