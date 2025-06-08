package main

import (
	"os"

	"github.com/chetan0402/veripass/internal"
)

func main() {
	dbUrl := os.Getenv("VERIPASS_DATABASE_URL")
	veripass.Run(dbUrl)
}
