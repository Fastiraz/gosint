package help

import ("fmt")

func Help() {
	fmt.Println(`
		GOSInt is an OSINT tool.

		FLAGS:
			-h, --help    : Display this message
			-n, --name    : Search for a name
			-p, --phone   : Phone number lookup (dev)
			-g, --github  : GitHub email lookup
			-d, --discord : Discord Token Recovery
			-@, --email   : Find where an email is register (dev)

		EXAMPLE:
			gosint -n name
	`)
}
