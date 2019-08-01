package main

import (
	"github.com/QOSGroup/kepler/cmd"
	_ "github.com/QOSGroup/kepler/docs"
)

// @title Kepler API
// @version v0.6.0
// @description  Kepler server.
// @BasePath /
// @Host 127.0.0.1:8080
func main() {
	cmd.Execute()
}
