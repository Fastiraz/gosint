package main

import (
	"fmt"
	"os"
	"gosint/lib/help"
	"gosint/lib/name" // osint 4 username
	"gosint/lib/phone" // osint 4 phone number
)

var printf = fmt.Printf

func main() {
	argv := os.Args
	for i:=1; i<len(argv); i++ {
		if argv[i] == "-h" {
			help.Help()
		} else if argv[i] == "-n" {
			//name := argv[i+1]
			name.Name(argv[i+1])
		} else if argv[i] == "-p" {
			phone.Phone(argv[i+1])
		} else {
			help.Help()
		}
	}
}

