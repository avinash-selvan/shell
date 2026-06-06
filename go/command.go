package main

import (
	"fmt"
	"log"
	"os/exec"
)

func commandPack() {
	out,err:= exec.Command("git", "--version").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n",out)
}
