package utils

import (
	"WALLET-AS-A-SERVICE/dto"
	"WALLET-AS-A-SERVICE/logger"
	"github.com/google/uuid"
)

// CheckError checking is there any error
func CheckError(e error) bool {
	logger.GetMyLogger().Warn("Error checking initiated...")
	if e != nil {
		logger.GetMyLogger().Error(e)
		//panic(e)
		return true
	}
	logger.GetMyLogger().Info("No Error Found")
	return false
}

// UUIDGenerator generating an uuid
func UUIDGenerator() string {
	id := uuid.New()
	logger.GetMyLogger().Info("UUID Generated....")
	return id.String()
}

// successful  response generation
func ConstructSuccessResponse(statusCode int, id int, message string) *dto.SuccessfulResponse {
	basicResponse := dto.Response{
		Id:         id,
		StatusCode: statusCode,
	}
	logger.GetMyLogger().Info(StatusCodeToMessage[statusCode])
	successResponse := dto.SuccessfulResponse{basicResponse, message}
	return &successResponse
}

// error response generation
func ConstructErrorResponse(statusCode int, id int, message string) *dto.ErrorResponse {
	basicResponse := dto.Response{
		Id:         id,
		StatusCode: statusCode,
	}
	logger.GetMyLogger().Error(StatusCodeToMessage[statusCode])
	errorResponse := dto.ErrorResponse{basicResponse, message}
	return &errorResponse
}
