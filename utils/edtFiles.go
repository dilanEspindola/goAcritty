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

func ReadFile(theme string) {
	themes := []string{
		"ayu-mirage",
		"ayu-dark",
		"one-dark",
		"dracula",
	}

	for _, value := range themes {
		if value == theme {
			file, errFile := os.ReadFile("./themes/" + theme + ".yml")
			CheckError(errFile)
			OverWriteFile(file)
			os.Exit(3)
		}
	}
	fmt.Println("theme does not exist")
}
