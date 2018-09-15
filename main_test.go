package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TestText = `A colander (or cullender) is a bowl-shaped kitchen
                utensil with holes in it used for draining food such
                as pasta or rice. A colander or also used to rinse
                vegetables. The perforated nature of the colander allows
                liquid to drain through while retaining the solids
                inside. It is sometimes also called a pasta strainer
                or kitchen sieve.  Conventionally, colanders are also
                of a light metal, such as aluminium or thinly rolled
                stainless steel. Colanders are also made of plastic,
                silicone, ceramic, and enamelware I like turtles i like
                turtles I Like Turtles I LIKE TURTLES I like Turtles`
)

func TestTopNTriples(t *testing.T) {
	reader := io.Reader(strings.NewReader(TestText))
	triples, err := TopNTriples(10, []io.Reader{reader})

	assert.Nil(t, err)
	assert.Equal(t, 10, len(triples))
	assert.Equal(t, `i like turtles`, triples[0].Value)
	assert.Equal(t, 5, triples[0].Count)
}
