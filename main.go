/*
Oats -- One-at-a-Time To-do's
*/
package main

import (
	"github.com/yamlinson/oats/cmd"
	"github.com/yamlinson/oats/internal/data"
)

func main() {
	err := data.MkDataDir()
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}
