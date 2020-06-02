package model

import (
	"strconv"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		user    User
		wantErr bool
	}{
		{User{}, true},
		{User{RegNo: "CA1"}, true},
		{User{RegNo: "CA1", Name: "name1"}, true},
		{User{RegNo: "CA1", Name: "name1", Phone: "123456789"}, false},
	}

	for i, test := range tests {
		testCase := "t" + strconv.Itoa(i)
		t.Run(testCase, func(tt *testing.T) {
			err := test.user.Validate()
			if err != nil && !test.wantErr {
				tt.Errorf("expected no error but got: %v", err)
			} else if err == nil && test.wantErr {
				tt.Error("expected error but got nil")
			}
		})
	}
}
