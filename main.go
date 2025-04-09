package main

import (
	"bytes"
	genContainer "dockgen/gen-container"
	"fmt"
	"golang.org/x/term"
	"io"
	"log"
	"os"
	"sync"
)

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	if len(os.Args) != 3 || os.Args[1] != "debug" {
		fmt.Println("Usage: dockgen debug <gen-container-id>")
		return
	}
	containerID := os.Args[2]
	fmt.Println(containerID)
	cli := genContainer.Client()
	execContainer, err := genContainer.CreateExecContainer(cli, "./Dockerfile", genContainer.CreateExecOptions{
		ImageName:     "hello:0.2",
		ContainerName: "hello1",
		Cmd:           "/bin/sh",
	})

	if err != nil {
		return
	}
	go genContainer.CopyIO(os.Stdout, execContainer.Reader)
	var w WaitGroup
	w.Do(func() {
		buf := make([]byte, 1)
		var n int
		for {

			if n, err = os.Stdin.Read(buf); err != nil {
				break
			}
			if _, err = io.Copy(execContainer.Conn, bytes.NewReader(buf[:n])); err != nil {
				break
			}

		}
	})
	w.Wait()
}

type WaitGroup struct {
	sync.WaitGroup
}

func (wait *WaitGroup) Do(do func()) {
	wait.Add(1)
	go func() {
		defer wait.Done()
		do()
	}()
}
