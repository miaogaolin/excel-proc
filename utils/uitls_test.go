package utils

import "testing"

func TestGetMoneyNum(t *testing.T) {
	tests := []struct {
		name    string
		money   string
		want    float64
		wantErr bool
	}{
		{"¥123.45", "¥123.45", 123.45, false},
		{"$1,234,567.89", "$1,234,567.89", 1234567.89, false},
		{"€1.234.567,89", "€1,234,567.89", 1234567.89, false},
		{"1,234,567", "1,234,567", 1234567, false},
		{`"1,234,567.89"`, `"1,234,567.89"`, 1234567.89, false},
		{"123", `123`, 123, false},
		{"¥1000.00", "¥1000.00", 1000, false},
		{"-0.09", "-0.09", -0.09, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMoneyNum(tt.money)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMoneyNum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetMoneyNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
