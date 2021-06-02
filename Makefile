# note: call scripts from /scripts

clean:
	rm -rf build/gas

build: clean
	go build -o build/gas cmd/gas/gas.go

build-linux: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/gas cmd/gas/gas.go
