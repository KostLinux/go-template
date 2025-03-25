package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"time"

	"go-template/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type responseCapturer struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (resp *responseCapturer) Write(body []byte) (int, error) {
	resp.body.Write(body)
	return resp.ResponseWriter.Write(body)
}

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestTime := time.Now()
		traceID := uuid.New().String()

		// Set trace ID
		ctx.Set("trace_id", traceID)
		ctx.Header("X-Trace-ID", traceID)

		// Read request body
		var requestBody interface{}
		if ctx.Request.Body != nil && ctx.Request.ContentLength > 0 {
			bodyBytes, _ := io.ReadAll(ctx.Request.Body)
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			if err := json.Unmarshal(bodyBytes, &requestBody); err != nil {
				logger.WarnwFields("Failed to parse request body", map[string]interface{}{
					"error":    err.Error(),
					"trace_id": traceID,
				})
			}
		}

		// Create custom response writer
		resp := &responseCapturer{
			ResponseWriter: ctx.Writer,
			body:           &bytes.Buffer{},
		}
		ctx.Writer = resp

		// Process request
		ctx.Next()

		// Capture response time
		responseTime := time.Now()

		// Parse response body
		var responseBody interface{}
		if resp.body.Len() > 0 {
			if err := json.Unmarshal(resp.body.Bytes(), &responseBody); err != nil {
				logger.WarnwFields("Failed to unmarshal response body", map[string]interface{}{
					"error":    err.Error(),
					"trace_id": traceID,
				})
				responseBody = resp.body.String()
			}

			if responseBody == nil {
				responseBody = gin.H{}
			}
		}

		// Create fields map for structured logging
		fields := map[string]interface{}{
			"trace_id":           traceID,
			"request_timestamp":  requestTime,
			"response_timestamp": responseTime,
			"method":             ctx.Request.Method,
			"path":               ctx.Request.URL.Path,
			"ip":                 ctx.ClientIP(),
			"user_agent":         ctx.Request.UserAgent(),
			"status_code":        ctx.Writer.Status(),
		}

		// Add optional fields if they have values
		if ctx.Request.URL.RawQuery != "" {
			fields["query"] = ctx.Request.URL.RawQuery
		}

		if requestBody != nil {
			fields["request_body"] = requestBody
		}

		if responseBody != nil {
			fields["response_body"] = responseBody
		}

		if latency := responseTime.Sub(requestTime).Milliseconds(); latency > 0 {
			fields["latency_ms"] = latency
		}

		if err := ctx.Errors.String(); err != "" {
			fields["error"] = err
		}

		// Log based on status code with structured fields
		switch status := ctx.Writer.Status(); {
		case status >= 500:
			logger.ErrorwFields("Request failed", fields)
		case status >= 400:
			logger.WarnwFields("Request error", fields)
		case status >= 200 && status < 400:
			logger.InfowFields("Request completed", fields)
		}
	}
}
