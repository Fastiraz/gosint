package help

import ("fmt")

func Help() {
	fmt.Println(`
		GOSInt is an osint tool.

			FLAGS:
				-h : Display this message
				-n : Looking for a name
				-p : Phone number
				-l : LinkedIn
				-i : Instagram
				-f : FaceBook
				-g : Google
				-t : Twitter

			EXAMPLE:
				gosint -n name
	`)
}
