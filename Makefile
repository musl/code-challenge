BIN := $(shell basename $(CURDIR))

.PHONY: all clean test

all: test

clean:
	rm -f $(BIN)

$(BIN): 
	go build .

test: $(BIN)
	go test -v .

