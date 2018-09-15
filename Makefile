BIN := $(shell basename $(CURDIR))

.PHONY: all clean test

all: test

clean:
	rm -f $(BIN)

$(BIN): 
	go build .

test: $(BIN)
	go test -v .

pg2009.txt:
	curl -sO http://www.gutenberg.org/cache/epub/2009/pg2009.txt

big.txt: pg2009.txt
	for i in `seq 1000` ; do cat pg2009.txt >> big.txt ; done

run: $(BIN) big.txt
	time ./$(BIN) big.txt	

