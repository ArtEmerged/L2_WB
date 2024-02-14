package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		want    string
		wantErr error
	}{
		{"test1", "a4bc2d5e", "aaaabccddddde", nil},
		{"test2", "abcd", "abcd", nil},
		{"test3", "45", "", ErrInvalidString},
		{"test4", "", "", nil},
		{"test5", "3abc", "", ErrInvalidString},
		{"test6", "aaa10b", "", ErrInvalidString},
		{"test7", "aaa0b", "aab", nil},
		{"test8", "Упс3.3", "Упссс...", nil},
		{"test8", "", "", nil},
	}

	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				got, goterr := UnpackString(test.text)
				if goterr != test.wantErr {
					t.Errorf("UnpackString() goterr = %v , want = %v", goterr, test.wantErr)
				}
				if got != test.want {
					t.Errorf("UnpackString() got = %v , want = %v", got, test.want)
				}
			},
		)
	}
}
