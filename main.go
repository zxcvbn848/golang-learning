package main

import (
	welcome "first-golang/01-welcome"
	basics "first-golang/02-basics"
	"fmt"
	"os"
)

var mapping = map[string]func(){
	"01-01": welcome.RunWelcome01, // 01-welcome/01-hello-world.go
	"01-04": welcome.RunWelcome04, // 01-welcome/04-playground.go
	"02-01": basics.RunBasics01,   // 02-basics/01-package.go
	"02-02": basics.RunBasics02,   // 02-basics/02-imports.go
	"02-03": basics.RunBasics03,   // 02-basics/03-exported-names.go
	"02-04": basics.RunBasics04,   // 02-basics/04-functions.go
	"02-05": basics.RunBasics05,   // 02-basics/05-functions-continued.go
	"02-06": basics.RunBasics06,   // 02-basics/06-multiple-results.go
	"02-07": basics.RunBasics07,   // 02-basics/07-named-return-values.go
}

func main() {
	if len(os.Args) < 2 {
		printAvailableOptions()
		os.Exit(1)
	}

	arg := os.Args[1]

	if fn, exists := mapping[arg]; exists {
		fn()
	} else {
		fmt.Printf("Unknown argument: %s\n", arg)
		fmt.Println("Available options:")
		for key := range mapping {
			fmt.Printf("  %s\n", key)
		}
		os.Exit(1)
	}
}

func printAvailableOptions() {
	fmt.Println("Usage: go run main.go <01-01|01-04|02-01|02-02|02-03|02-04>")
	fmt.Println("  01-01 - Run Hello, world! (01-welcome/01-hello-world.go)")
	fmt.Println("  01-04 - Run Welcome to the playground! (01-welcome/04-playground.go)")
	fmt.Println("  02-01 - Run My favorite number is (02-basics/01-package.go)")
	fmt.Println("  02-02 - Run Now you have problems. (02-basics/02-imports.go)")
	fmt.Println("  02-03 - Run Exported names. (02-basics/03-exported-names.go)")
	fmt.Println("  02-04 - Run Functions. (02-basics/04-functions.go)")
	fmt.Println("  02-05 - Run Functions continued. (02-basics/05-functions-continued.go)")
	fmt.Println("  02-06 - Run Multiple results. (02-basics/06-multiple-results.go)")
	fmt.Println("  02-07 - Run Named return values. (02-basics/07-named-return-values.go)")
}
