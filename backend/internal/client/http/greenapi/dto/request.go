package dto

type SendMessageRequest struct {
	ChatID  string `json:"chatId"`
	Message string `json:"message"`
}

type SendFileRequest struct {
	ChatID   string `json:"chatId"`
	URLFile  string `json:"urlFile"`
	FileName string `json:"fileName,omitempty"`
}
