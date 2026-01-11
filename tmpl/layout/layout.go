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
			Link(Rel("stylesheet"), Href("/assets/output.css")),
		),
		Body(Group(children)),
	)
}

func BaseLayout(title string, children ...Node) Node {
	return RootLayout(title,
		Div(Class("flex min-h-screen flex-col"),
			Header(
				Class("bg-surface fixed top-0 right-0 left-0 z-10 h-20"),
				Div(
					Class("container mx-auto flex h-full items-center"),
					H1(Class("text-3xl font-bold"), A(Href("/"), Text("My Little Blog"))),
				),
			),
			Main(Class("flex-1"),
				Div(Class("container mx-auto pt-20 pb-16"),
					H2(Class("my-8 text-center text-4xl leading-normal font-bold"), Text(title)),
					Group(children),
				),
			),
			Footer(
				Class("bg-surface py-12"),
				Div(Class("container mx-auto text-center"),
					Raw("&copy; 2026 by Karol Moroz. This project is licensed under the BSD-3 license."),
				),
			),
		),
	)
}
