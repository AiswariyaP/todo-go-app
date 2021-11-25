package utility

import (
	"github.com/segmentio/ksuid"
)

// GetGUID will give UUID
func GetGUID() string {
	//TODO: Check GUID version
	return ksuid.New().String()
}
