package whatsapp

import "testing"

func TestFormatChatID(t *testing.T) {
	tests := []struct {
		name  string
		phone string
		want  string
	}{
		{
			name:  "russian phone",
			phone: "79001234567",
			want:  "79001234567@c.us",
		},
		{
			name:  "us phone",
			phone: "14155551234",
			want:  "14155551234@c.us",
		},
		{
			name:  "short phone",
			phone: "1234567890",
			want:  "1234567890@c.us",
		},
		{
			name:  "empty phone",
			phone: "",
			want:  "@c.us",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatChatID(tt.phone)
			if got != tt.want {
				t.Errorf("FormatChatID(%q) = %q, want %q", tt.phone, got, tt.want)
			}
		})
	}
}
