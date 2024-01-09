package global

// The image source configuration
var Registry = map[string]string{
	// "CTFHub":    "registry.cn-hangzhou.aliyuncs.com/",
	"DockerHub": "ctfhub",
}

// Challenge types
var ChallengeType = map[string]string{
	"Web":    "web",
	"Pwn":    "pwn",
	"Socket": "misc",
}

// Languages
var Language = map[string]string{
	"PHP":    "php",
	"HTML":   "html",
	"Python": "python",
	// "NodeJS": "nodejs",
	// "Java":   "java",
	// "Ruby":   "ruby",
}

var PHPVersion = []string{
	"5.6", "7.4", "8.0",
}
var PythonVersion = []string{
	"2.7", "3.11",
}
var NodeJSVersion = []string{
	"12", "14", "16", "18",
}
var JavaVersion = []string{
	"8", "11", "15",
}
var RubyVersion = []string{
	"2.5", "2.6", "2.7",
}

var DBType = map[string]string{
	"Nothing/SQLite": "",
	"MySQL":    "mysql",
	// "MongoDB":  "mongodb",
}

// PHP web server
var PHPWebServer = map[string]string{
	"Apache HTTPd": "httpd",
	"Nginx":        "nginx",
}

// Python web server
var PythonWebServer = map[string]string{
	"gunicorn":   "gunicorn",
	"supervisor": "supervisor",
}

// Java web server
var JavaServer = map[string]string{
	"Tomcat": "tomcat",
}

// Pwn architecture
var PwnArch = map[string]string{
	"x86/x64 Binary": "binary",
	// "x86/x64 Kernel":        "kernel",
	// "arm/arm64/mips/mips64": "qemu",
}

// How to start Pwn
var PwnServer = map[string]string{
	"socat":   "socat",
	"xinet.d": "xinetd",
}

var FileTemplate = map[string]string{
	"flag.sh":  "#!/bin/bash\nflag",
	"start.sh": "#!/bin/bash\nstart",
	"db.sql":   "db",
}

// Difficulty
var Level = []string{"Warm-up", "Easy", "Medium", "Hard"}
