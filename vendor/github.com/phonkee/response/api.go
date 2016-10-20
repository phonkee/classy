package response

import (
	"context"
	"net/http"
)

const (
	statusKey = 118999
)

/*
New returns new response instance, if no status is given StatusOK is used
*/
func New(statuses ...int) (result Response) {
	result = &response{
		data:    map[string]interface{}{},
		headers: map[string]string{},
	}
	result.ContentType("application/json")
	if len(statuses) > 0 {
		result.Status(statuses[0])
	} else {
		result.Status(http.StatusOK)
	}
	return
}

/*
Body is helper to create response
*/
func Body(body interface{}) Response {
	return New().Body(body)
}

/*
Data is helper to create status ok response.
*/
func Data(key string, value interface{}) Response {
	return New().Data(key, value)
}

/*
Error is helper to create status ok response.
*/
func Error(err interface{}) Response {
	return New(http.StatusInternalServerError).Error(err)
}

/*
HTML returns response set to HTML
*/
func HTML(html string) Response {
	return New().HTML(html)
}

/*
Result is helper to create status ok response.
*/
func Result(result interface{}) Response {
	return New().Result(result)
}

/*
GetStatus returns status for given context, if not found returns default or StatusNotFound if no default given
*/
func GetStatus(ctx context.Context, def ...int) int {
	if value, ok := ctx.Value(statusKey).(int); !ok {
		if len(def) > 0 {
			return def[0]
		} else {
			return http.StatusNotFound
		}
	} else {
		return value
	}
}

/*
SetStatus sets status for given context
*/
func SetStatus(ctx context.Context, status int) context.Context {
	context.WithValue(ctx, statusKey, status)
	return ctx
}

/*
SliceResult is helper to create status ok response.
*/
func SliceResult(result interface{}) Response {
	return New().SliceResult(result)
}

/*
Write writes to response and returns error
*/
func Write(w http.ResponseWriter, r *http.Request) error {
	return New().Write(w, r)
}
