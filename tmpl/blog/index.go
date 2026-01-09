package blog

import (
	"fmt"

	"github.com/moroz/go-altcha-video/tmpl/layout"
	"github.com/moroz/go-altcha-video/types"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Index(posts []*types.PostListDto) Node {
	return layout.BaseLayout("Blog",
		Map(posts, func(post *types.PostListDto) Node {
			return Article(
				H3(
					A(
						Href(fmt.Sprintf("/blog/%s", post.Slug)),
						Text(post.Title),
					),
				),
			)
		}),
	)
}
