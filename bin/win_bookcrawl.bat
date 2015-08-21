set http_proxy=http://10.3.100.207:8080
flipgo.exe -queues="bookcrawl" -concurrency=20 -interval=5 -redis="10.109.11.216" -uri=redis://10.109.11.216:6379/
