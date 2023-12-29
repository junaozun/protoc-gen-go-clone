package main

import (
	"github.com/junaozun/protoc-gen-go-clone/generator"
	"github.com/junaozun/protoc-gen-go-clone/plugin/clone"
)

func main() {
	generator.Main(&clone.Plugin{})
}
