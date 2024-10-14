package hash

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ValidPassword", args{"password123"}, false},
		{"EmptyPassword", args{""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := HashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComparePassword(t *testing.T) {
	type args struct {
		hashedPassword string
		password       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ValidPassword", args{"$2a$10$hE/A8m6Dpdv4vN5HNVoBHeh1EGEtHiqUVK5p1lwhNedEh8J82PSXa", "password123"}, false},
		{"InvalidPassword", args{"$2a$10$7EqJtq98hPqEX7fNZaFWoOa8u5FQw1u1Q1u1Q1u1Q1u1Q1u1Q1u1Q", "wrongpassword"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ComparePassword(tt.args.hashedPassword, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComparePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
