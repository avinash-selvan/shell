package main

import (
	"fmt"
	"log"
	"os/exec"
)

func lookPath() {
	path, err := exec.LookPath("git")
	if err != nil {
		log.Fatal("installing git is in your future")
	}
	fmt.Printf("Git is available at %s\n", path)
}

