package grep

import (
	"bytes"
	"strings"
	"testing"
)

func TestGrep(t *testing.T) {
	testsCase := []struct {
		name      string
		inputData string
		options   *Options
		want      string
		wantErr   error
	}{
		{
			name:      "basicRegexp",
			inputData: "",
			options:   NewOptions("book *about", []string{"testdata/input.txt"}, 0, 0, 0, false, false, false, false, false),
			want:      "or any book    about automata theory.\n",
			wantErr:   nil,
		},
		{
			name:      "flag -A After",
			inputData: "",
			options:   NewOptions("identify", []string{"testdata/input.txt"}, 10, 0, 0, false, false, false, false, false),
			want:      "expression and identify the matched text.\nTheir names are matched by this regular expression:\n",
			wantErr:   nil,
		},
		{
			name:      "flag -B Before",
			inputData: "",
			options:   NewOptions("implements", []string{"testdata/input.txt"}, 0, 20, 0, false, false, false, false, false),
			want:      "Overview\nPackage regexp implements regular expression search.\n",
			wantErr:   nil,
		},
		{
			name:      "flag -C Context",
			inputData: "",
			options:   NewOptions("Python,", []string{"testdata/input.txt"}, 0, 1, 2, false, false, false, false, false),
			want:      "same general syntax used by Perl,\nPython,\nand other languages.\nMore precisely,\n",
			wantErr:   nil,
		},
		{
			name:      "flag -c Count",
			inputData: "",
			options:   NewOptions("regexp", []string{"testdata/input.txt"}, 0, 0, 0, true, true, false, false, false),
			want:      "5\n",
			wantErr:   nil,
		},
		{
			name:      "flag -i IgnoreCase",
			inputData: "",
			options:   NewOptions("more", []string{"testdata/input.txt"}, 0, 0, 0, false, true, false, false, false),
			want:      "More precisely,\nFor more information about this property, see\n",
			wantErr:   nil,
		},
		{
			name:      "flag -v Invert",
			inputData: "Dog\nCat\nTom\nMouse\nJerry\n",
			options:   NewOptions("Dog", []string{}, 0, 0, 0, false, false, true, false, false),
			want:      "Cat\nTom\nMouse\nJerry\n",
			wantErr:   nil,
		},
		{
			name:      "flag -F Fixed",
			inputData: "",
			options:   NewOptions("y.", []string{"testdata/input.txt"}, 0, 0, 0, false, false, false, true, false),
			want:      "or any book    about automata theory.\n",
			wantErr:   nil,
		},
		{
			name:      "flag -n LineNum",
			inputData: "",
			options:   NewOptions("Python", []string{"testdata/input.txt"}, 0, 0, 1, false, false, false, false, true),
			want:      "6-same general syntax used by Perl,\n7:Python,\n8-and other languages.\n",
			wantErr:   nil,
		},
		{
			name:      "Two file",
			inputData: "",
			options:   NewOptions("Python", []string{"testdata/input.txt", "testdata/input2.txt"}, 0, 0, 0, false, false, false, false, true),
			want:      "testdata/input.txt:7:Python,\ntestdata/input2.txt:7:Python,\n",
			wantErr:   nil,
		},
		{
			name:      "No such file",
			inputData: "",
			options:   NewOptions("Python", []string{"wrong_path.txt"}, 0, 0, 0, false, false, false, false, false),
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

