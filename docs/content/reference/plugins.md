---
linkTitle: Plugins
title: How to write plugins for gqlgen
description: Use plugins to customize code generation and integrate with other libraries
menu: { main: { parent: "reference", weight: 10 } }
---

Plugins provide a way to hook into the gqlgen code generation lifecycle. In order to use anything other than the
default plugins you will need to create your own entrypoint:

## Using a plugin

To use a plugin during code generation, you need to create a new entry point. Create `generate.go` in the same folder as `resolver.go` with the following code:

```go
// go:build ignore

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/john-markham/gqlgen/api"
	"github.com/john-markham/gqlgen/codegen/config"
	"github.com/john-markham/gqlgen/plugin/stubgen"
)

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}


	err = api.Generate(cfg,
		api.AddPlugin(yourplugin.New()), // This is the magic line
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}

```

In `resolver.go`, add `//go:generate go run generate.go`. Now you can run `go generate ./...` instead of `go run github.com/john-markham/gqlgen generate` to generate the code.

## Writing a plugin

There are currently only two hooks:

- MutateConfig: Allows a plugin to mutate the config before codegen starts. This allows plugins to add
  custom directives, define types, and implement resolvers. see
  [modelgen](https://github.com/john-markham/gqlgen/tree/master/plugin/modelgen) for an example
- GenerateCode: Allows a plugin to generate a new output file, see
  [stubgen](https://github.com/john-markham/gqlgen/tree/master/plugin/stubgen) for an example

Take a look at [plugin.go](https://github.com/john-markham/gqlgen/blob/master/plugin/plugin.go) for the full list of
available hooks. These are likely to change with each release.
