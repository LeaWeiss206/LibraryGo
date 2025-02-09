package utils

import (
    "LibraryGo/internal/model"
    "encoding/json"
    "net/http"
    "time"

    "github.com/google/uuid"
)

// ResponseBuilder helps construct API responses
type ResponseBuilder struct {
    response model.APIResponse
}

// NewResponse creates a new ResponseBuilder
func NewResponse() *ResponseBuilder {
    return &ResponseBuilder{
        response: model.APIResponse{
            RequestID: uuid.New().String(),
            Timestamp: time.Now(),
        },
    }
}

// WithSuccess sets the success status
func (rb *ResponseBuilder) WithSuccess(success bool) *ResponseBuilder {
    rb.response.Success = success
    return rb
}

// WithStatus sets the HTTP status information
func (rb *ResponseBuilder) WithStatus(code int) *ResponseBuilder {
    rb.response.Status = model.StatusInfo{
        Code:    code,
        Message: http.StatusText(code),
    }
    return rb
}

// WithData sets the response data
func (rb *ResponseBuilder) WithData(data interface{}) *ResponseBuilder {
    rb.response.Data = data
    return rb
}

// WithError sets the error information
func (rb *ResponseBuilder) WithError(code, message, details string) *ResponseBuilder {
    rb.response.Error = &model.ErrorInfo{
        Code:    code,
        Message: message,
        Details: details,
    }
    return rb
}

// WithMeta sets the metadata
func (rb *ResponseBuilder) WithMeta(meta *model.MetaData) *ResponseBuilder {
    rb.response.Meta = meta
    return rb
}

// Send writes the response to http.ResponseWriter
func (rb *ResponseBuilder) Send(w http.ResponseWriter, statusCode int) error {
    rb.WithStatus(statusCode) // Set the status code in the response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    return json.NewEncoder(w).Encode(rb.response)
}