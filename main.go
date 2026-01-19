package main

import (
	"fmt"
	"gohm/calculate"
	"gohm/cli"
	"gohm/identify"
	"os"
)

const GOHM_VERSION = "0.0.1-beta"
const GOHM_DESCRIPTION = `A zero dependency, very small footprint, electrical engineering CLI utility for calculations and component identification`

func main() {
	c := cli.NewCLI("gohm", GOHM_VERSION, GOHM_DESCRIPTION)

	c.AddCommand(calculate.GetCommand())
	c.AddCommand(identify.GetCommand())

	fmt.Println(c.Run(os.Args))
}
