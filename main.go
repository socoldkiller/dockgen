package main

import (
	genContainer "dockgen/gen-container"
	"fmt"
	"golang.org/x/term"
	"log"
	"os"
)

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	_, err = os.Create("bash_history")
	if err != nil {

		return
	}
	defer os.Remove("bash_history")
	attachContainer, err := genContainer.CreateAttachContainer(genContainer.Client(), genContainer.CreateBuildAndContainerOptions{
		Tag:           "debugger:0.2",
		ContainerName: "",
		Cmd:           "/bin/sh",
		Tty:           true,
		Mounts: map[string]string{
			"bash_history": "/root/.bash_history",
		},
		Env: []string{
			"HISTFILE=/root/.bash_history",
			"HISTSIZE=10000",
			"HISTFILESIZE=20000",
		},
	})

	if err != nil {
		log.Println(err)
	}
	go genContainer.CopyIO(attachContainer.Conn, os.Stdin)
	genContainer.CopyIO(os.Stdout, attachContainer.Reader)
	term.Restore(int(os.Stdin.Fd()), oldState)
	f, _ := os.ReadFile("bash_history")
	defer fmt.Printf("%s", f)
}
