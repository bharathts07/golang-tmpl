package model

import "time"

// BlogPreview holds enough information for a short preview for display in main page
type BlogPreview struct {
	ID      string    `firestore:"id" json:"id"`
	Title   string    `firestore:"title" json:"title"`
	Author  string    `firestore:"author" json:"author"`
	Summary string    `firestore:"summary" json:"summary"`
	Tags    []string  `firestore:"tags" json:"tags"`
	Photo   string    `firestore:"photo" json:"photo"`
	Created time.Time `firestore:"created" json:"created"`
	Updated time.Time `firestore:"updated" json:"updated"`
}

// BlogItem holds complete information for display in the entire content
type BlogItem struct {
	ID      string    `firestore:"id" json:"id" `
	Title   string    `firestore:"title" json:"title"`
	Author  string    `firestore:"author" json:"author"`
	Body    string    `firestore:"body" json:"body"`
	Tags    []string  `firestore:"tags" json:"tags"`
	Photo   []string  `firestore:"photo" json:"photo"`
	Created time.Time `firestore:"created" json:"created"`
	Updated time.Time `firestore:"updated" json:"updated"`
}
