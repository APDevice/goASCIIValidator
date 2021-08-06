.DEFAULT_GOAL = build
PROG_NAME = asciiValidator
FILES = interface.go
ARGS = 

run:
	go run $(FILES) $(ARGS)

build:
	GOOS=linux GOARCH=amd64 go build -o bin/$(PROG_NAME).linux $(FILES)
	GOOS=windows GOARCH=amd64 go build -o bin/$(PROG_NAME).exe $(FILES)
	GOOS=darwin GOARCH=amd64 go build -o bin/$(PROG_NAME) $(FILES)