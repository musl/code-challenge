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
                as pasta or rice. A colander is also used to rinse
                vegetables. The perforated nature of the colander allows
                liquid to drain through while retaining the solids
                inside. It is sometimes also called a pasta strainer
                or kitchen sieve.  Conventionally, colanders are made
                of a light metal, such as aluminium or thinly rolled
                stainless steel. Colanders are also made of plastic,
                silicone, ceramic, and enamelware`
)

func TestTopNTriples(t *testing.T) {
	reader := io.Reader(strings.NewReader(TestText))
	triples, err := TopNTriples(100, reader)

	assert.Nil(t, err)
	assert.Equal(t, 10, len(triples))
}
