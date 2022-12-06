package endpoints

type HelloRequest struct {
	Name string `json:"name,omitempty"`
}

type HelloResponse struct {
	Phrase string `json:"phrase"`
	Err    string `json:"err,omitempty"`
}
