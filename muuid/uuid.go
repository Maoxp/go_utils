package muuid

import (
	"github.com/google/uuid"
)

//Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services.
//A UUID is a 16 byte (128 bit) array. UUIDs may be used as keys to maps or compared directly.
func GetUUID() string {
	u, _ := uuid.NewRandom()
	return u.String()
}