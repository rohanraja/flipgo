#!/bin/sh

cd ..
go build
./flipgo -queues="catsavedb" -concurrency=10

