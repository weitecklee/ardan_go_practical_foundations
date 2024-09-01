package nlp_test

import (
	"fmt"

	"github.com/weitecklee/nlp"
)

func ExampleTokenize() {
	text := "Who's on first?"
	tokens := nlp.Tokenize(text)
	fmt.Println(tokens)

	// Output:
	// [who on first]
}

func Example_tokenize() {
	text := "Who's on second?"
	tokens := nlp.Tokenize(text)
	fmt.Println(tokens)

	// Output:
	// [who on second]
}

/*

go test
go test -v
go test ./...
go test -v ./...
go test -v -race ./...
go test -v -race -cover ./...

Test discovery:
For every file ending with _test.go, run every function that matches:
- Example[A-Z_].*
- Test[A-Z_].*

Body must include "// Output:" comment
*/
