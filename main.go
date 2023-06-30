package main

import (
	"os"
	"gosint/lib/help"
	"gosint/lib/name" // osint 4 username
	"gosint/lib/phone" // osint 4 phone number
	"gosint/lib/github" // osint 4 github email
	"gosint/lib/discord"
)

func main() {
	argv := os.Args
	for i := 1; i < len(argv)-1; i++ {
		switch argv[i] {
		case "-h", "--help":
			help.Help()
		case "-n", "--name":
			name.Name(argv[i+1])
		case "-p", "--phone":
			phone.Phone(argv[i+1])
		case "-g", "--github":
			github.Email(argv[i+1])
		case "-d", "--discord":
			discord.Token(argv[i+1])
		/*case "-@", "--email":
			email.Email(argv[i+1])*/
		default:
			help.Help()
		}
	}
}
