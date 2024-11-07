package wingfoilnews

import "time"

type Link struct {
	Id      int
	Title   string
	URI     string
	Tags    []string
	Created time.Time
}
