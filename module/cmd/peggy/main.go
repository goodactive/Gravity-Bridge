package main

import (
	"os"

	"github.com/althea-net/peggy/module/cmd/peggy/cmd"
	"github.com/cosmos/cosmos-sdk/server"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)
		default:
			os.Exit(1)
		}
	}
}
