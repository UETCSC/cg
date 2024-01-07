package cmdutil

import (
	"cg/pkg/global"
	"cg/pkg/tpl"
	"cg/pkg/util"
	"fmt"
	"os"

	"github.com/gookit/color"
	"gopkg.in/yaml.v2"
)

// ConfigSet reads and sets config. yaml from template and writes it to config directory for use in config
func ConfigSet() {
	config := global.Config{}
	_ = yaml.Unmarshal(tpl.Config, &config)
	// Changes to content
	config.Author = util.InputLine("Enter your ID: ")
	config.Contact = util.InputLine("Enter your email address: ")
	config.RegistryUrl = util.SelectOne("Choose the default registry", global.Registry)
	// Writing documents
	writeData, _ := yaml.Marshal(&config)
	UserHomeDir, _ := os.UserHomeDir()
	os.MkdirAll(UserHomeDir+"/.config/cg/", os.ModePerm)
	util.WriteFile(UserHomeDir+"/.config/cg/config.yaml", string(writeData), 0644)
}

// ConfigGet reads the config file and outputs the information to the screen. It is called when the user starts
func ConfigGet() {
	config := global.Config{}
	UserHomeDir, _ := os.UserHomeDir()
	data, err := util.ReadFileByte(UserHomeDir + "/.config/cg/config.yaml")
	// This function will read the config file and check if it has been set.
	if err != nil {
		fmt.Println("Failed to read the config file.", err)
		fmt.Println("Please make sure the config file has been set.")
		return
	}
	_ = yaml.Unmarshal(data, &config)
	Cyan := color.FgCyan.Render
	fmt.Println("Author ID: " + Cyan(config.Author))
	fmt.Println("Author Email: ", Cyan(config.Contact))
	fmt.Println("Default registry: ", Cyan(config.RegistryUrl))
}

// ConfigClean cleans config. yaml and user home dir. It does not check if config is valid
func ConfigClean() {
	UserHomeDir, _ := os.UserHomeDir()
	err := os.Remove(UserHomeDir + "/.config/cg/config.yaml")
	// if err nil print out the error message
	if err != nil {
		fmt.Println("清除错误: ", err)
		return
	}
	fmt.Println("Configuration cleaned.")
}
