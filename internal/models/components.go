package models

// ComponentType is the identifier to identify the type of component present in the enclosing struct
type ComponentType int

const (
	Banner    ComponentType = 1
	MenuBar   ComponentType = 2
	TestBlock ComponentType = 3
)

// Component holds complete information of a single displayable component
type Component struct {
	Type ComponentType `json:"type"`
	Data *interface{}  `json:"data"`
}

// Header defines information to be displayed at the top of any view
type Header struct {
	Title     string      `json:"title"`
	Component []Component `json:"components"`
}
