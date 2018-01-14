package grifts

import (
	"github.com/gobuffalo/buffalo"

	"github.com/robvdl/vuerecipe/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
