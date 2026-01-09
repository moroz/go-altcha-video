package layout

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func RootLayout(title string, children ...Node) Node {
	return HTML(
		Lang("en-US"),
		Head(
			Meta(Charset("UTF-8")),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
			TitleEl(Text(title+" | Altcha Demo")),
		),
		Body(Group(children)),
	)
}
