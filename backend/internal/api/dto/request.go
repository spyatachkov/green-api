package dto

type SendMessageRequest struct {
	IDInstance       string `json:"idInstance"`
	APITokenInstance string `json:"apiTokenInstance"`
	ChatID           string `json:"chatId"`
	Message          string `json:"message"`
}

type SendFileRequest struct {
	IDInstance       string `json:"idInstance"`
	APITokenInstance string `json:"apiTokenInstance"`
	ChatID           string `json:"chatId"`
	FileURL          string `json:"fileUrl"`
	FileName         string `json:"fileName,omitempty"`
}
