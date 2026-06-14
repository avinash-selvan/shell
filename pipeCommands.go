package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runPipeCommands(left []string, right []string) error{
	// 2 commands ready for execution
	// A pipe
	// connect them
	// remove the other read and write descriptor
	// execute the commands.

	cmd1 := exec.Command(left[0], left[1:]...)
	cmd2 := exec.Command(right[0], right[1:]...)

	r,w,err := os.Pipe()

	if err!=nil{
		return err
	}

	cmd1.Stdout = w
	cmd2.Stdin = r
	cmd2.Stdout = os.Stdout


	err = cmd1.Start()
	if err!=nil{
		fmt.Println("error from first command")
		return err
	}

	err = cmd2.Start()
	if err != nil{
		fmt.Println("error from second command")
		return err
	}

	err = w.Close()
        if err!= nil{
                fmt.Println("error when closing write")
                return err
        }

        err = r.Close()
        if err!=nil{
                fmt.Println("error when closing read")
                return err
        }

	err = cmd1.Wait()
	if err!=nil{
		fmt.Println("problem with wait 1")
		return err
	}
	err = cmd2.Wait()
	if err!=nil{
		fmt.Println("problem with wait 2")
		return err
	}

	return nil
}

