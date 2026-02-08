package internal_http_service

type ActionRequest struct {
	Data string `json:"data"`
}

type ActionResponse struct {
	Code int64 `json:"code"`
}
