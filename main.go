package main

import (
	"flag"

	"github.com/dilanEspindola/goAcritty/utils"
)

func main() {

	theme := flag.String("theme", "none", "type a theme")

	flag.Parse()

	utils.ReadFile(*theme)
}
