package cmdutil

import (
	"cg/pkg/global"
	"cg/pkg/util"
	"fmt"
	"os"

	"github.com/gookit/color"
	"gopkg.in/yaml.v2"
)

/**
 * @Description: Generated topic template
 * @param baseImageName Title Basic mirror name
 * @param challengeType Subject types
 * @param challengeName Title Title
 * @param language Subtitles
 * @param hasDB Do you need a database?
 */
func Generate(challengeInfo map[string]string) {
	// Flag location confirmed
	challengeInfo["need_flag"] = util.SelectArray("Do you need to handle the flag logic separately? (flag.sh)", []string{"no", "yes"})
	challengeInfo["need_start"] = util.SelectArray("Do you need to handle part of the service startup separately? (start.sh)", []string{"no", "yes"})
	// Selection of subjects
	challengeInfo["level"] = util.SelectArray("Challenge difficulty", global.Level)

	// //DEBUG test output
	// The first is the "Calling Info", which is a reference to the "Calling Info" series.
	// I'm not sure what I'm saying.

	// Second confirmation
	confirm := util.SelectArray("Challenge creation confirmation "+challengeInfo["challenge_name"]+" ?", []string{"confirm", "cancel"})
	if confirm == "cancel" {
		os.Exit(0)
	}

	// Create a template directory of topics
	os.Mkdir(challengeInfo["challenge_name"], 0755)
	os.Chdir(challengeInfo["challenge_name"])

	// Creating folders
	dirTree := []string{
		"environment/src/",
		"environment/files/",
		"writeup/",
	}
	for _, path := range dirTree {
		os.MkdirAll(path, os.ModePerm)
	}
	// Write to the docker file
	GenerateDockerFile(challengeInfo)

	// The following is the list of the official languages of the United States.
	GenerateDockerCompose(challengeInfo)

	// Write in meta.yml
	GenerateMeta(challengeInfo)

	// Write to the database
	GenerateDB(challengeInfo)

	// Write to flag.sh process
	GenerateFlag(challengeInfo)

	// Write to start.sh process
	GenerateStart(challengeInfo)

	// Readme.md is a free app.
	GenerateReadme(challengeInfo)

	// Successful input
	fmt.Println("")
	fmt.Println(challengeInfo["challenge_name"] + " created successfully, please follow the steps below:")
	fmt.Println("1. Initialize Git repository.")
	fmt.Println("2. Edit " + challengeInfo["challenge_name"] + "/meta.yml to modify challenge information")
	fmt.Println("3. Enter " + challengeInfo["challenge_name"] + "/environment for challenge testing")
	fmt.Println("4. Put your writeups in " + challengeInfo["challenge_name"] + "/writeup")
	fmt.Println("5. Push to your remote Git repo")
}

func Wizard() {
	challengeInfo := map[string]string{
		"type":             "", // Subject types
		"language":         "", // Use of language
		"language_version": "", // Language version, HTML title left blank
		"webserver":        "", // Web server, non-web topic left empty
		"db":               "", // Database, no database left empty
		"pwn_arch":         "", // Pwn theme architecture, not Pwn theme left blank
		"pwn_server":       "", // Pwn title server, non-Pwn title left empty
		"need_flag":        "", // If you need flag.sh
		"need_start":       "", // Do you need start.sh?
		"level":            "", // Title ranking
		"base_image_name":  "", // Name of the base mirror
		"base_registry":    "", // The source address of the basic mirror
		"challenge_name":   "", // Title of the title
	}
	// Determine if the default mirror source is set in the configuration
	config := global.Config{}
	UserHomeDir, _ := os.UserHomeDir()
	data, err := util.ReadFileByte(UserHomeDir + "/.config/cg/config.yaml")
	if err != nil {
		fmt.Println("No config file detected. It is recommended to set the registry first.")
		registry := util.SelectOne("Choose the registry you want to use:", global.Registry)
		challengeInfo["base_registry"] = registry + "/"
	} else {
		_ = yaml.Unmarshal(data, &config)
		fmt.Println("Configuration file detected, the registry in the configuration file will be used.")
		fmt.Println("Registry URL：" + color.FgCyan.Render(config.RegistryUrl))
		fmt.Println()
		challengeInfo["base_registry"] = config.RegistryUrl + "/"
	}
	color.Green.Println("If you make the wrong selection, press Ctrl+C to terminate the program and re-execute the wizard.")
	challengeInfo["type"] = util.SelectOne("Select the question type you want to create:", global.ChallengeType)
	challengeInfo["base_image_name"] = challengeInfo["type"]
	switch challengeInfo["type"] {
	case "web":
		challengeInfo = WizardWeb(challengeInfo)
	case "pwn":
		challengeInfo = WizardPwn(challengeInfo)
	case "misc":
		challengeInfo = WizardSocket(challengeInfo)
	}
	color.Cyan.Println("The challenge name should be lowercase and not contain special characters, in the form below:")
	color.Cyan.Println("(year)_(ctf_abbreviation_name)_(category)_(challenge_name)")
	fmt.Println("")
	color.Cyan.Println("Ex 1: 2021 N1CTF 'babysqli' (Web), 2021_n1ctf_web_babysqli.")
	color.Cyan.Println("Ex 2: 2019 SCTF 'babyheap' (Pwn), 2019_sctf_pwn_babyheap.")
	// Keep getting input until there's content
	for {
		challengeInfo["challenge_name"] = util.InputLine("Enter challenge image name:")
		if len(challengeInfo["challenge_name"]) != 0 {
			break
		}
		color.Red.Println("Please re-enter the challenge image name.")
	}
	// Created
	Generate(challengeInfo)
}

