package whatsapp

import "fmt"

func FormatChatID(phone string) string {
	return fmt.Sprintf("%s@c.us", phone)
}
