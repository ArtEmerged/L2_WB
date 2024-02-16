package main

import (
	"strings"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name       string
		parameters Parameters
		want       string
		wantErr    error
	}{
		{
			name:       "basic_test1",
			parameters: Parameters{1, false, false, false, []string{"./testdata/test1.txt"}},
			want:       "./testdata/test1_want.txt",
			wantErr:    nil,
		},
		{
			name:       "basic_test2",
			parameters: Parameters{1, false, false, false, []string{"./testdata/symbols_test2.txt"}},
			want:       "./testdata/symbols_test2_want.txt",
			wantErr:    nil,
		},
		{
			name:       "basic_test3",
			parameters: Parameters{1, false, false, false, []string{"./testdata/numbers_test3.txt"}},
			want:       "./testdata/numbers_test3_want.txt",
			wantErr:    nil,
		},
		{
			name:       "basic_test4",
			parameters: Parameters{1, false, false, false, []string{"./testdata/latin_alfovit_test5.txt"}},
			want:       "./testdata/latin_alfovit_test5_want.txt",
			wantErr:    nil,
		},
		{
			name:       "basic_test5",
			parameters: Parameters{1, false, false, false, []string{"./testdata/cyrillic_alphabet_test4.txt"}},
			want:       "./testdata/cyrillic_alphabet_test4_want.txt",
			wantErr:    nil,
		},
		{
			name:       "revers_test",
			parameters: Parameters{1, false, true, false, []string{"./testdata/latin_alfovit_test5.txt"}},
			want:       "./testdata/latin_alfovit_test5_want_r.txt",
			wantErr:    nil,
		},
		{
			name:       "column_test_k3",
			parameters: Parameters{3, false, false, false, []string{"./testdata/number_column_test6.txt"}},
			want:       "./testdata/number_column_test6_want_k3.txt",
			wantErr:    nil,
		},
		{
			name:       "column_test_k5",
			parameters: Parameters{5, false, false, false, []string{"./testdata/number_column_test6.txt"}},
			want:       "./testdata/number_column_test6_want_k5.txt",
			wantErr:    nil,
		},
		{
			name:       "columnErr_test",
			parameters: Parameters{0, false, false, false, []string{"./testdata/number_column_test6.txt"}},
			want:       "",
			wantErr:    ErrColumn,
		},
		{
			name:       "argsErr_test1",
			parameters: Parameters{1, false, false, false, []string{"./testdata/number_column_test6.txt", "hello"}},
			want:       "",
			wantErr:    ErrArgs,
		},
		{
			name:       "argsErr_test2",
			parameters: Parameters{1, false, false, false, []string{}},
			want:       "",
			wantErr:    ErrArgs,
		},
		{
			name:       "noSuchFileErr_test2",
			parameters: Parameters{1, false, false, false, []string{"sffwuwuwwrwo0322"}},
			want:       "",
			wantErr:    ErrNoSuchFile,
		},
		{
			name:       "nubmers_test",
			parameters: Parameters{1, true, false, false, []string{"./testdata/numbers_test3.txt"}},
			want:       "./testdata/numbers_test3_want_n.txt",
			wantErr:    nil,
		},
		{
			name:       "nubmers+column_test",
			parameters: Parameters{3, true, false, false, []string{"./testdata/numbers_test3.txt"}},
			want:       "./testdata/numbers_test3_want_nk3.txt",
			wantErr:    nil,
		},
		{
			name:       "nubmers+column+revers_test",
			parameters: Parameters{3, true, true, false, []string{"./testdata/numbers_test3.txt"}},
			want:       "./testdata/numbers_test3_want_nk3r.txt",
			wantErr:    nil,
		},
		{
			name:       "unique_test",
			parameters: Parameters{1, false, false, true, []string{"./testdata/test1.txt"}},
			want:       "./testdata/test1_want_u.txt",
			wantErr:    nil,
		},
		{
			name:       "all_flags_test",
			parameters: Parameters{3, true, true, true, []string{"./testdata/all_flags_test.txt"}},
			want:       "./testdata/all_flags_test_want.txt",
			wantErr:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, gotErr := Run(test.parameters)

			if gotErr != nil {
				if got != nil {
					t.Errorf("Run() got = %v , want = %v", got, nil)
				}
				if gotErr != test.wantErr {
					t.Errorf("Run() goterr = %v , want = %v", gotErr, test.wantErr)
				}
				return
			}

			if gotErr != test.wantErr {
				t.Errorf("Run() goterr = %v , want = %v", gotErr, test.wantErr)
			}
			wantText, _ := readFile(test.want)
			if strings.Join(got, "\n") != strings.Join(wantText, "\n") {
				t.Errorf("Run() got = \n%v\n , want = \n%v", got, wantText)
			}

		})
	}
}
