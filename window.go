package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"os/exec"
	"io"
)

func takeInput() (string, error){
	reader := bufio.NewReader(os.Stdin)
	clBuff, err := reader.ReadString('\n')
	if err != nil{
		return "", err
	}
	clBuff = strings.TrimSpace(clBuff)

	return clBuff, nil
}

func cd(dir string) error{
	err := os.Chdir(dir)
	if (err != nil){
		return err
	}
	return nil
}

func runShell(){
	for{
		fmt.Print("Avi's Shell (Type 'exit' to quit)> ")

		command,err := takeInput()
		if err == io.EOF{
			fmt.Println("0 bytes available in input stream")
			break
		}
		if err != nil{
			fmt.Println("Error during Read:",err)
		}
		token := strings.Fields(command)
		if len(token) == 0{
			continue
		}
		if(token[0] =="exit"){
			fmt.Println("Bye-Bye")
			break
		}
		if(token[0] == "cd") {
			fmt.Println("attempting to change dir to:",token[1:])
			err := cd(token[1:][0])
			if(err!=nil){
				fmt.Println("Error while change dir",err)
			}
			continue
		}

		pro := exec.Command(token[0], token[1:]...)
		pro.Stdout = os.Stdout
		pro.Stderr = os.Stderr
		err = pro.Run()
		if err != nil {
                        fmt.Println(err)
                }
	}
}
