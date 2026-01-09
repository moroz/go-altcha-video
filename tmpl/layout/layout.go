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
		Div(Class("flex flex-col min-h-screen"),
			Header(
				Class("bg-slate-100 h-20 dark:bg-slate-700 fixed top-0 left-0 right-0 z-10"),
				Div(
					Class("container mx-auto flex items-center h-full"),
					H1(Class("text-3xl font-bold"), A(Href("/"), Text("My Little Blog"))),
				),
			),
			Main(Class("flex-1"),
				Div(Class("container mx-auto pb-16 pt-20"),
					H2(Class("text-center text-4xl font-bold leading-normal my-8"), Text(title)),
					Group(children),
				),
			),
			Footer(
				Class("bg-slate-100 py-12 dark:bg-slate-700"),
				Div(Class("container mx-auto text-center"),
					Raw("&copy; 2026 by Karol Moroz. This project is licensed under the BSD-3 license."),
				),
			),
		),
	)
}
