package domain

import "net/url"

const (
	open = iota
	merged
	review
)

type PR struct {
	Type     string
	Title    string
	User     string
	Openlink url.URL
}
