#!/bin/sh

cd ..
go build
./flipgo -queues="bookinfosavedb" -concurrency=1 -interval=1

