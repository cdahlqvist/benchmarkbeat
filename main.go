package main

import (
	"os"

	"github.com/cdahlqvist/benchmarkbeat/cmd"

	_ "github.com/cdahlqvist/benchmarkbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
