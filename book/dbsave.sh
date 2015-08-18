#!/bin/sh

cd ..
go build
./flipgo -queues="booksavedb" -concurrency=1 -interval=1

