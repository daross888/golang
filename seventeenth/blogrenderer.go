package seventeenth

import (
	"fmt"
	"io"
)

type Post struct {
	Title       string
	Description string
	Body        string
	Tags        []string
}

func Render(w io.Writer, post Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>", post.Title)

	return err
}