func WizardWeb(challengeInfo map[string]string) map[string]string {
	// Judging Language
	challengeInfo["language"] = util.SelectOne("Select the language:", global.Language)
	// Judging the language version
	if challengeInfo["language"] != "html" {
		languageVersion := []string{}
		switch challengeInfo["language"] {
		case "php":
			languageVersion = global.PHPVersion
		case "python":
			languageVersion = global.PythonVersion
		case "nodejs":
			languageVersion = global.NodeJSVersion
		case "java":
			languageVersion = global.JavaVersion
		case "ruby":
			languageVersion = global.RubyVersion
		}
		challengeInfo["language_version"] = util.SelectArray("Select the version:", languageVersion)
	}
	// Judging Web Servers
	switch challengeInfo["language"] {
	case "php", "html":
		challengeInfo["webserver"] = util.SelectOne("Select the PHP webserver", global.PHPWebServer)
	case "java":
		challengeInfo["webserver"] = util.SelectOne("Select the Java webserver", global.JavaServer)
	case "python":
		challengeInfo["webserver"] = util.SelectOne("Select the Python webserver", global.PythonWebServer)
	}
	// Judging the database
	if challengeInfo["language"] != "html" {
		challengeInfo["db"] = util.SelectOne("Do you need a database?", global.DBType)
	}
	// Spelling of the mirror name
	baseImageName := ""
	switch challengeInfo["language"] {
	case "php", "python", "java", "html":
		baseImageName += "_" + challengeInfo["webserver"]
	case "nodejs", "ruby":
		// No need for a web server
	}
	if challengeInfo["db"] != "" {
		baseImageName += "_" + challengeInfo["db"]
	}
	if challengeInfo["language"] != "html" {
		baseImageName += "_" + challengeInfo["language"]
		baseImageName += "_" + challengeInfo["language_version"]
	}
	challengeInfo["base_image_name"] += baseImageName
	return challengeInfo
}

func WizardPwn(challengeInfo map[string]string) map[string]string {
	challengeInfo["pwn_server"] = util.SelectOne("Select your desired startup method", global.PwnServer)
	challengeInfo["pwn_arch"] = util.SelectOne("Select the architecture", global.PwnArch)
	// Spelling of the mirror name
	baseImageName := ""
	switch challengeInfo["pwn_arch"] {
	case "":
		baseImageName += ""
	case "kernel", "qemu":
		baseImageName += "_" + challengeInfo["pwn_arch"]
	}
	baseImageName += "_" + challengeInfo["pwn_server"]
	challengeInfo["base_image_name"] += baseImageName
	return challengeInfo
}

func WizardSocket(challengeInfo map[string]string) map[string]string {
	selectLanguage := global.Language
	delete(selectLanguage, "HTML")
	delete(selectLanguage, "PHP")
	delete(selectLanguage, "Ruby")
	challengeInfo["language"] = util.SelectOne("Select the language", selectLanguage)
	challengeInfo["db"] = util.SelectOne("Do you need a database?", global.DBType)
	// Judging the language version
	languageVersion := []string{}
	switch challengeInfo["language"] {
	case "python":
		languageVersion = global.PythonVersion
	case "nodejs":
		languageVersion = global.NodeJSVersion
	case "java":
		languageVersion = global.JavaVersion
	}
	challengeInfo["language_version"] = util.SelectArray("Select the language version", languageVersion)
	// Spelling of the mirror name
	baseImageName := ""
	if challengeInfo["db"] != "" {
		baseImageName += "_" + challengeInfo["db"]
	}
	baseImageName += "_" + challengeInfo["language"]
	baseImageName += "_" + challengeInfo["language_version"]
	// "Please select the version you want to use", global.SelectVersion[language])
	// The following is a list of the most commonly used database types:
	// This is the first time that the user has ever seen a video.
	// If hasDB!= "No" {
	// 	The main feature of the database is the database.
	// 	The first is the "Street".
	// ♪ I'm not going to be here
	// The first is the image of the user.
	challengeInfo["base_image_name"] += baseImageName
	return challengeInfo
}
