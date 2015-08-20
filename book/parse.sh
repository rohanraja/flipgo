#!/bin/sh

cd ..
go build
./flipgo -queues="bookcrawl" -concurrency=90 -interval=1
