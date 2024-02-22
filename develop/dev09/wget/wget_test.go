package wget

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		wantErr        bool
		expectedFile   string // ожидаемое имя файла
		expectedErrMsg string 
	}{
		{
			name:           "Download HTML file",
			args:           []string{"http://google.com"},
			wantErr:        false,
			expectedFile:   "google.com.html",
			expectedErrMsg: "",
		},
		{
			name:           "Download gif file",
			args:           []string{"https://media.tenor.com/1MfIiC1yllEAAAAM/cat-huh-cat-huh-etr.gif"},
			wantErr:        false,
			expectedFile:   "cat-huh-cat-huh-etr.gif",
			expectedErrMsg: "",
		},
		{
			name:           "Download image my name",
			args:           []string{"-o", "my_cat.jpg", "https://upload.wikimedia.org/wikipedia/commons/f/f9/Surprised_young_cat.JPG"},
			wantErr:        false,
			expectedFile:   "my_cat.jpg",
			expectedErrMsg: "",
		},
		{
			name:           "Missing URL",
			args:           []string{},
			wantErr:        true,
			expectedFile:   "",
			expectedErrMsg: ErrMissURL.Error(),
		},
		{
			name:           "Not Found",
			args:           []string{"https://github.com/Go/Go"},
			wantErr:        true,
			expectedFile:   "",
			expectedErrMsg: fmt.Sprintf("wget: ERROR %d: %s", http.StatusNotFound, http.StatusText(http.StatusNotFound)),
		},
		{
			name:           "Invalid URL",
			args:           []string{"invalid-url"},
			wantErr:        true,
			expectedFile:   "",
			expectedErrMsg: "Get \"invalid-url\": unsupported protocol scheme \"\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Run(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err.Error() != tt.expectedErrMsg {
				t.Errorf("Run() error = %v, want error message %v", err, tt.expectedErrMsg)
			}

			if !tt.wantErr {
				if _, err := os.Stat(tt.expectedFile); os.IsNotExist(err) {
					t.Errorf("Run() expected file %q, but not found", tt.expectedFile)
				}
				defer os.Remove(tt.expectedFile)
			}
		})
	}
}
