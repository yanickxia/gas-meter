# note: call scripts from /scripts

clean:
	rm -rf build/gas

build: clean
	go build -o build/gas cmd/gas/gas.go
