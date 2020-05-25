package utils

import "testing"

func TestIsEmptyString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name        string
		args        args
		wantIsEmpty bool
	}{
		{name: "Should return true", args: args{str: ""}, wantIsEmpty: true},
		{name: "Should return false", args: args{str: "a"}, wantIsEmpty: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsEmpty := IsEmptyString(tt.args.str); gotIsEmpty != tt.wantIsEmpty {
				t.Errorf("IsEmptyString() = %v, want %v", gotIsEmpty, tt.wantIsEmpty)
			}
		})
	}
}
