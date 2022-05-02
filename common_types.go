package altcraft

type baseList struct {
	Error      int    `json:"error"`
	ErrorText  string `json:"error_text"`
	TotalCount int    `json:"total_count"`
}
