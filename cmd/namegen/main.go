package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	/* parse command line arguments */
	numNames := flag.Int("n", 15, "Number of names to generate")
	slowMode := flag.Bool("slowmode", false, "If set the names will be output at a constant rate defined by --slowmode-rate")
	slowModeRate := flag.Int64("slowmode-rate", 1, "Defines the rate at which names are output in slowmode. Defaults to 1/second")

	flag.Parse()
	rand.Seed(time.Now().UnixNano()) // initialize RNG seed

	// generate numNames names
	for i := 0; i < *numNames; i++ {
		// generate and output a name
		pickedName := generateName()
		fmt.Println(pickedName)
		if *slowMode {
			// if slow mode is enabled, delay next name by slowModeRate seconds
			time.Sleep(time.Duration(*slowModeRate) * time.Second)
		}
	}
}

/* generateName generates one name by selecting a base name,
 * and adding a prefix or a suffix */
func generateName() string {
	pickedName := getRandomName()
	pickedName = addSuffixOrPrefix(pickedName)
	return pickedName
}

/* getRandomName returns a random name from a list */
func getRandomName() string {
	randomIndex := rand.Intn(len(Names))
	pickedName := Names[randomIndex]
	return pickedName
}

/* getRandomPrefix returns a random prefix from a list */
func getRandomPrefix() string {
	randomIndex := rand.Intn(len(Prefixes))
	prefix := Prefixes[randomIndex]
	return prefix
}

/* getRandomName returns a random suffix from a list */
func getRandomSuffix() string {
	randomIndex := rand.Intn(len(Suffixes))
	suffix := Suffixes[randomIndex]
	return suffix
}

/* addSuffixOrPrefix randomly adds a prefix or a suffix to a given name */
func addSuffixOrPrefix(pickedName string) string {
	prefixOrSuffix := rand.Intn(2)
	if prefixOrSuffix > 0 {
		// add a prefix
		pickedName = getRandomPrefix() + strings.ToLower(pickedName)
	} else {
		// add a suffix
		pickedName = pickedName + strings.ToLower(getRandomSuffix())
	}
	return pickedName
}
