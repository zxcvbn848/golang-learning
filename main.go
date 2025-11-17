package main

import (
	welcome "first-golang/01-welcome"
	basics "first-golang/02-basics"
	flowControl "first-golang/03-flow-control"
	moreTypes "first-golang/04-more-types"
	"fmt"
	"os"
)

type option struct {
	code        string
	description string
	filePath    string
	runFunc     func()
}

var options = []option{
	{"01-01", "Run Hello, world!", "01-welcome/01-hello-world.go", welcome.RunWelcome01},
	{"01-04", "Run Welcome to the playground!", "01-welcome/04-playground.go", welcome.RunWelcome04},
	{"02-01", "Run My favorite number is", "02-basics/01-package.go", basics.RunBasics01},
	{"02-02", "Run Now you have problems.", "02-basics/02-imports.go", basics.RunBasics02},
	{"02-03", "Run Exported names.", "02-basics/03-exported-names.go", basics.RunBasics03},
	{"02-04", "Run Functions.", "02-basics/04-functions.go", basics.RunBasics04},
	{"02-05", "Run Functions continued.", "02-basics/05-functions-continued.go", basics.RunBasics05},
	{"02-06", "Run Multiple results.", "02-basics/06-multiple-results.go", basics.RunBasics06},
	{"02-07", "Run Named return values.", "02-basics/07-named-return-values.go", basics.RunBasics07},
	{"02-08", "Run Variables.", "02-basics/08-variables.go", basics.RunBasics08},
	{"02-09", "Run Variables with initializers.", "02-basics/09-variables-with-initializers.go", basics.RunBasics09},
	{"02-10", "Run Short variable declarations.", "02-basics/10-short-variable-declarations.go", basics.RunBasics10},
	{"02-11", "Run Basic types.", "02-basics/11-basic-types.go", basics.RunBasics11},
	{"02-12", "Run Zero values.", "02-basics/12-zero-values.go", basics.RunBasics12},
	{"02-13", "Run Type conversions.", "02-basics/13-type-conversions.go", basics.RunBasics13},
	{"02-14", "Run Type inference.", "02-basics/14-type-inference.go", basics.RunBasics14},
	{"02-15", "Run Constants.", "02-basics/15-constants.go", basics.RunBasics15},
	{"02-16", "Run Numeric Constants.", "02-basics/16-numeric-constants.go", basics.RunBasics16},
	{"03-01", "Run For.", "03-flow-control/01-for.go", flowControl.RunFlowControl01},
	{"03-02", "Run For continued.", "03-flow-control/02-for-continued.go", flowControl.RunFlowControl02},
	// 03-03 同 03-02
	{"03-04", "Run Forever.", "03-flow-control/04-forever.go", flowControl.RunFlowControl04},
	{"03-05", "Run If.", "03-flow-control/05-if.go", flowControl.RunFlowControl05},
	{"03-06", "Run If with a short statement.", "03-flow-control/06-if-with-a-short-statement.go", flowControl.RunFlowControl06},
	{"03-07", "Run If and else.", "03-flow-control/07-if-and-else.go", flowControl.RunFlowControl07},
	{"03-08", "Run Exercise: Loops and Functions.", "03-flow-control/08-exercise-loops-and-functions.go", flowControl.RunFlowControl08},
	{"03-09", "Run Switch.", "03-flow-control/09-switch.go", flowControl.RunFlowControl09},
	{"03-10", "Run Switch evaluation order.", "03-flow-control/10-switch-evaluation-order.go", flowControl.RunFlowControl10},
	{"03-11", "Run Switch with no condition.", "03-flow-control/11-switch-with-no-condition.go", flowControl.RunFlowControl11},
	{"03-12", "Run Defer.", "03-flow-control/12-defer.go", flowControl.RunFlowControl12},
	{"03-13", "Run Stacking defers.", "03-flow-control/13-stacking-defers.go", flowControl.RunFlowControl13},
	{"04-01", "Run Pointers.", "04-more-types/01-pointers.go", moreTypes.RunMoreTypes01},
	{"04-02", "Run Structs.", "04-more-types/02-structs.go", moreTypes.RunMoreTypes02},
	{"04-03", "Run Struct fields.", "04-more-types/03-struct-fields.go", moreTypes.RunMoreTypes03},
	{"04-04", "Run Pointers to structs.", "04-more-types/04-pointers-to-structs.go", moreTypes.RunMoreTypes04},
	{"04-05", "Run Struct literals.", "04-more-types/05-struct-literals.go", moreTypes.RunMoreTypes05},
	{"04-06", "Run Arrays.", "04-more-types/06-arrays.go", moreTypes.RunMoreTypes06},
	{"04-07", "Run Slices.", "04-more-types/07-slices.go", moreTypes.RunMoreTypes07},
	{"04-08", "Run Slices are like references to arrays.", "04-more-types/08-slices-are-like-references-to-arrays.go", moreTypes.RunMoreTypes08},
}

var mapping = func() map[string]func() {
	m := make(map[string]func())
	for _, opt := range options {
		m[opt.code] = opt.runFunc
	}
	return m
}()

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
		for _, opt := range options {
			fmt.Printf(" %s,", opt.code)
		}
		fmt.Println()
		os.Exit(1)
	}
}

func printAvailableOptions() {
	// Build usage string from options
	usageCodes := ""
	for i, opt := range options {
		if i > 0 {
			usageCodes += "|"
		}
		usageCodes += opt.code
	}
	fmt.Printf("Usage: go run main.go <%s>\n", usageCodes)

	// Print all options from the options slice
	for _, opt := range options {
		fmt.Printf("  %s - %s (%s)\n", opt.code, opt.description, opt.filePath)
	}
}
