# Challenge!
This code satisfies the requirements for a code challenge. It reads text
and prints the top 100 most frequently occurring three word sequences.

## Requirements
This code requires go 1.11 for module support.

## Testing
To test this code, run `make`. You should see the output from successful code compilation and tests.

## Large Example Run
To run this example against a large file, run `make run`. This should
take a little while. When this was written, it would download a ~3 MB
text file, copy it into a 1.2 GB text file. This should print out the
list of three word sequences and also the total processing time taken.

## Usage
You can run this program and give a list of one or more files. You can also run it by piping data in over standard input.
Examples:

		# list of files
		./code-challenge <file> [file] ...
		
		# reading from stdin
		echo big.txt | code-challenge
		
You should see the top-100 list of three word sequences printed out, one
per line.

