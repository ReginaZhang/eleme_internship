package scheduler

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseHost(t *testing.T) {
	a := assert.New(t)

	res, err := parseHosts("example[123:125].com")
	a.Nil(err)
	for _, h := range res {
		fmt.Println(h)
	}
	fmt.Println()

	res, err = parseHosts("example[9:10].com")
	a.Nil(err)
	for _, h := range res {
		fmt.Println(h)
	}
	fmt.Println()

	res, err = parseHosts("example[09:010].com")
	a.Error(err)
	for _, h := range res {
		fmt.Println(h)
	}
	fmt.Println()

	res, err = parseHosts("example[009:10].com")
	a.Error(err)
	for _, h := range res {
		fmt.Println(h)
	}
	fmt.Println()

	res, err = parseHosts("example[09:10].com")
	a.Nil(err)
	for _, h := range res {
		fmt.Println(h)
	}
	fmt.Println()

	res, err = parseHosts("example[009:010].com")
	a.Nil(err)
	for _, h := range res {
		fmt.Println(h)
	}
	fmt.Println()

	res, err = parseHosts("example[a:c].com")
	a.Nil(err)
	for _, h := range res {
		fmt.Println(h)
	}
	fmt.Println()

	res, err = parseHosts("example[ab:bc].com")
	a.Error(err)
	for _, h := range res {
		fmt.Println(h)
	}
	fmt.Println()

	res, err = parseHosts("example[d:a].com")
	a.Error(err)
	for _, h := range res {
		fmt.Println(h)
	}
	fmt.Println()
}
