package main

import (
	"log"
	"os"

	"github.com/GeraldLoeffler/go-greetings/greetings"
)

func main() {
	args := os.Args
	switch len(args) {
	case 0, 1: // no name
		log.Fatal("Missing command-line arg")
	case 2: // one name
		n := args[1]

		g, err := greetings.Greet(n)
		if err != nil {
			log.Fatal(err)
		}

		log.Print("Greeting ", n, " with: ", g.Greeting)
	default: // many names
		ns := args[1:]

		gs, err := greetings.GreetAll(ns)
		if err != nil {
			log.Fatal(err)
		}

		// could loop over range of map, but want to keep order of names
		for _, n := range ns {
			g := gs[n]
			log.Print("Greeting ", n, " with: ", g.Greeting)
		}
	}
}
