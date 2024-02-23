package telnet

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name    string
		in      io.Reader
		args    []string
		wantOut string
		wantErr error
	}{
		{
			name:    "Valid Input",
			in:      bytes.NewBufferString("Hello\n"),
			args:    []string{"127.0.0.1", "5050"},
			wantOut: "server: Hello\nConnection closed.\n",
		},
		{
			name:    "Connection closure",
			in:      bytes.NewBufferString(""),
			args:    []string{"127.0.0.1", "5050"},
			wantOut: "Connection closed.\n",
		},
		{
			name:    "Miss Args",
			in:      bytes.NewBufferString(""),
			args:    []string{"5050"},
			wantOut: "",
			wantErr: ErrMissArgs,
		},
		{
			name:    "Invalid Port",
			in:      bytes.NewBufferString(""),
			args:    []string{"127.0.0.1", "invalidport"},
			wantOut: "",
			wantErr: fmt.Errorf("dial tcp: lookup tcp/%s: unknown port", "invalidport"),
		},
		{
			name:    "Invalid Host",
			in:      bytes.NewBufferString(""),
			args:    []string{"invalidhost", "5050"},
			wantOut: "",
			wantErr: fmt.Errorf("dial tcp: lookup %s: i/o timeout", "invalidhost"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			if err := Run(tt.in, out, tt.args); err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("Run() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
