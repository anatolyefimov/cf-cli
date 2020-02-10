package main

import (
	"os"

	"github.com/anatolyefimov/cf-cli/cmd"
)

func main() {
	// g := color.New(color.FgGreen)
	// ticker := time.Tick(time.Second)
	// for i := 1; i <= 10; i++ {
	// 	<-ticker
	// 	fmt.Print("\rOn /10")
	// }
	// g.Printf("\rOk\n")

	args := os.Args[1:]
	if args[0] == "login" {
		cmd.Login()
	} else if args[0] == "submit" {
		cmd.Submit(args[1], args[2])
	}

}
