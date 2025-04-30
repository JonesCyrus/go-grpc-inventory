package utils

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(v interface{}, label ...string) {
	prettyJSON, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Println("----------------------------------------------")
		fmt.Println("----------------------------------------------")
		fmt.Println("Failed to generate json", err)
		fmt.Println("----------------------------------------------")
		fmt.Println("----------------------------------------------")
	}

	if len(label) > 0 {
		fmt.Println("----------------------------------------------")
		fmt.Println("----------------------------------------------")
		fmt.Print(label[0] + " : ")
		fmt.Printf("%s\n", string(prettyJSON))
		fmt.Println("----------------------------------------------")
		fmt.Println("----------------------------------------------")
		return
	}

	fmt.Println("----------------------------------------------")
	fmt.Println("----------------------------------------------")
	fmt.Printf("%s\n", string(prettyJSON))
	fmt.Println("----------------------------------------------")
	fmt.Println("----------------------------------------------")
}
