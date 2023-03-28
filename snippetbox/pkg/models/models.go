package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("model: no matching record found")

type Snippet struct {
	Id      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
