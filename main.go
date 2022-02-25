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
	x = gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("x"))
	y = gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("y"))
	z, err = gorgonia.Add(x, y)
	if err != nil {
		log.Fatal(err)
	}
	// create a VM to run the program on
	machi