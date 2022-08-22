package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type YmlConfig2 struct {
	Shell struct {
		Program string `yaml:"program"`
	} `yaml:"shell"`
	WorkingDirectory interface{} `yaml:"working_directory"`
	Font             struct {
		Normal struct {
			Family string `yaml:"family"`
		} `yaml:"normal"`
		Size int `yaml:"size"`
	} `yaml:"font"`
	Window struct {
		Position struct {
			X int `yaml:"x"`
		} `yaml:"position"`

		Dimensions struct {
			Columns int `yaml:"columns"`
			Lines   int `yaml:"lines"`
		} `yaml:"dimensions"`

		Padding struct {
			X int `yaml:"x"`
		} `yaml:"padding"`
		Startup_mode string  `yaml:"startup_mode"`
		Opacity      float64 `yaml:"opacity"`
	} `yaml:"window"`
	Colors struct {
		Primary struct {
			Background string `yaml:"background"`
			Foreground string `yaml:"foreground"`
		} `yaml:"primary"`
		Normal struct {
			Black   string `yaml:"black"`
			Red     string `yaml:"red"`
			Green   string `yaml:"green"`
			Yellow  string `yaml:"yellow"`
			Blue    string `yaml:"blue"`
			Magenta string `yaml:"magenta"`
			Cyan    string `yaml:"cyan"`
			White   string `yaml:"white"`
		} `yaml:"normal"`
		Bright struct {
			Black   string `yaml:"black"`
			Red     string `yaml:"red"`
			Green   string `yaml:"green"`
			Yellow  string `yaml:"yellow"`
			Blue    string `yaml:"blue"`
			Magenta string `yaml:"magenta"`
			Cyan    string `yaml:"cyan"`
			White   string `yaml:"white"`
		} `yaml:"bright"`
	} `yaml:"colors"`
	DrawBoldTextWithBrightColors bool `yaml:"draw_bold_text_with_bright_colors"`
	Live_Config_Reload           bool `yaml:"live_config_reload"`
	Scrolling                    struct {
		History    int `yaml:"history"`
		Multiplier int `yaml:"multiplier"`
	} `yaml:"scrolling"`
	Cursor struct {
		Unfocused_Hollow bool `yaml:"unfocused_hollow"`
	} `yaml:"cursor"`
	Mouse struct {
		DoubleClick struct {
			Threshold int `yaml:"threshold"`
		} `yaml:"double_click"`
		TripleClick struct {
			Threshold int `yaml:"threshold"`
		} `yaml:"triple_click"`
		HideWhenTyping bool `yaml:"hide_when_typing"`
	} `yaml:"mouse"`
	Selection struct {
		Semantic_Escape_Chars string `yaml:"semantic_escape_chars"`

		Save_To_Clipboard bool `yaml:"save_to_clipboard"`
	} `yaml:"selection"`
}

func OverWriteFileContent(file []byte, fontstyle string, theme string) []byte {
	var config YmlConfig2

	err := yaml.Unmarshal(file, &config)
	CheckError(err)

	if fontstyle != "none" {
		config.Font.Normal.Family = fontstyle

		newFile, err2 := yaml.Marshal(&config)
		CheckError(err2)

		err3 := ioutil.WriteFile("./themes/"+theme+".yml", newFile, 0777)
		CheckError(err3)
		return newFile
	}
	newFile, err2 := yaml.Marshal(&config)
	CheckError(err2)

	return newFile
}

func OverWriteFileContentNoTheme(file []byte, fontstyle string) []byte {
	var config YmlConfig2

	err := yaml.Unmarshal(file, &config)
	CheckError(err)
	config.Font.Normal.Family = fontstyle

	newFile, err2 := yaml.Marshal(&config)
	CheckError(err2)

	return newFile
}
