osx:
	go build -o bin/flip_osx
linux:
	GOOS=linux GOARCH=amd64 go build -o bin/flip_linux
windows:
	GOOS=windows GOARCH=386 go build -o bin/flipgo.exe
windows64:
	GOOS=windows GOARCH=amd64 go build -o bin/flipgo64.exe


clean:
	rm bin/*
