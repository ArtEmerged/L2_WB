package main

import (
	"reflect"
	"testing"
)

func TestAnagram(t *testing.T) {
	testsCases := []struct {
		name       string
		dictionary []string
		want       map[string][]string
		wantErr    error
	}{
		{
			name:       "basic_test",
			dictionary: []string{"пятак", "тяпка", "пятка", "пудра", "пятка", "листок", "мох", "ГоРа", "рОга", "слиток", "столик"},
			want:       map[string][]string{"гора": {"гора", "рога"}, "листок": {"листок", "слиток", "столик"}, "пятак": {"пятак", "пятка", "тяпка"}},
			wantErr:    nil,
		},
		{
			name:       "error_test",
			dictionary: []string{"пятак", "тяпка", "пяtка", "hello", "пудра", "пятка", "листок", "мох", "ГоРа", "рОга", "слиток", "столик"},
			want:       nil,
			wantErr:    ErrLetter,
		},
		{
			name:       "empty_test",
			dictionary: []string{},
			want:       map[string][]string{},
			wantErr:    nil,
		},
	}
	for _, test := range testsCases {
		t.Run(test.name, func(t *testing.T) {
			got, gotErr := searchForAnagramSets(&test.dictionary)
			if gotErr != test.wantErr {
				t.Errorf("gotErr %v, want %v", gotErr, test.wantErr)
			}
			if gotErr != nil {
				return
			}
			if !reflect.DeepEqual(*got, test.want) {
				t.Errorf("got %v, want %v", *got, test.want)
			}
		})

	}
}
