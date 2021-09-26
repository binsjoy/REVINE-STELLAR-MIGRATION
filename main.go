package main

import (
	"fmt"
	"log"

	"github.com/chewxy/gorgonia"
	"gorgonia.org/tensor"
)

func main() {
	g := gorgonia.NewGraph()
	var x, y, z *gorgonia.Node
	var err error
	// define the expression
	x = gorgonia