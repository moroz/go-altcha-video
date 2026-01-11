package blog

import (
	"fmt"
	"time"

	"github.com/moroz/go-altcha-video/db/queries"
	"github.com/moroz/go-altcha-video/tmpl/components"
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
		Class("my-12 space-y-6"),
		H3(Class("text-3xl font-bold"), Text("Leave a comment on this post")),
		components.InputField(&components.InputFieldProps{
			Name:     "signature",
			Required: true,
			Label:    "Signature:",
		}),
		components.InputField(&components.InputFieldProps{
			Name:  "website",
			Label: "Website:",
		}),
		components.TextareaField(&components.TextareaProps{
			Name:     "body",
			Required: true,
			Label:    "Your comment:",
			Rows:     6,
			Class:    "font-mono",
		}),
		Section(
			El("altcha-widget", Attr("challengeurl", "/api/challenge"), Lang("en-US")),
		),
		Button(
			Type("submit"),
			Class("bg-primary hover:bg-primary-hover flex h-12 w-full cursor-pointer items-center justify-center rounded-sm text-center text-base font-bold text-white transition-all dark:text-black"),
			Text("Submit"),
		),
	)
}

func CommentSection(post *types.PostDetailsDto) Node {
	return Section(
		Class("my-12"),
		H3(
			Class("text-3xl font-bold"),
			Text("Comments"),
			Span(Class("text-lg opacity-80"), Text(fmt.Sprintf(" (%v)", len(post.Comments)))),
		),
		If(len(post.Comments) == 0, P(Text("No comments."))),
		Div(Class("mt-8 space-y-12"),
			Map(post.Comments, func(comment *queries.Comment) Node {
				return Article(
					Class("space-y-2"),
					P(
						Strong(Text(comment.Signature)),
						Iff(comment.Website != nil && *comment.Website != "", func() Node {
							return Group{
								Text(" ("),
								A(Class("text-primary hover:text-primary-hover underline underline-offset-2 transition-all"), Href(*comment.Website), Target("_blank"), Rel("noopener noreferrer"), Text("website")),
								Text(")"),
							}
						}),
						Text(", on "), Time(DateTime(comment.CreatedAt.Format(time.RFC3339)), Text(comment.CreatedAt.Format("January 02, 2006, 03:04 PM"))),
					),
					P(Class("text-lg"), Text(comment.Body)),
				)
			}),
		),
	)
}

func Show(post *types.PostDetailsDto) Node {
	return layout.BaseLayout(post.Title,
		Script(Type("module"), Src("https://esm.sh/altcha@2.3.0/es2022/altcha.mjs")),
		Header(
			Class("my-6 text-center opacity-80"),
			P(
				Text("Published on "),
				Time(
					DateTime(post.PublishedAt.Format(time.DateOnly)),
					Text(post.PublishedAt.Format("January 02, 2006")),
				),
			),
		),
		Article(
			Class("prose lg:prose-xl dark:prose-invert mx-auto max-w-full"),
			Raw(renderMarkdown(post.Body)),
		),
		CommentSection(post),
		CommentForm(post),
	)
}
