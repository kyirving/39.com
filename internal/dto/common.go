package dto

type CommonParams struct {
	AccessToken string `json:"access_token"`
	Udid        string `json:"udid" binding:"required"`
	Timestamp   string `json:"timestamp" binding:"required"`
	Version     string `json:"version" binding:"required"`
	SignType    string `json:"sign_type" binding:"required"`
	Sign        string `json:"sign" binding:"required"`
	RequestId   string `json:"request_id" binding:"required"`
}
