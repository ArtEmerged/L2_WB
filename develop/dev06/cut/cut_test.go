package cut

import (
	"bytes"
	"strings"
	"testing"
)

func TestCut(t *testing.T) {
	testsCase := []struct {
		name      string
		inputData string
		options   *Options
		want      string
		wantErr   error
	}{
		{
			name:      "basic",
			inputData: "",
			options:   NewOptions([]string{"testdata/test3.txt"}, "1", "\t", false),
			want:      "Hello\nHello\nHello\n",
			wantErr:   nil,
		},
		{
			name:      "flagF_test1",
			inputData: "",
			options:   NewOptions([]string{"testdata/test3.txt"}, "1,3", "\t", false),
			want:      "Hello\tMen\nHello\tMac\nHello\tssds\n",
			wantErr:   nil,
		},
		{
			name:      "flagF_test2(range)",
			inputData: "",
			options:   NewOptions([]string{"testdata/test3.txt"}, "1-3", "\t", false),
			want:      "Hello\tGood\tMen\nHello\tGood\tMac\nHello\tdsss\tssds\n",
			wantErr:   nil,
		},
		{
			name:      "flagF_test2.1(range)",
			inputData: "",
			options:   NewOptions([]string{"testdata/test3.txt"}, "1-100", "\t", false),
			want:      "Hello\tGood\tMen\nHello\tGood\tMac\nHello\tdsss\tssds\n",
			wantErr:   nil,
		},
		{
			name:      "flagF_test2.2(two range)",
			inputData: "",
			options:   NewOptions([]string{"testdata/test4.txt"}, "4-5,1-2", ".", false),
			want:      "1.2.4.5\n",
			wantErr:   nil,
		},
		{
			name:      "flagF_test3(interval)",
			inputData: "",
			options:   NewOptions([]string{"testdata/test4.txt"}, "-4", ".", false),
			want:      "1.2.3.4\n",
			wantErr:   nil,
		},
		{
			name:      "flagF_test3.1(interval)",
			inputData: "",
			options:   NewOptions([]string{"testdata/test4.txt"}, "12-", ".", false),
			want:      "12.13.14.15.16\n",
			wantErr:   nil,
		},
		{
			name:      "flagF_test3.3(interval)",
			inputData: "",
			options:   NewOptions([]string{"testdata/test4.txt"}, "-4,12-", ".", false),
			want:      "1.2.3.4.12.13.14.15.16\n",
			wantErr:   nil,
		},
		{
			name:      "flagF_test3.4(interval)",
			inputData: "",
			options:   NewOptions([]string{"testdata/test4.txt"}, "12-,-4", ".", false),
			want:      "1.2.3.4.12.13.14.15.16\n",
			wantErr:   nil,
		},
		{
			name:      "flagD_test1",
			inputData: "",
			options:   NewOptions([]string{"testdata/test1.txt"}, "1,3,2", ":", false),
			want:      "Winter: white: snow\nSpring: green: grass\nSummer: colorful: blossom\nAutumn: yellow: leaves\n",
			wantErr:   nil,
		},
		{
			name:      "flagS_test1",
			inputData: "",
			options:   NewOptions([]string{"testdata/test2.txt"}, "2", ":", true),
			want:      " extraordinarily particular about politeness in others.\n extraordinarily particular about politeness\n",
			wantErr:   nil,
		},
		{
			name:      "two file test",
			inputData: "",
			options:   NewOptions([]string{"testdata/test1.txt", "testdata/test2.txt"}, "1", ":", true),
			want:      "Winter\nSpring\nSummer\nAutumn\n1.He was\n3.He was\n",
			wantErr:   nil,
		},
		{
			name:      "stdin test",
			inputData: "1.2.3.4.5.6.7.8.9.10.11.12.13.14.15.16",
			options:   NewOptions([]string{}, "5,2-6,3-7,14-,8,7,-3", ".", false),
			want:      "1.2.3.4.5.6.7.8.14.15.16\n",
			wantErr:   nil,
		},
		{
			name:      "error check1:Not Character",
			inputData: "",
			options:   NewOptions([]string{"testdata/test2.txt"}, "2", "aa", false),
			want:      "",
			wantErr:   ErrNotCharacter,
		},
		{
			name:      "error check2:Not Number",
			inputData: "",
			options:   NewOptions([]string{"testdata/test2.txt"}, "two", ":", false),
			want:      "",
			wantErr:   ErrNotNumber,
		},
		{
			name:      "error check3:Nubmer Less One",
			inputData: "",
			options:   NewOptions([]string{"testdata/test2.txt"}, "0", ":", false),
			want:      "",
			wantErr:   ErrNubmerLessOne,
		},
		{
			name:      "error check3.1:Nubmer Less One",
			inputData: "",
			options:   NewOptions([]string{"testdata/test2.txt"}, "", ":", false),
			want:      "",
			wantErr:   ErrNubmerLessOne,
		},
		{
			name:      "error check4:Invalid Range",
			inputData: "",
			options:   NewOptions([]string{"testdata/test4.txt"}, "4-1", ".", false),
			want:      "",
			wantErr:   ErrInvalidRange,
		},
		{
			name:      "error check5:No Such File",
			inputData: "",
			options:   NewOptions([]string{"wrong_path.txt"}, "1", ":", false),
			want:      "",
			wantErr:   ErrNoSuchFile,
		},
	}

	for _, test := range testsCase {
		t.Run(test.name, func(t *testing.T) {
			reader := strings.NewReader(test.inputData)
			var buffer bytes.Buffer
			gotErr := Run(reader, &buffer, test.options)
			got := buffer.String()
			if gotErr != nil {
				if gotErr.Error() != test.wantErr.Error() {
					t.Errorf("gotErr:%s wantErr:%s", gotErr.Error(), test.wantErr.Error())
				}
			}
			if got != test.want {
				t.Errorf("got:%s want:%s", got, test.want)
			}
		})
	}

}
