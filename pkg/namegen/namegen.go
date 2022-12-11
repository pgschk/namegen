package namegen

import (
	"math/rand"
	"strings"
	"time"
)

/* generateName generates one name by selecting a base name,
 * and adding a prefix or a suffix */
func GenerateName() string {
	pickedName := GetRandomName()
	pickedName = addSuffixOrPrefix(pickedName)
	return pickedName
}

/* getRandomName returns a random name from a list */
func GetRandomName() string {
	rand.Seed(time.Now().UnixNano()) // initialize RNG seed
	randomIndex := rand.Intn(len(Names))
	pickedName := Names[randomIndex]
	return pickedName
}

/* getRandomPrefix returns a random prefix from a list */
func GetRandomPrefix() string {
	randomIndex := rand.Intn(len(Prefixes))
	prefix := Prefixes[randomIndex]
	return prefix
}

/* getRandomName returns a random suffix from a list */
func GetRandomSuffix() string {
	randomIndex := rand.Intn(len(Suffixes))
	suffix := Suffixes[randomIndex]
	return suffix
}

/* addSuffixOrPrefix randomly adds a prefix or a suffix to a given name */
func addSuffixOrPrefix(pickedName string) string {
	prefixOrSuffix := rand.Intn(2)
	if prefixOrSuffix > 0 {
		// add a prefix
		pickedName = GetRandomPrefix() + strings.ToLower(pickedName)
	} else {
		// add a suffix
		pickedName = pickedName + strings.ToLower(GetRandomSuffix())
	}
	return pickedName
}
