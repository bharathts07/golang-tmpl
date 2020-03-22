package utils

import "github.com/rs/xid"

// GetUniqueID returns a globally unique uuid used for storing uniques ID's in the web
func GetUniqueID() string {
	guid := xid.New()

	return guid.String()
}
