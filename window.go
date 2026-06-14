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

func backgroundProcess(cmd *exec.Cmd) (int, error){
	err := cmd.Start()
	if err!= nil{
		return 0, err
	}

	go func() {
		err := cmd.Wait()
		if err != nil {
			fmt.Println("Process exited with an error:",err)
		}
	}()
	return cmd.Process.Pid,nil
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

		pipe := false

		for i,tok := range token{
			if tok == "|"{
				pipe = true
				left := token[:i]
				right := token [i+1:]
				err := runPipeCommands(left, right)
				if err!=nil{
					fmt.Println("Error executing the commands:",err)
				}
			}
		}
		if pipe{
			continue
		}
		if(token[0] =="exit"){
			fmt.Println("Bye-Bye")
			break
		} else if(token[0] == "cd") {
			fmt.Println("attempting to change dir to:",token[1:])
			err := cd(token[1:][0])
			if(err!=nil){
				fmt.Println("Error while change dir",err)
			}
			continue
		} else if(token[len(token)-1] == "&"){
			backPro := exec.Command(token[0], token[1:len(token)-1]...)
			backPID, err := backgroundProcess(backPro)
			if(err != nil){
				fmt.Println("Error during start:",err)
				continue
			}
			fmt.Println(backPID)
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
