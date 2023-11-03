package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/karamaru-alpha/loopvar"
)

func main() { unitchecker.Main(loopvar.Analyzer) }
