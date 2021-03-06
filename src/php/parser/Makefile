PHPFILE=example.php

all: compile fmt build run

fmt:
	find . -type f -iregex '.*\.go' -exec gofmt -l -s -w '{}' +

build:
	go generate ./...
	go build

run:
	./php-parser -d go $(PHPFILE)

test:
	go test ./...

cover:
	go test ./... --cover

bench:
	go test -benchmem -bench=. ./php7

compile: ./php7/php7.go ./scanner/scanner.go fmt
	sed -i '' -e 's/yyErrorVerbose = false/yyErrorVerbose = true/g' ./php7/php7.go
	rm -f y.output

./scanner/scanner.go: ./scanner/scanner.rl
	ragel -Z -G2 -o $@ $<

./php7/php7.go: ./php7/php7.y
	goyacc -o $@ $<

cpu_pprof:
	go test -cpuprofile cpu.pprof -bench=. -benchtime=20s ./php7
	go tool pprof ./php7.test cpu.pprof

mem_pprof:
	go test -memprofile mem.pprof -bench=. -benchtime=20s -benchmem ./php7
	go tool pprof -alloc_objects ./php7.test mem.pprof
