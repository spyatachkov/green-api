package validator

import (
	"testing"

	"github.com/spyatachkov/green-api/backend/internal/models"
)

func TestValidatePhone(t *testing.T) {
	tests := []struct {
		name    string
		phone   string
		wantErr error
	}{
		{
			name:    "valid phone 11 digits",
			phone:   "79001234567",
			wantErr: nil,
		},
		{
			name:    "valid phone 10 digits",
			phone:   "1234567890",
			wantErr: nil,
		},
		{
			name:    "valid phone 15 digits",
			phone:   "123456789012345",
			wantErr: nil,
		},
		{
			name:    "empty phone",
			phone:   "",
			wantErr: models.ErrPhoneEmpty,
		},
		{
			name:    "phone with spaces",
			phone:   "7 900 123 45 67",
			wantErr: models.ErrPhoneInvalidFormat,
		},
		{
			name:    "phone with plus",
			phone:   "+79001234567",
			wantErr: models.ErrPhoneInvalidFormat,
		},
		{
			name:    "phone with dashes",
			phone:   "7-900-123-45-67",
			wantErr: models.ErrPhoneInvalidFormat,
		},
		{
			name:    "phone with letters",
			phone:   "790012345ab",
			wantErr: models.ErrPhoneInvalidFormat,
		},
		{
			name:    "phone too short",
			phone:   "123456789",
			wantErr: models.ErrPhoneInvalidLength,
		},
		{
			name:    "phone too long",
			phone:   "1234567890123456",
			wantErr: models.ErrPhoneInvalidLength,
		},
		{
			name:    "only letters",
			phone:   "abcdefghij",
			wantErr: models.ErrPhoneInvalidFormat,
		},
		{
			name:    "special characters",
			phone:   "!@#$%^&*()",
			wantErr: models.ErrPhoneInvalidFormat,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePhone(tt.phone)
			if err != tt.wantErr {
				t.Errorf("ValidatePhone(%q) error = %v, wantErr %v", tt.phone, err, tt.wantErr)
			}
		})
	}
}

func TestIsValidPhone(t *testing.T) {
	tests := []struct {
		name  string
		phone string
		want  bool
	}{
		{
			name:  "valid phone",
			phone: "79001234567",
			want:  true,
		},
		{
			name:  "invalid phone with letters",
			phone: "790012345ab",
			want:  false,
		},
		{
			name:  "empty phone",
			phone: "",
			want:  false,
		},
		{
			name:  "phone with plus",
			phone: "+79001234567",
			want:  false,
		},
		{
			name:  "too short",
			phone: "123",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidPhone(tt.phone)
			if got != tt.want {
				t.Errorf("IsValidPhone(%q) = %v, want %v", tt.phone, got, tt.want)
			}
		})
	}
}
