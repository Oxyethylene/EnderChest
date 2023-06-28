.PHONY: build

clean:
	rm -rf target

build:
	mkdir -p target/data
	cp -R data target/
	cp application.yaml target/
	go build -o target/littlebox .

run:
	target/littlebox