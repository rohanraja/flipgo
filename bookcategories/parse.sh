#!/bin/sh

cd ..
go build
./flipgo -queues="catcrawl" -concurrency=60
