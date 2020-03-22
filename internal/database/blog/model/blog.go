package model

import "time"

// BlogPreview holds enough information for a short preview for display in main page
type BlogPreview struct {
	ID      int
	Title   string
	Author  string
	Summary string
	Tags    []string
	Photo   string
	Created time.Time
	Updated time.Time
}

// BlogItem holds complete information for display in the entire content
type BlogItem struct {
	ID      int
	Title   string
	Author  string
	Body    string
	Tags    []string
	Photo   []string
	Created time.Time
	Updated time.Time
}
