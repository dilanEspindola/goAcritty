package utils

import (
	"fmt"
	"os"
)

func OverWriteFile(file []byte) []byte {

	username := GetUserName()

	err := os.WriteFile("C:/Users/"+username+"/AppData/Roaming/alacritty/alacritty.yml", file, 0777)

	CheckError(err)

	return file
}

func ReadFile(theme, fontstyle string) {
	themes := []string{
		"ayu-mirage",
		"one-dark",
		"dracula",
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

	if theme != "none" && fontstyle != "none" {
		for _, value := range themes {
			if value == theme {
				file, errFile := os.ReadFile("./themes/" + theme + ".yml")
				CheckError(errFile)
				newFile := OverWriteFileContent(file, fontstyle, theme)
				if newFile != nil {
					OverWriteFile(newFile)
				} else {
					OverWriteFile(file)
				}
				os.Exit(3)
			}
		}
	}

	fmt.Println("theme does not exist")
}
