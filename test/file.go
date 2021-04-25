package test

import (
	_ "embed"
)

var (
	//go:embed reference/plantuml/simple-example.puml
	SimpleExample string
)
