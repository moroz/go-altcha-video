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
		Class("space-y-6 my-12"),
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
			El("altcha-widget", Attr("challengeurl", "/api/v1/challenge"), ID("altcha")),
		),
		Button(
			Type("submit"),
			Class("w-full bg-primary text-inherit flex h-12 font-bold text-center items-center justify-center rounded-sm dark:hover:bg-blue-700 cursor-pointer transition-all text-base"),
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
			Span(Class("opacity-80 text-lg"), Text(fmt.Sprintf(" (%v)", len(post.Comments)))),
		),
		If(len(post.Comments) == 0, P(Text("No comments."))),
		Div(Class("space-y-12 mt-8"),
			Map(post.Comments, func(comment *queries.Comment) Node {
				return Article(
					Class("space-y-2"),
					P(
						Strong(Text(comment.Signature)),
						Iff(comment.Website != nil && *comment.Website != "", func() Node {
							return Group{
								Text(" ("),
								A(Class("text-blue-400 underline underline-offset-2 hover:text-blue-500 transition-all"), Href(*comment.Website), Target("_blank"), Rel("noopener noreferrer"), Text("website")),
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
		Script(Type("module"), Src("/assets/app.mjs")),
		Header(
			Class("text-center my-6 opacity-80"),
			P(
				Text("Published on "),
				Time(
					DateTime(post.PublishedAt.Format(time.DateOnly)),
					Text(post.PublishedAt.Format("January 02, 2006")),
				),
			),
		),
		Article(
			Class("prose lg:prose-xl mx-auto dark:prose-invert max-w-full"),
			Raw(renderMarkdown(post.Body)),
		),
		CommentSection(post),
		CommentForm(post),
	)
}
