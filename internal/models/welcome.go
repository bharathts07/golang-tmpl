package models

// Welcome defines the fields returned to client
type Welcome struct {
	Message string `json:"message"`
	Home string    `json:"home_link"`
}

// Home contains the layout and an ordered list of  components to be displayed
type Home struct {
	Layout HomeLayout `json:"home_layout"`
	Components []Component `json:"components"`
}