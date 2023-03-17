.PHONY: build

clean:
	rm -rf target
	mkdir -p target

build:
	mkdir -p target/data
	cp -R data target/
	go build -o target/littlebox .

run: build
	target/littlebox