package response

// Response is wrap response api
type Response struct {
	Data       interface{}       `json:"data,omitempty"`
	Link       NextPrevlinks     `json:"link,omitempty'"`
	Meta       interface{}       `json:"meta,omitempty"`
	Message    string            `json:"message,omitempty"`
	StatusCode int               `json:"status_code,omitempty"`
	Status     string            `json:"status,omitempty"`
	Error      interface{}       `json:"error,omitempty"`
	Parameters interface{}       `json:"parameters,omitempty"`
	Body       string            `json:"body,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
}

//NextPrevlinks -
type NextPrevlinks struct {
	Self    string `json:"self,omitempty"`
	Next    string `json:"next,omitempty"`
	Preview string `json:"preview,omitempty"`
}
