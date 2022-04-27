package main

import (
	"gitlab.uuzu.com/war/pbtool/generator"
	"gitlab.uuzu.com/war/pbtool/plugin/clone"
)

func main() {
	generator.Main(&clone.Plugin{})
}
