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
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	cmd.Execute()
}
