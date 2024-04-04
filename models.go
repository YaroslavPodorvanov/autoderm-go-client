package autoderm_go_client

type QueryResponse struct {
	Success     bool         `json:"success"`
	Message     string       `json:"message"`
	ID          string       `json:"id"`
	Predictions []Prediction `json:"predictions"`
}

type Prediction struct {
	Confidence       float64 `json:"confidence"`
	Icd              string  `json:"icd"`
	Name             string  `json:"name"`
	ClassificationID string  `json:"classificationId"`
	ReadMoreURL      string  `json:"readMoreUrl"`
}
