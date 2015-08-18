#!/bin/sh

redis-cli -r 1 RPUSH resque:queue:bookcrawl '{"class":"Bookcrawl","args":["/lc/pr/pv1/spotList1/spot1/productList?sid=bks&pincode=721302&filterNone=true&start=41&ajax=true","bks"]}'
