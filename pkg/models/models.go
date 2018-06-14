package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   sql.NullString
	Content sql.NullString
	Created time.Time
	Expires time.Time
}

type Snippets *[]Snippet
