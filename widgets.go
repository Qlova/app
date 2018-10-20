package app

func ToolBar() *Web {
	return NewToolBar()
}

func NewToolBar() *Web {
	app := New()
	app.SetLayout("flex")
	app.SetSticky()
	return app
}


func Text() *Web {
	app := New()
	app.tag = "p"
	return app
}


func Script(content string) *Web {
	app := New()
	app.tag = "script"
	app.content = []byte(content)
	return app
}

func Header() *Web {
	app := New()
	app.tag = "h1"
	return app
}

func FilePicker(types string) *Web {
	app := New()
	app.tag = "input"
	app.attr = `type="file" accept="`+types+`"`
	return app
}

func TextBox() *Web {
	app := New()
	app.tag = "input"
	return app
}

func TextArea() *Web {
	app := New()
	app.tag = "textarea"
	return app
}

func Button() *Web {
	app := New()
	app.tag = "button"
	return app
}
