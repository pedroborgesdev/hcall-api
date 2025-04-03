package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response codes
const (
	CodeSuccess         = "success"
	CodeValidationError = "validation_error"
	CodeNotFound        = "not_found"
	CodeUnauthorized    = "unauthorized"
	CodeForbidden       = "forbidden"
	CodeTooManyRequests = "too_many_requests"
	CodeInternalError   = "internal_error"
	CodeBadRequest      = "bad_request"
	CodeRateLimitError  = "rate_limit_error"
	CodeInvalidInput    = "invalid_input"
	CodeDuplicateEntry  = "duplicate_entry"
)

// APIResponse is the standard response structure for all API endpoints
type APIResponse struct {
	Code    string      `json:"code"`            // Machine-readable code
	Message string      `json:"message"`         // Human-readable message
	Data    interface{} `json:"data,omitempty"`  // Optional data payload
	Error   interface{} `json:"error,omitempty"` // Detailed error message (only for errors)
	Status  int         `json:"status"`          // HTTP status code
}

// Standard HTTP status code mapping
var statusCodes = map[string]int{
	CodeSuccess:         http.StatusOK,
	CodeValidationError: http.StatusBadRequest,
	CodeNotFound:        http.StatusNotFound,
	CodeUnauthorized:    http.StatusUnauthorized,
	CodeForbidden:       http.StatusForbidden,
	CodeTooManyRequests: http.StatusTooManyRequests,
	CodeInternalError:   http.StatusInternalServerError,
	CodeBadRequest:      http.StatusBadRequest,
	CodeRateLimitError:  http.StatusTooManyRequests,
	CodeInvalidInput:    http.StatusBadRequest,
	CodeDuplicateEntry:  http.StatusConflict,
}

// SendSuccess sends a success response
func SendSuccess(c *gin.Context, message string, data interface{}) {
	status := statusCodes[CodeSuccess]
	c.JSON(status, APIResponse{
		Code:    CodeSuccess,
		Message: message,
		Data:    data,
		Status:  status,
	})
}

// SendError sends an error response
func SendError(c *gin.Context, code string, message string, err interface{}) {
	status := statusCodes[code]
	c.JSON(status, APIResponse{
		Code:    code,
		Message: message,
		Error:   err,
		Status:  status,
	})
}

// AbortWithError aborts the request with an error response
func AbortWithError(c *gin.Context, code string, message string, err interface{}) {
	status := statusCodes[code]
	c.AbortWithStatusJSON(status, APIResponse{
		Code:    code,
		Message: message,
		Error:   err,
		Status:  status,
	})
}

// Common response messages
const (
	MsgInvalidCredentials = "Invalid email or password"
	MsgUnauthorized       = "Unauthorized access"
	MsgForbidden          = "Access forbidden"
	MsgInvalidInput       = "Invalid input data"
	MsgNotFound           = "Resource not found"
	MsgInternalError      = "Internal server error"
	MsgDatabaseError      = "Database operation failed"
	MsgDuplicateEntry     = "Resource already exists"
	MsgValidationError    = "Validation failed"
	MsgRateLimitError     = "Rate limit exceeded"
	MsgSuccess            = "Operation successful"
	MsgTooManyRequests    = "Too many requests"
	MsgBadRequest         = "Bad request"
)
