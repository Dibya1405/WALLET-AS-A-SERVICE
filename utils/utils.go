package utils

import "github.com/google/uuid"

// CheckError checking is there any error
func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

// UUIDGenerator generating an uuid
func UUIDGenerator() string {
	id := uuid.New()
	return id.String()
}
