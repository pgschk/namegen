package main

import (
	"flag"
	"fmt"
	"time"

	NameGen "github.com/pgschk/namegen/pkg/namegen"
)

func main() {
	/* parse command line arguments */
	numNames := flag.Int("n", 15, "Number of names to generate")
	slowMode := flag.Bool("slowmode", false, "If set the names will be output at a constant rate defined by --slowmode-rate")
	slowModeRate := flag.Int64("slowmode-rate", 1, "Defines the rate at which names are output in slowmode. Defaults to 1/second")

	flag.Parse()

	// generate numNames names
	for i := 0; i < *numNames; i++ {
		// generate and output a name
		pickedName := NameGen.GenerateName()
		fmt.Println(pickedName)
		if *slowMode {
			// if slow mode is enabled, delay next name by slowModeRate seconds
			time.Sleep(time.Duration(*slowModeRate) * time.Second)
		}
	}
}
