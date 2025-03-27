package goBinaryConverter

import (
	"testing"
)

func TestImpl_ToBinary(t *testing.T) {
	tests := []struct {
		name   string
		args   int64
		expect string
	}{
		{
			name:   "Success - convert decimal to binary",
			args:   30,
			expect: "11110",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			converter := NewBinaryConverter()
			actual := converter.ToBinary(tt.args)
			if actual != tt.expect {
				t.Errorf("BinaryConverter.ToBinary() = %v, want %v", actual, tt.expect)
			}
		})
	}
}

func TestImpl_ToDecimal(t *testing.T) {
	tests := []struct {
		name   string
		args   string
		expect int64
		error  bool
	}{
		{
			name:   "Success - convert binary to decimal",
			args:   "11110",
			expect: 30,
			error:  false,
		},
		{
			name:   "Error - convert binary to decimal",
			args:   "AB120",
			expect: 0,
			error:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			converter := NewBinaryConverter()
			actual, err := converter.ToDecimal(tt.args)
			if (err != nil) != tt.error {
				t.Errorf("BinaryConverter.ToDecimal() error = %v, wantErr %v", err, tt.error)
			}

			if actual != tt.expect {
				t.Errorf("BinaryConverter.ToBinary() = %v, want %v", actual, tt.expect)
			}
		})
	}
}
