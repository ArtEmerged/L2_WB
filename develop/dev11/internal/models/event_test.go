package models

import (
	"develop/dev11/internal/errors"
	"testing"
	"time"
)

func TestEventValidate(t *testing.T) {

	tests := []struct {
		name    string
		event   *Event
		wantErr error
	}{
		{
			name:    "OK",
			event:   NewEvent(5, 0, time.Date(2026, 5, 12, 15, 04, 04, 0, time.UTC), "Test", "Test"),
			wantErr: nil,
		},
		{
			name:    "UserID must be positive",
			event:   NewEvent(-5, 0, time.Date(2026, 5, 12, 15, 04, 04, 0, time.UTC), "Test", "Test"),
			wantErr: errors.NewBadRequestError("user_id must be positive"),
		},
		{
			name:    "EventID must be positive",
			event:   NewEvent(5, -5, time.Date(2026, 5, 12, 15, 04, 04, 0, time.UTC), "Test", "Test"),
			wantErr: errors.NewBadRequestError("id must be positive"),
		},
		{
			name:    "Past Event Date",
			event:   NewEvent(5, 5, time.Date(2007, 5, 12, 15, 04, 04, 0, time.UTC), "Test", "Test"),
			wantErr: errors.NewBadRequestError("event date cannot be in the past"),
		},
		{
			name:    "Empty Title",
			event:   NewEvent(5, 5, time.Date(2026, 5, 12, 15, 04, 04, 0, time.UTC), "", "Test"),
			wantErr: errors.NewBadRequestError("empty parameter: title"),
		},
		{
			name:    "Title Too Long",
			event:   NewEvent(5, 5, time.Date(2026, 5, 12, 15, 04, 04, 0, time.UTC), string(make([]rune, 21)), "Test"),
			wantErr: errors.NewBadRequestError("title parameter is too long, maximum length 20 symbols"),
		},
		{
			name:    "Description Too Long",
			event:   NewEvent(5, 5, time.Date(2026, 5, 12, 15, 04, 04, 0, time.UTC), "Test", string(make([]rune, 51))),
			wantErr: errors.NewBadRequestError("description parameter is too long, maximum length 50 symbols"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.event.Validate(); err != nil {
				if tt.wantErr == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("Event.Validate() error=%v\nwantErr=%v", err, tt.wantErr)
				}
			}
		})
	}
}
