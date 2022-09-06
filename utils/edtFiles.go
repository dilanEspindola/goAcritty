package utils

import (
	"fmt"
	"os"
)

func ReadFile(theme, fontstyle string, opacity float64) {
	themes := []string{
		"ayu-mirage",
		"one-dark",
		"dracula",
		"tokyo-night",
		"material-ocean",
		"grubvox-dark",
		"cyberpunk",
	}

	if fontstyle == "none" {
		username := GetUserName()
		file, err := os.ReadFile("C:/Users/" + username + "/AppData/Roaming/alacritty/alacritty.yml")
		CheckError(err)
		newFile := OverWriteFileContentNotFont(file, theme)
		OverWriteFile(newFile)
		os.Exit(3)
	}

	if theme == "none" {
		username := GetUserName()
		file, err := os.ReadFile("C:/Users/" + username + "/AppData/Roaming/alacritty/alacritty.yml")
		CheckError(err)
		newFile := OverWriteFileContentNoTheme(file, fontstyle)
		OverWriteFile(newFile)
		os.Exit(3)
	}

	if opacity == 0.0 {
		username := GetUserName()
		file, err := os.ReadFile("C:/Users/" + username + "/AppData/Roaming/alacritty/alacritty.yml")
		CheckError(err)
		newFile := OverWriteFileContentNoOpacity(file, theme)
		OverWriteFile(newFile)
	}

	if theme == "none" || fontstyle == "none" || opacity != 0.0 {
		username := GetUserName()
		file, err := os.ReadFile("C:/Users/" + username + "/AppData/Roaming/alacritty/alacritty.yml")
		CheckError(err)
		OverWriteFileContentOpacity(file, opacity)
	}
	for _, value := range themes {
		if value == theme {
			file, errFile := os.ReadFile("./themes/" + theme + ".yml")
			CheckError(errFile)
			newFile := OverWriteFileContent(file, fontstyle, theme, opacity)
			if newFile != nil {
				OverWriteFile(newFile)
			} else {
				OverWriteFile(file)
			}
			os.Exit(3)
		}
	}

	fmt.Println("theme does not exist")
}

func OverWriteFile(file []byte) []byte {

	username := GetUserName()

	err := os.WriteFile("C:/Users/"+username+"/AppData/Roaming/alacritty/alacritty.yml", file, 0777)

	CheckError(err)

	return file
}
