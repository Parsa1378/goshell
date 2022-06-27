package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

// improvment needed
// func error_handling()

func execInput(input string) error {
	//cleaning input
	input = strings.TrimSuffix(input, "\n")

	//Split the input to separate the command and the arguments
	args := strings.Split(input, " ")

	//Pass the program and the arguments separately
	cmd := exec.Command(args[0], args[1:]...)

	switch args[0] {

	case "cd":

		if len(args[0]) < 2 {
			return errors.New("Path Not Specified")
		}

		os.Chdir(args[1])

	case "exit":
		os.Exit(0)
	}

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func main() {

	// get the current user

	me, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Printf("%s -> ", me.Username)
		//get input
		input, err := reader.ReadString('\n')

		//check for errors
		if err != nil {
			log.Fatal(err)
		}

		if err := execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}
