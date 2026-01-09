package blog

import (
	"fmt"

	"github.com/moroz/go-altcha-video/tmpl/layout"
	"github.com/moroz/go-altcha-video/types"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func renderMarkdown(body string) string {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.Footnotes | parser.NoEmptyLineBeforeBlock | parser.Autolink
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(body))

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer))
}

func CommentForm(post *types.PostDetailsDto) Node {
	return Form(
		Action(fmt.Sprintf("/blog/%s/comments", post.Slug)),
		Method("POST"),
		Class("space-y-6 my-12"),
		H3(Class("text-3xl font-bold"), Text("Leave a comment on this post")),
		Div(Class("grid gap-2"),
			Label(Class("font-bold"), Text("Signature: "), Span(Class("text-red-700 dark:text-red-300"), Text("*"))),
			Input(Name("signature"), Class("h-12 border-border border-solid border bg-surface")),
		),
		Div(Class("grid gap-2"),
			Label(Class("font-bold"), Text("Your comment: "), Span(Class("text-red-700 dark:text-red-300"), Text("*"))),
			Textarea(Name("body"), Class("bg-surface border-border border-solid border font-mono"), Rows("6")),
		),
	)
}

func Show(post *types.PostDetailsDto) Node {
	return layout.BaseLayout(post.Title,
		Article(Class("prose lg:prose-xl mx-auto dark:prose-invert max-w-full"), Raw(renderMarkdown(post.Body))),
		CommentForm(post),
	)
}
