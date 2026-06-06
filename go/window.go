package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"os/exec"
	"log"
)

func runShell(){
	exit:=false
	for exit!=true{
		fmt.Print("myshell> ")
		reader := bufio.NewReader(os.Stdin)
		command,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println("Error:",err)
			exit = true
			continue
		}
		command = strings.TrimSpace(command)
		fmt.Println("You typed:",command)

		token := strings.Fields(command)

		pro := exec.Command(token[0])
		err = pro.Run()
		if err != nil {
                        log.Fatal(err)
                }
		stdOut, err := pro.StdoutPipe() 
                if err != nil{
                        log.Fatal(err)
                }

		if err = pro.Wait(); err!=nil{
			log.Fatal(err)
		}
		fmt.Printf("%s\n",stdOut)
		if(command == "exit"){
			exit = true
		}

	}
	fmt.Println("Out of loop")
}
