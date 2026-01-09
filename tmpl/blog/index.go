package blog

import (
	"fmt"
	"time"

	"github.com/moroz/go-altcha-video/tmpl/layout"
	"github.com/moroz/go-altcha-video/types"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Index(posts []*types.PostListDto) Node {
	return layout.BaseLayout("Blog",
		Map(posts, func(post *types.PostListDto) Node {
			return Article(
				Class("space-y-1"),
				H3(
					A(
						Class("text-3xl font-bold hover:underline underline-offset-2 dark:hover:text-blue-500 transition-all"),
						Href(fmt.Sprintf("/blog/%s", post.Slug)),
						Text(post.Title),
					),
				),
				P(
					Text("By "),
					Strong(Text(post.Author)),
					Text(", on "),
					Time(
						Class("font-bold"),
						DateTime(post.PublishedAt.Format(time.DateOnly)),
						Text(post.PublishedAt.Format("January 02, 2006")),
					),
				),
				P(
					A(
						Class("font-bold hover:underline underline-offset-2 text-blue-500"),
						Href(fmt.Sprintf("/blog/%s", post.Slug)),
						Text("Read more..."),
					),
				),
			)
		}),
	)
}
