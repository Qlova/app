package textarea

import (
	"qlova.org/seed"
	"qlova.org/seed/css"
	"qlova.org/seed/html/attr"
	"qlova.org/seed/js"
	"qlova.org/seed/script"
	"qlova.org/seed/state"

	"qlova.org/seed/s/html/textarea"
)

//New returns a new textbox widget.
func New(options ...seed.Option) seed.Seed {
	return textarea.New(css.SetResize(css.None), seed.Options(options))
}

//Var returns text with a variable text argument.
func Var(text state.String, options ...seed.Option) seed.Seed {
	if text.Null() {
		return New(options...)
	}
	return New(seed.NewOption(func(c seed.Seed) {
		c.With(script.On("input", func(q script.Ctx) {
			text.Set(js.String{js.NewValue(script.Scope(c, q).Element() + `.value`)})(q)
		}), text.SetValue())
	}), seed.Options(options))
}

//SetPlaceholder sets the placeholder of the textbox.
func SetPlaceholder(placeholder string) seed.Option {
	return attr.Set("placeholder", placeholder)
}

//SetReadOnly sets the textbox to be readonly.
func SetReadOnly() seed.Option {
	return attr.Set("readonly", "")
}

//Focus focuses the textbox.
func Focus(c seed.Seed) script.Script {
	return func(q script.Ctx) {
		q(script.Element(c).Run(`focus`))
	}
}