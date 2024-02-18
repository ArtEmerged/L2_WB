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
	}{
		{
			name:       "basic_test",
			dictionary: []string{"пятак", "тяпка", "пятка", "пудра", "пятка", "листок", "мох", "ГоРа", "рОга", "слиток", "столик"},
			want:       map[string][]string{"гора": {"гора", "рога"}, "листок": {"листок", "слиток", "столик"}, "пятак": {"пятак", "пятка", "тяпка"}},
		},
		{
			name:       "incorrect letter entry_test",
			dictionary: []string{"пятак", "тяпка", "пяtка", "hello", "пудра", "пятка", "листок", "мох", "ГоРа", "рОга", "слиток", "столик"},
			want:       map[string][]string{},
		},
		{
			name:       "empty_test",
			dictionary: []string{},
			want:       map[string][]string{},
		},
	}
	for _, test := range testsCases {
		t.Run(test.name, func(t *testing.T) {
			got := searchForAnagramSets(&test.dictionary)

			if !reflect.DeepEqual(*got, test.want) {
				t.Errorf("got %v, want %v", *got, test.want)
			}
		})

	}
}
