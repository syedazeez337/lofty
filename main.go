package main

import (
	"bufio"
	"fmt"
	"os"
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

		if strmsg == "exit" {
			return
		}

		fmt.Println("=>", strmsg)
	}
}
