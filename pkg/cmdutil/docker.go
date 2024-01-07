package cmdutil

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func exec_cmd(command string) {
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Env = append(cmd.Env, "COMPOSE_DOCKER_CLI_BUILD=1", "DOCKER_BUILDKIT=1")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalln("stderr pipe ", err)
	}
	defer stderr.Close()
	if err := cmd.Start(); err != nil {
		log.Fatalln("start ", err)
	}
	go func() {
		serr := bufio.NewReader(stderr)
		for {
			line, _, err2 := serr.ReadLine()
			if err2 != nil || io.EOF == err2 {
				break
			}
			fmt.Println(string(line))
		}
	}()
	if err := cmd.Wait(); err != nil {
		log.Fatalln("wait ", err)
	}
}

func Auto() {
	Stop()
	Build()
	Run()
}

func Build() {
	exec_cmd("docker compose build")
}

func Run() {
	exec_cmd("docker compose up -d")
}

func Stop() {
	exec_cmd("docker compose down")
}

func Log() {
	exec_cmd("docker compose logs")
}

func Save() {
	fmt.Printf("TODO")
}
