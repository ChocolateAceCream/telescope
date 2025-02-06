package response

type ClassifyResponse struct {
	ClassName  string  `json:"class_name"`
	Confidence float64 `json:"confidence"`
}
