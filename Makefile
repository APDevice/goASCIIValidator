.DEFAULT_GOAL = build
PROG_NAME = asciiValidator
FILES = interface.go
ARGS = 

run:
	go run $(FILES) $(ARGS)

build:
	go build -o $(PROG_NAME) $(FILES)