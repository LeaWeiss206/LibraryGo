package model

import (
    "time"
)

// APIResponse represents the standard response structure for all API endpoints
type APIResponse struct {
    Success    bool        `json:"success"`
    Status     StatusInfo  `json:"status"`          // Added StatusInfo struct
    Data       interface{} `json:"data,omitempty"`
    Error      *ErrorInfo  `json:"error,omitempty"`
    Meta       *MetaData   `json:"meta,omitempty"`
    RequestID  string      `json:"requestId"`
    Timestamp  time.Time   `json:"timestamp"`
}

// StatusInfo contains HTTP status information
type StatusInfo struct {
    Code    int    `json:"code"`    // HTTP status code
    Message string `json:"message"` // HTTP status message
}

// ErrorInfo contains detailed error information
type ErrorInfo struct {
    Code    string `json:"code"`              // Machine-readable error code
    Message string `json:"message"`           // Human-readable error message
    Details string `json:"details,omitempty"` // Additional error details if any
}

// MetaData contains additional information about the response
type MetaData struct {
    Total       int       `json:"total,omitempty"`
    Count       int       `json:"count,omitempty"`
    Page        int       `json:"page,omitempty"`
    PerPage     int       `json:"perPage,omitempty"`
    TotalPages  int       `json:"totalPages,omitempty"`
    NextPage    *int      `json:"nextPage,omitempty"`
    PrevPage    *int      `json:"prevPage,omitempty"`
    ProcessedAt time.Time `json:"processedAt,omitempty"`
}
