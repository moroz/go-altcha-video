package blog

import (
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

func Show(post *types.PostDetailsDto) Node {
	return layout.BaseLayout(post.Title,
		Article(Class("prose lg:prose-xl mx-auto dark:prose-invert max-w-full"), Raw(renderMarkdown(post.Body))),
	)
}
