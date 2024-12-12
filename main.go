package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/syedazeez337/lofty/tokeniser"
)

func main() {

	fmt.Println("Welcome to lofty compiler...")

	for {
		writer := bufio.NewWriter(os.Stdout)
		input := bufio.NewReader(os.Stdin)

		_, err := writer.WriteString("> ")
		if err != nil {
			fmt.Println("Error", err)
			return
		}

		err = writer.Flush()
		if err != nil {
			fmt.Println("Error", err)
		}

		msg, _, err := input.ReadLine()
		if err != nil {
			fmt.Println("Error", err)
			return
		}

		strmsg := string(msg)

		// exit message
		if strmsg == "exit" {
			fmt.Println("Exiting repl...")
			return
		}

		out := tokeniser.DisplayMsg(strmsg)
		fmt.Println("=>", out)

		/*
		if strmsg == "1 + 2 + 3" {
			out := tokeniser.DisplayMsg(strmsg)
			fmt.Println("=>", out)
		} else {
			fmt.Println("=>", strmsg)
		}
		*/
	}
}
