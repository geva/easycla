// Copyright The Linux Foundation and each contributor to CommunityBridge.
// SPDX-License-Identifier: MIT

package utils

import (
	"fmt"

	v1Models "github.com/communitybridge/easycla/cla-backend-go/gen/models"
	"github.com/communitybridge/easycla/cla-backend-go/gen/v2/models"
)

// ToV1ErrorResponse is a wrapper function to convert a v2 swagger error response to a v1 swagger error response
func ToV1ErrorResponse(err *models.ErrorResponse) *v1Models.ErrorResponse {
	return &v1Models.ErrorResponse{
		Code:       err.Code,
		Message:    err.Message,
		XRequestID: err.XRequestID,
	}
}

// ErrorResponseBadRequest Helper function to generate a bad request error response
func ErrorResponseBadRequest(reqID, msg string) *models.ErrorResponse {
	return &models.ErrorResponse{
		Code:       String400,
		Message:    fmt.Sprintf("%s - %s", EasyCLA400BadRequest, msg),
		XRequestID: reqID,
	}
}

// ErrorResponseBadRequestWithError Helper function to generate a bad request error response
func ErrorResponseBadRequestWithError(reqID, msg string, err error) *models.ErrorResponse {
	return &models.ErrorResponse{
		Code:       String400,
		Message:    fmt.Sprintf("%s - %s - error: %+v", EasyCLA400BadRequest, msg, err),
		XRequestID: reqID,
	}
}

// ErrorResponseForbidden Helper function to generate a forbidden error response
func ErrorResponseForbidden(reqID, msg string) *models.ErrorResponse {
	return &models.ErrorResponse{
		Code:       String403,
		Message:    fmt.Sprintf("%s - %s", EasyCLA403Forbidden, msg),
		XRequestID: reqID,
	}
}

// ErrorResponseForbiddenWithError Helper function to generate a forbidden error response
func ErrorResponseForbiddenWithError(reqID, msg string, err error) *models.ErrorResponse {
	return &models.ErrorResponse{
		Code:       String403,
		Message:    fmt.Sprintf("%s - %s - error: %+v", EasyCLA403Forbidden, msg, err),
		XRequestID: reqID,
	}
}

// ErrorResponseNotFound Helper function to generate a not found error response
func ErrorResponseNotFound(reqID, msg string) *models.ErrorResponse {
	return &models.ErrorResponse{
		Code:       String404,
		Message:    fmt.Sprintf("%s - %s", EasyCLA404NotFound, msg),
		XRequestID: reqID,
	}
}

// ErrorResponseNotFoundWithError Helper function to generate a not found error response
func ErrorResponseNotFoundWithError(reqID, msg string, err error) *models.ErrorResponse {
	return &models.ErrorResponse{
		Code:       String404,
		Message:    fmt.Sprintf("%s - %s - error: %+v", EasyCLA404NotFound, msg, err),
		XRequestID: reqID,
	}
}

// ErrorResponseInternalServerError Helper function to generate an internal server error response
func ErrorResponseInternalServerError(reqID, msg string) *models.ErrorResponse {
	return &models.ErrorResponse{
		Code:       String500,
		Message:    fmt.Sprintf("%s - %s", EasyCLA500InternalServerError, msg),
		XRequestID: reqID,
	}
}

// ErrorResponseInternalServerErrorWithError Helper function to generate an internal server error response
func ErrorResponseInternalServerErrorWithError(reqID, msg string, err error) *models.ErrorResponse {
	return &models.ErrorResponse{
		Code:       String500,
		Message:    fmt.Sprintf("%s - %s - error: %+v", EasyCLA500InternalServerError, msg, err),
		XRequestID: reqID,
	}
}
