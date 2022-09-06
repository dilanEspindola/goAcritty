package main

import (
	"flag"

	"github.com/dilanEspindola/goAcritty/utils"
)

func main() {
	theme := flag.String("theme", "none", "type a theme")
	fontStyle := flag.String("f", "none", "type a font")
	opacity := flag.Float64("o", 0.0, "type an opacity")

	flag.Parse()

	utils.ReadFile(*theme, *fontStyle, *opacity)

}
