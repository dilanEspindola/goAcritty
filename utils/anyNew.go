package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type YmlConfig struct {
	Shell struct {
		Program string `yaml:"program"`
	} `yaml:"shell"`
	Working_directory string `yaml:"working_directory"`
	Font              struct {
		Normal struct {
			Family string `yaml:"family"`
		} `yaml:"normal"`
		Size int `yaml:"size"`
	} `yaml:"font"`
	Window struct {
		Position struct {
			X int `yaml:"x"`
			Y int `yaml:"Y"`
		} `yaml:"position"`
		Dimensions struct {
			Columns int `yaml:"columns"`
			Lines   int `yaml:"lines"`
		} `yaml:"dimensions"`
		Padding struct {
			X int `yaml:"x"`
			Y int `yaml:"y"`
		} `yaml:"padding"`
		Startup_mode string `yaml:"startup_mode"`
		Opacity      int    `yaml:"opacity"`
		Colors       struct {
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
			Draw_bold_text_with_bright_colors string `yaml:"draw_bold_text_with_bright_colors"`
		} `yaml:"colors"`
		Live_config_reload bool `yaml:"live_config_reload"`
		Scrolling          struct {
			History    int `yaml:"history"`
			multiplier int `yaml:"multiplier"`
		} `yaml:"scrolling"`
		Cursor struct {
			Unfocused_hollow bool `yaml:"unfocused_hollow"`
		} `yaml:"colors"`
		Mouse struct {
			Double_click struct {
				Threshold int `yaml:"threshold"`
			} `yaml:"double_click"`
			Triple_click struct {
				Threshold int `yaml:"threshold"`
			} `yaml:"triple"`
			Hide_when_typing bool `yaml:"hide_when_typing"`
		} `yaml:"mouse"`
		Selection struct {
			semantic_escape_chars string `yaml:"semantic_escape_chars"`
			save_to_clipboard     bool   `yaml:"save_to_clipboard"`
		} `yaml:"selection"`
	} `yaml:"window"`
}

func GetSomething() {
	var settings YmlConfig
	file, err := os.ReadFile("../themes/ayu-mirage.yml")
	if err != nil {
		panic(err)
	}

	if err2 := yaml.Unmarshal(file, &settings); err2 != nil {
		panic(err2)
	}

	fmt.Println(settings)
}
