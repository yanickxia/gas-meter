# note: call scripts from /scripts

.PHONY: test

clean:
	rm -rf build/gas

build: clean
	go build -o build/gas cmd/gas/gas.go

build-linux: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/gas cmd/gas/gas.go

test:
	go test -race -coverprofile=coverage.out ./...
    go tool cover -func coverage.out | tail -n 1 | awk '{ print "Total coverage: " $$3 }'
