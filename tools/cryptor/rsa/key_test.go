package rsa

import "testing"

func TestGetRsaKey(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetRsaKey(); (err != nil) != tt.wantErr {
				t.Errorf("GetRsaKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
