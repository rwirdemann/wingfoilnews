package wingfoilnews

import (
	"fmt"
	"github.com/xeonx/timeago"
	"net/url"
	"strings"
	"time"
)

type Link struct {
	Id      int
	Title   string
	URI     string
	Tags    []string
	Created time.Time
}

func (l Link) Source() string {
	u, err := url.Parse(strings.Trim(l.URI, " "))
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s://%s", u.Scheme, u.Host)
}

func (l Link) Ago() string {
	return timeago.German.Format(l.Created)
}
