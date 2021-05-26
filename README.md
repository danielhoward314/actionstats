## To Use as a Library Class

Since the assignment calls for a "small library class", I've structured the code for use as an import. The library is written in Go and assumes the following prerequisites:

- your machine's architecture and OS meet the minimum requirements documented [here](https://github.com/golang/go/wiki/MinimumRequirements)
- you've installed the Go programming language following the instructions [here](https://golang.org/doc/install)
- your installed Go version is 1.16+

With the prerequisites covered, follow these steps for a sample use of the imported library.

(Steps 1-5 and 6 are to be run in the terminal. These commands assume bash shell and *nix system; run equivalent for your OS)
1. `mkdir <yourdirname>`
2. `cd <yourdirname>`
3. `go mod init github.com/<yournamespace>/<yourdirname>` (substitute gitlab/bitbucket/etc. as needed; change templated namespace and project name)
4. `go get github.com/danielhoward314/actionstats@v0.1.5`
5. `touch main.go`
6. Open `main.go` in a code editor and paste the code snippet below:
7. Back in terminal, run `main.go`

```
package main

import (
	"fmt"
	"sync"

	"github.com/danielhoward314/actionstats"
)

func main() {
	store := actionstats.NewActionStore()
	testCases := []string{
		`{"action":"jump", "time":100}`,
		`{"action":"run", "time":75}`,
		`{"action":"jump", "time":200}`,
	}
	c := make(chan string)
	go example(store, testCases, c)
	statsJson := <-c
	fmt.Println(statsJson)
}

func example(store *actionstats.ActionStore, tests []string, c chan string) {
	var wg sync.WaitGroup
	wg.Add(len(tests))
	for _, s := range tests {
		current := s
		go func() {
			defer wg.Done()
			store.AddAction(current)
		}()
	}
	wg.Wait()
	c <- store.GetStats()
}
```

This example shows how to use the library concurrently. It prints the expected JSON to the console. A real use case may return `statsJson`, store it in a variable and use it elsewhere, or slot this library in as handlers for an http server exposing endpoints like `POST /actions` and `GET /stats`.

## To Run Test Cases

Running the test cases assumes all of the same prerequisites outlined above. 

(These are terminal commands. The same assumptions as above apply.)
1. `git clone https://github.com/danielhoward314/actionstats.git`
2. `cd actionstats`
3. `go test` or, for verbose output, `go test -v`