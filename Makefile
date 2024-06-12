.PHONY: clean
clean:
	rm -rf target

.PHONY: build
build:
	mkdir -p target/data
	cp -R data target/
	cp application.yaml target/
	go build -o target/ender_chest .

.PHONY: build_linux
build_linux:
	mkdir -p target/data
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o target/ender_chest .

.PHONY: run
run:
	go run .

.PHONY: start
start:
	target/ender_chest