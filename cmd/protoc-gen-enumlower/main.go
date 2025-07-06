package main

import (
	"github.com/thejezzi/protoc-gen-enumlower/pkg/enumlower"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		for _, file := range plugin.Files {
			if !file.Generate {
				continue
			}
			enumlower.Generate(plugin, file)
		}
		return nil
	})
}
