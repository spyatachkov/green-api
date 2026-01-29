package models

type GreenAPICredentials struct {
	IDInstance       string
	APITokenInstance string
}

type SendMessageRequest struct {
	Credentials GreenAPICredentials
	ChatID      string
	Message     string
}

type SendFileRequest struct {
	Credentials GreenAPICredentials
	ChatID      string
	FileURL     string
	FileName    string
}

type GreenAPIResponse struct {
	Success bool
	Data    interface{}
	Error   string
}
