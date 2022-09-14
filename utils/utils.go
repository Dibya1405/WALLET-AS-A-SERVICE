package utils

import (
	"WALLET-AS-A-SERVICE/logger"
	"github.com/google/uuid"
)

// CheckError checking is there any error
func CheckError(e error) {
	logger.GetMyLogger().Warn("Error checking initiated...")
	if e != nil {
		logger.GetMyLogger().Error(e)
		panic(e)
	}
	logger.GetMyLogger().Info("No Error Found")
}

// UUIDGenerator generating an uuid
func UUIDGenerator() string {
	id := uuid.New()
	logger.GetMyLogger().Info("UUID Generated....")
	return id.String()
}
