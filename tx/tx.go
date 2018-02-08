package tx

import (
	"log"

	"github.com/pkg/errors"
)

type TX struct {
	err     error
	Verbose bool
}

func (tx *TX) log(tpl string, args ...interface{}) {
	if !tx.Verbose {
		return
	}
	log.Printf(tpl+"\n", args...)
}

func (tx *TX) Run(step string, fn func() error) {
	if tx.err != nil {
		return
	}

	tx.log("[%s] started", step)
	tx.err = errors.Wrap(fn(), step)

	if tx.err != nil {
		tx.log("[%s] error: %s", step, tx.err.Error())
	} else {
		tx.log("[%s] finished", step)
	}
}

func (tx *TX) Error() error {
	return tx.err
}
