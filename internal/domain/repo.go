package domain

import "net/url"

type Repository struct {
	Type     string
	Title    string
	User     string
	Openlink url.URL
}

func (r *Repository) isValid() bool {
	return !(r.Title == "" || r.User == "" || r.Type == "")
}
