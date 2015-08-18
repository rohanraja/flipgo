#!/bin/sh

redis-cli -r 1 RPUSH resque:queue:catcrawl '{"class":"Catcrawl","args":["/books/pr?sid=bks","bks"]}'
