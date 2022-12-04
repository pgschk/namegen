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
	rand.Seed(time.Now().UnixNano())

	flag.Parse()

	for i := 0; i < *numNames; i++ {
		generateName()
	}

}

/* generateName generates one name by selecting a base name,
 * and adding a prefix or a suffix */
func generateName() string {
	pickedName := getRandomName()
	pickedName = addSuffixOrPrefix(pickedName)
	fmt.Println(pickedName)
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
