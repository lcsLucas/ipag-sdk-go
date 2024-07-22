package utils

// Pretty prints the given data in a pretty format
// This is useful for debugging purposes
// Example usage:
//  utils.PrintPretty(data)

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrintPretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}
