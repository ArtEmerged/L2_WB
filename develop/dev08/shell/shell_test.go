package shell

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {

	tests := []struct {
		name    string
		input   string
		wantOut string
		wantErr bool
	}{
		{
			name:    "cd",
			input:   "cd ..\n/quit",
			wantOut: "",
			wantErr: false,
		},
		{
			name:    "cd: too many arg",
			input:   "cd .. ..\n/quit",
			wantOut: ErrCdTooManyArg.Error() + "\n",
			wantErr: false,
		},
		{
			name:    "cd: no such file or dir",
			input:   "cd qwertyuiop123456\n/quit",
			wantOut: ErrCdNoSuchFileOrDir.Error() + "\n",
			wantErr: false,
		},
		{
			name:    "echo",
			input:   "echo Hello Everyone Golang\n/quit",
			wantOut: "Hello Everyone Golang\n",
			wantErr: false,
		},
		{
			name:    "kill: empty arg",
			input:   "kill\n/quit",
			wantOut: ErrKill.Error() + "\n",
			wantErr: false,
		},
		{
			name:    "kill: not a number",
			input:   "kill GO\n/quit",
			wantOut: ErrKill.Error() + "\n",
			wantErr: false,
		},
		{
			name:    "kill: no such process",
			input:   "kill -10000\n/quit",
			wantOut: ErrKillNoSuchProcess.Error() + "\n",
			wantErr: false,
		},
		{
			name:    "exec",
			input:   "echo 1.2.3.4.5 | cut -d. -f 2,4 \n/quit",
			wantOut: "2.4\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			out := &bytes.Buffer{}
			if err := Run(reader, out); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("Run() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
