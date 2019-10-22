package models

// Layout return information regarding the general display order of different components in view
type Layout struct {
	Alignment    string          `json:"alignment"`
	Order        OrderMap        `json:"order_map"`
	ComponentIDs []ComponentType `json:"components_ids"`
}

// OrderMap holds information regarding the component size and number of components
// TODO elaborate and change later
type OrderMap struct {
	Blocks []int `json:"blocks"`
}
