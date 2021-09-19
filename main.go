/* Program for basic matrix manipulations. */
package main

import(
	"fmt"
	"os"
	"matrix/sum"
)

func
main() {

	utilsMap := map[string] interface{} {
		"sum": sum.Run,
	}

	args := os.Args
	if len(args) < 2 {
		for k, _ := range utilsMap {
			fmt.Println(k) 
		}
		os.Exit(0)
	}
	utilName := args[1]
	args = args[1:]

	if _, ok := utilsMap[utilName] ; !ok {
		fmt.Printf("%s: %s: no such util\n", args[0], utilName)
		os.Exit(1)
	}
	utilsMap[utilName].(func([]string))(args)
}
