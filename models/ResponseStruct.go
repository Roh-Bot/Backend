package models

type Response struct {
	StatusCode int
	Error      map[string]string
	Data       map[string]string
}

func (r *Response) ResponseWriter(statusCode int, error map[string]string, data map[string]string) *Response {
	r.StatusCode = statusCode
	r.Error = error
	r.Data = data
	return r
}
