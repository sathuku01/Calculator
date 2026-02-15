package parser

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []Token
		wantErr bool
	}{
		{
			name:  "simple addition",
			input: "2+3",
			want: []Token{
				{Number: 2, IsNum: true},
				{Operator: "+", IsNum: false},
				{Number: 3, IsNum: true},
			},
			wantErr: false,
		},
		{
			name:  "mixed operators",
			input: "10*5-2",
			want: []Token{
				{Number: 10, IsNum: true},
				{Operator: "*"},
				{Number: 5, IsNum: true},
				{Operator: "-"},
				{Number: 2, IsNum: true},
			},
			wantErr: false,
		},
		{
			name:    "invalid character",
			input:   "2+a",
			wantErr: true,
		},
		{
			name:    "operator first",
			input:   "+2",
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   "",
			wantErr: false, // depends on your implementation
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Tokenize(tt.input)

			if (err != nil) != tt.wantErr {
				t.Fatalf("expected error: %v, got error: %v", tt.wantErr, err)
			}

			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected %+v, got %+v", tt.want, got)
			}
		})
	}
}
