package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

var count int = 0

func add() int {
	count++
	return count
}

func prettyJson(input string) (string, error) {
	var raw any
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

func main() {
	fmt.Println("Go Web Assembly")

	js.Global().Set("add", js.FuncOf(func(t js.Value, args []js.Value) any {
		return add()
	}))

	js.Global().Set("formatJSON", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "invalid args"
		}
		inputJSON := args[0].String()
		pretty, err := prettyJson(inputJSON)
		if err != nil {
			fmt.Printf("unable to convert to json %s\n", err)
			return err.Error()
		}
		return pretty
	}))

	<-make(chan struct{})
}
