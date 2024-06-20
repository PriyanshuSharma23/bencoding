package bencoding

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	var buff *bytes.Buffer = nil
	var enc *Encoder = nil

	var setup = func() {
		buff = bytes.NewBuffer([]byte{})
		enc = NewEncoder(buff)
	}

	tests := []struct {
		name     string
		input    any
		expected string
		hasError bool
	}{
		{"test string", "hello", "5:hello", false},
		{"test number", 5, "i5e", false},
		{"test list of strings", []any{"hello", "world"}, "l5:hello5:worlde", false},
		{"test list of mixed types", []any{"hello", 42}, "l5:helloi42ee", false},
		{"test empty list", []any{}, "le", false},
		{"test dictionary", map[string]any{"age": 25, "name": "john"}, "d3:agei25e4:name4:johne", false},
		{"test empty dictionary", map[string]any{}, "de", false},
		{"test unsupported type", 3.14, "", true},
		{"test non-string map key", map[any]any{1: "one"}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup()
			err := enc.Encode(tt.input)
			if tt.hasError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if buff.String() != tt.expected {
				t.Errorf("expected %s got %s", tt.expected, buff.String())
			}
		})
	}
}
