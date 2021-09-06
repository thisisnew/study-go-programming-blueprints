package main

import (
	"log"
	"os"
	"os/exec"
)

var cmdChain = []*exec.Cmd{
	exec.Command("C:/Users/dhkim/go/src/study-go-programming-blueprints/ch4/domainfinder/lib/synonyms"),
	exec.Command("C:/Users/dhkim/go/src/study-go-programming-blueprints/ch4/domainfinder/lib/sprinkle"),
	exec.Command("C:/Users/dhkim/go/src/study-go-programming-blueprints/ch4/domainfinder/lib/coolify"),
	exec.Command("C:/Users/dhkim/go/src/study-go-programming-blueprints/ch4/domainfinder/lib/domainify"),
	exec.Command("C:/Users/dhkim/go/src/study-go-programming-blueprints/ch4/domainfinder/lib/available"),
}

func main() {

	cmdChain[0].Stdin = os.Stdin
	cmdChain[len(cmdChain)-1].Stdout = os.Stdout

	for i := 0; i < len(cmdChain)-1; i++ {
		thisCmd := cmdChain[i]
		nextCmd := cmdChain[i+1]
		stdout, err := thisCmd.StdoutPipe()
		if err != nil {
			log.Fatalln(err)
		}
		nextCmd.Stdin = stdout
	}

	for _, cmd := range cmdChain {
		if err := cmd.Start(); err != nil {
			log.Fatalln(err)
		} else {
			defer cmd.Process.Kill()
		}
	}

	for _, cmd := range cmdChain {
		if err := cmd.Wait(); err != nil {
			log.Fatalln(err)
		}
	}

}
