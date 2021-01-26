package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// input := []string{}
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	input = append(input, fmt.Sprintf("%s\n", scanner.Text()))
	// }

	input := os.Args[1]

	json, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Stdout.Write(json)
}
