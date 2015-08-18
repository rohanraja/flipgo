#!/bin/sh

go build
./flipgo -queues="bookinfocrawl" -concurrency=1
