.PHONY: build build_linux clean run start

clean:
	rm -rf target

build:
	mkdir -p target/data
	cp -R data target/
	cp application.yaml target/
	go build -o target/littlebox .

build_linux:
	mkdir -p target/data
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o target/littlebox .

run:
	go run .

start:
	target/littlebox