package main

import (
	"os"
	"gosint/lib/name" // osint 4 username (sherlock)
	"gosint/lib/phone" // osint 4 phone number (phoneinfoga)
	"gosint/lib/github" // osint 4 github email
	"gosint/lib/discord"
	"gosint/lib/email" // osint 4 github email (holehe)
	"gosint/lib/card" // credit card information recovery
)

func main() {
	argv := os.Args
	for i := 1; i < len(argv); i++ {
		switch argv[i] {
			case "-n", "--name":
				name.Name(argv[i+1])
				os.Exit(0)
			case "-p", "--phone":
				phone.Phone(argv[i+1])
				os.Exit(0)
			case "-g", "--github":
				github.Email(argv[i+1])
				os.Exit(0)
			case "-d", "--discord":
				discord.Token(argv[i+1])
				os.Exit(0)
			case "-@", "--email":
				email.Email(argv[i+1])
				os.Exit(0)
			case "-c", "--card":
				card.Card(argv[i+1])
				os.Exit(0)
			default:
				os.Stdout.WriteString(`
					GOSInt is an OSINT tool.

					FLAGS:
						-h, --help    : Display this message
						-n, --name    : Search for a name
						-p, --phone   : Phone number lookup (dev)
						-g, --github  : GitHub email lookup
						-d, --discord : Discord Token Recovery
						-@, --email   : Find where an email is registered (dev)

					EXAMPLE:
						gosint -n name
				`)
		}
	}
}
