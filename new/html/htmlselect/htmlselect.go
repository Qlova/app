package htmlselect

import (
	"qlova.org/seed"
	"qlova.org/seed/use/html"
)

//New returns a new HTML select element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("select"),

		seed.Options(options),
	)
}
