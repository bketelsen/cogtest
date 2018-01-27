package grifts

import (
	"github.com/bketelsen/cogtest/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
