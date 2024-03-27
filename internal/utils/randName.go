package utils

import (
	"flag"

	petname "github.com/dustinkirkland/golang-petname"
)

var (
	words     = flag.Int("words", 2, "The number of words in the pet name")
	separator = flag.String("separator", "-", "The separator between words in the pet name")
)

func PetNameGen() string {
	return petname.Generate(*words, *separator)
}
