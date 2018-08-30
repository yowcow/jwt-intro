KEYSIZE := 2048

all: key/app.key key/app.key.pub build

key/app.key: key
	openssl genrsa $(KEYSIZE) > $@

key/app.key.pub: key/app.key
	openssl rsa -in $< -pubout > $@

key:
	mkdir -p $@

build:
	go build -o ./main

run:
	./main

clean:
	rm -rf key

.PHONY: all build run clean
