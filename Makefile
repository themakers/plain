

.DEFAULT:
all: install run

install:
	go install

run:
	go build
	./plain
