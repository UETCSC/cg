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

func ConfigSet() {
	config := global.Config{}
	_ = yaml.Unmarshal(tpl.Config, &config)
	config.Author = util.InputLine("Enter your ID: ")
	config.Contact = util.InputLine("Enter your email address: ")
	config.RegistryUrl = util.SelectOne("Choose the default registry", global.Registry)
	writeData, _ := yaml.Marshal(&config)
	UserHomeDir, _ := os.UserHomeDir()
	os.MkdirAll(UserHomeDir+"/.config/cg/", os.ModePerm)
	util.WriteFile(UserHomeDir+"/.config/cg/config.yaml", string(writeData), 0644)
}

func ConfigGet() {
	config := global.Config{}
	UserHomeDir, _ := os.UserHomeDir()
	data, err := util.ReadFileByte(UserHomeDir + "/.config/cg/config.yaml")
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

func ConfigClean() {
	UserHomeDir, _ := os.UserHomeDir()
	err := os.Remove(UserHomeDir + "/.config/cg/config.yaml")
	if err != nil {
		fmt.Println("Clean error: ", err)
		return
	}
	fmt.Println("Configuration cleaned.")
}
