package service

import (
	"calculator/internal/core"
	"calculator/internal/service"
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      []service.Token
		expectErr error
	}{
		{
			name:  "simple addition",
			input: "3+5",
			want: []service.Token{
				{Number: 3, IsNum: true},
				{Operator: "+", IsNum: false},
				{Number: 5, IsNum: true},
			},
			expectErr: nil,
		},
		{
			name:  "subtraction and multiplication",
			input: "12-4*2",
			want: []service.Token{
				{Number: 12, IsNum: true},
				{Operator: "-", IsNum: false},
				{Number: 4, IsNum: true},
				{Operator: "*", IsNum: false},
				{Number: 2, IsNum: true},
			},
			expectErr: nil,
		},
		{
			name:  "expression with decimals",
			input: "3.5+2.1",
			want: []service.Token{
				{Number: 3.5, IsNum: true},
				{Operator: "+", IsNum: false},
				{Number: 2.1, IsNum: true},
			},
			expectErr: nil,
		},
		{
			name:      "invalid character",
			input:     "3+a",
			want:      nil,
			expectErr: core.ErrInvalidInput,
		},
		{
			name:      "operator at start",
			input:     "+5",
			want:      nil,
			expectErr: core.ErrInvalidExpression,
		},
		{
			name:      "empty string",
			input:     "",
			want:      []service.Token{},
			expectErr: nil,
		},
		{
			name:  "multiple digit numbers",
			input: "123*456+78",
			want: []service.Token{
				{Number: 123, IsNum: true},
				{Operator: "*", IsNum: false},
				{Number: 456, IsNum: true},
				{Operator: "+", IsNum: false},
				{Number: 78, IsNum: true},
			},
			expectErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.Tokenize(tt.input)
			if err != tt.expectErr {
				t.Fatalf("Tokenize() error = %v, want %v", err, tt.expectErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
