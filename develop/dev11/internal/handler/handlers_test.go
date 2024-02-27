package handler

import (
	"develop/dev11/internal/data"
	"develop/dev11/internal/models"
	"develop/dev11/internal/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHandlerCreateEvent(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		method     string
		body       string
		want       string
		wantStatus int
	}{
		{
			name:       "OK",
			url:        "http://localhost:8080/create_event",
			method:     "POST",
			body:       "user_id=5&date=2026-05-12 15:04:05&title=Test&description=Test",
			want:       "{\"result\":{\"eventID\":0}}\n",
			wantStatus: http.StatusCreated,
		},
		{
			name:       "Not Found Path",
			url:        "http://localhost:8080/wrong_path",
			method:     "GET",
			body:       "user_id=5&date=2026-05-12 15:04:05&title=Test&description=Test",
			want:       "{\"error\":\"path /wrong_path not found\"}\n",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "Method",
			url:        "http://localhost:8080/create_event",
			method:     "GET",
			body:       "user_id=5&date=2026-05-12 15:04:05&title=Test&description=Test",
			want:       "{\"error\":\"method not allowed: bad method GET, method must be POST\"}\n",
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:       "No UserID",
			url:        "http://localhost:8080/create_event",
			method:     "POST",
			body:       "date=2026-05-12 15:04:05&title=Test&description=Test",
			want:       "{\"error\":\"bad request: empty parameter: user_id\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Invalid UserID",
			url:        "http://localhost:8080/create_event",
			method:     "POST",
			body:       "user_id=five&&date=2026-05-12 15:04:05&title=Test&description=Test",
			want:       "{\"error\":\"bad request: invalid user_id: use only numbers\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "No Title",
			url:        "http://localhost:8080/create_event",
			method:     "POST",
			body:       "user_id=5&date=2026-05-12 15:04:05&description=Test",
			want:       "{\"error\":\"bad request: empty parameter: title\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "No Description",
			url:        "http://localhost:8080/create_event",
			method:     "POST",
			body:       "user_id=5&date=2026-05-12 15:04:05&title=Test",
			want:       "{\"result\":{\"eventID\":0}}\n",
			wantStatus: http.StatusCreated,
		},
		{
			name:       "No Date",
			url:        "http://localhost:8080/create_event",
			method:     "POST",
			body:       "user_id=5&title=Test&description=Test",
			want:       "{\"error\":\"bad request: empty parameter: date\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Format Date",
			url:        "http://localhost:8080/create_event",
			method:     "POST",
			body:       "user_id=5&date=2026.05.12 15:04:05&title=Test&description=Test",
			want:       "{\"error\":\"bad request: invalid date format: correct format 2006-01-02 15:04:05\"}\n",
			wantStatus: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, err := http.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
			if err != nil {
				t.Error(err)
				return
			}
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			responseRecorder := httptest.NewRecorder()
			handler := New(service.New(data.New())).InitRouter()
			handler.ServeHTTP(responseRecorder, request)
			if responseRecorder.Code != tt.wantStatus {
				t.Errorf("status: got %v want %v", responseRecorder.Code, tt.wantStatus)
			}
			if responseRecorder.Body.String() != tt.want {
				t.Errorf("result: got %v want %v", responseRecorder.Body.String(), tt.want)
			}
		})
	}
}

func TestHandlerUpdateEvent(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		method     string
		body       string
		events     []*models.Event
		want       string
		wantStatus int
	}{
		{
			name:       "OK Update Title",
			url:        "http://localhost:8080/update_event",
			method:     "POST",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 12, 15, 04, 04, 0, time.UTC), "Test", "Test")},
			body:       "user_id=5&id=0&date=2026-05-12 15:04:05&title=UpdateTitle&description=Test",
			want:       "{\"result\":\"OK\"}\n",
			wantStatus: http.StatusOK,
		},
		{
			name:       "Not Found Path",
			url:        "http://localhost:8080/wrong_path",
			method:     "GET",
			body:       "user_id=5&date=2026-05-12 15:04:05&title=Test&description=Test",
			want:       "{\"error\":\"path /wrong_path not found\"}\n",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "OK Update Description",
			url:        "http://localhost:8080/update_event",
			method:     "POST",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 12, 15, 04, 04, 0, time.UTC), "Test", "Test")},
			body:       "user_id=5&id=0&date=2026-05-12 15:04:05&title=Title&description=UpdateTest",
			want:       "{\"result\":\"OK\"}\n",
			wantStatus: http.StatusOK,
		},
		{
			name:       "OK Update Date",
			url:        "http://localhost:8080/update_event",
			method:     "POST",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 12, 15, 04, 04, 0, time.UTC), "Test", "Test")},
			body:       "user_id=5&id=0&date=2027-07-07 15:04:05&title=Title&description=Test",
			want:       "{\"result\":\"OK\"}\n",
			wantStatus: http.StatusOK,
		},
		{
			name:       "UserID Not Found",
			url:        "http://localhost:8080/update_event",
			method:     "POST",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 12, 15, 04, 04, 0, time.UTC), "Test", "Test")},
			body:       "user_id=1&id=0&date=2027-07-07 15:04:05&title=Title&description=Test",
			want:       "{\"error\":\"user_id 1 not found\"}\n",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "EventID Not Found",
			url:        "http://localhost:8080/update_event",
			method:     "POST",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 12, 15, 04, 04, 0, time.UTC), "Test", "Test")},
			body:       "user_id=5&id=10&date=2027-07-07 15:04:05&title=Title&description=Test",
			want:       "{\"error\":\"event id 10 not found\"}\n",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "Method",
			url:        "http://localhost:8080/update_event",
			method:     "GET",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 12, 15, 04, 04, 0, time.UTC), "Test", "Test")},
			body:       "user_id=5&id=0date=2026-05-12 15:04:05&title=Test&description=Test",
			want:       "{\"error\":\"method not allowed: bad method GET, method must be POST\"}\n",
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:       "No UserID",
			url:        "http://localhost:8080/update_event",
			method:     "POST",
			body:       "id=10&date=2026-05-12 15:04:05&title=Test&description=Test",
			want:       "{\"error\":\"bad request: empty parameter: user_id\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "No Event ID",
			url:        "http://localhost:8080/update_event",
			method:     "POST",
			body:       "user_id=5date=2026-05-12 15:04:05&title=Test&description=Test",
			want:       "{\"error\":\"bad request: empty parameter: id\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Invalid UserID",
			url:        "http://localhost:8080/update_event",
			method:     "POST",
			body:       "user_id=five&id=0&date=2026-05-12 15:04:05&title=Test&description=Test",
			want:       "{\"error\":\"bad request: invalid user_id: use only numbers\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Invalid Event ID",
			url:        "http://localhost:8080/update_event",
			method:     "POST",
			body:       "user_id=5&id=zero&date=2026-05-12 15:04:05&title=Test&description=Test",
			want:       "{\"error\":\"bad request: invalid id: use only numbers\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "No Title",
			url:        "http://localhost:8080/update_event",
			method:     "POST",
			body:       "user_id=5&id=0&date=2026-05-12 15:04:05&description=Test",
			want:       "{\"error\":\"bad request: empty parameter: title\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "No Description",
			url:        "http://localhost:8080/update_event",
			method:     "POST",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 12, 15, 04, 04, 0, time.UTC), "Test", "Test")},
			body:       "user_id=5&id=0&date=2026-05-12 15:04:05&title=TestUpdate",
			want:       "{\"result\":\"OK\"}\n",
			wantStatus: http.StatusOK,
		},
		{
			name:       "No Date",
			url:        "http://localhost:8080/update_event",
			method:     "POST",
			body:       "user_id=5&id=0&title=Test&description=Test",
			want:       "{\"error\":\"bad request: empty parameter: date\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Format Date",
			url:        "http://localhost:8080/update_event",
			method:     "POST",
			body:       "user_id=5&id=0&date=2026.05.12 15:04:05&title=Test&description=Test",
			want:       "{\"error\":\"bad request: invalid date format: correct format 2006-01-02 15:04:05\"}\n",
			wantStatus: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, err := http.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
			if err != nil {
				t.Error(err)
				return
			}
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			responseRecorder := httptest.NewRecorder()
			service := service.New(data.New())
			for _, event := range tt.events {
				service.Create(event)
			}
			handler := New(service).InitRouter()
			handler.ServeHTTP(responseRecorder, request)
			if responseRecorder.Code != tt.wantStatus {
				t.Errorf("status: got %v want %v", responseRecorder.Code, tt.wantStatus)
			}
			if responseRecorder.Body.String() != tt.want {
				t.Errorf("result: got %v want %v", responseRecorder.Body.String(), tt.want)
			}
		})
	}
}

func TestHandlerDeleteEvent(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		method     string
		body       string
		events     []*models.Event
		want       string
		wantStatus int
	}{
		{
			name:       "OK Delete",
			url:        "http://localhost:8080/delete_event",
			method:     "POST",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 12, 15, 04, 04, 0, time.UTC), "Test", "Test")},
			body:       "user_id=5&id=0",
			want:       "{\"result\":\"OK\"}\n",
			wantStatus: http.StatusOK,
		},
		{
			name:       "Not Found Path",
			url:        "http://localhost:8080/wrong_path",
			method:     "GET",
			body:       "user_id=5&date=2026-05-12 15:04:05&title=Test&description=Test",
			want:       "{\"error\":\"path /wrong_path not found\"}\n",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "UserID Not Found",
			url:        "http://localhost:8080/delete_event",
			method:     "POST",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 12, 15, 04, 04, 0, time.UTC), "Test", "Test")},
			body:       "user_id=1&id=0",
			want:       "{\"error\":\"user_id 1 not found\"}\n",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "EventID Not Found",
			url:        "http://localhost:8080/delete_event",
			method:     "POST",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 12, 15, 04, 04, 0, time.UTC), "Test", "Test")},
			body:       "user_id=5&id=10",
			want:       "{\"error\":\"event id 10 not found\"}\n",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "Method",
			url:        "http://localhost:8080/delete_event",
			method:     "GET",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 12, 15, 04, 04, 0, time.UTC), "Test", "Test")},
			body:       "user_id=5&id=0",
			want:       "{\"error\":\"method not allowed: bad method GET, method must be POST\"}\n",
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:       "No UserID",
			url:        "http://localhost:8080/delete_event",
			method:     "POST",
			body:       "id=0",
			want:       "{\"error\":\"bad request: empty parameter: user_id\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "No Event ID",
			url:        "http://localhost:8080/delete_event",
			method:     "POST",
			body:       "user_id=5",
			want:       "{\"error\":\"bad request: empty parameter: id\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Invalid UserID",
			url:        "http://localhost:8080/delete_event",
			method:     "POST",
			body:       "user_id=five&id=0",
			want:       "{\"error\":\"bad request: invalid user_id: use only numbers\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Invalid Event ID",
			url:        "http://localhost:8080/delete_event",
			method:     "POST",
			body:       "user_id=5&id=zero",
			want:       "{\"error\":\"bad request: invalid id: use only numbers\"}\n",
			wantStatus: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, err := http.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
			if err != nil {
				t.Error(err)
				return
			}
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			responseRecorder := httptest.NewRecorder()
			service := service.New(data.New())
			for _, event := range tt.events {
				service.Create(event)
			}
			handler := New(service).InitRouter()
			handler.ServeHTTP(responseRecorder, request)
			if responseRecorder.Code != tt.wantStatus {
				t.Errorf("status: got %v want %v", responseRecorder.Code, tt.wantStatus)
			}
			if responseRecorder.Body.String() != tt.want {
				t.Errorf("result: got %v want %v", responseRecorder.Body.String(), tt.want)
			}
		})
	}
}

func TestHandlerEventsForDay(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		url        string
		events     []*models.Event
		want       string
		wantStatus int
	}{
		{
			name:   "OK Event For Day",
			method: "GET",
			events: []*models.Event{
				models.NewEvent(5, 0, time.Date(2026, 5, 12, 14, 04, 04, 0, time.UTC), "Test1", "Test1"),
				models.NewEvent(5, 0, time.Date(2026, 5, 18, 14, 04, 04, 0, time.UTC), "Test2", "Test2"),
				models.NewEvent(5, 0, time.Date(2026, 5, 30, 14, 04, 04, 0, time.UTC), "Test3", "Test3"),
				models.NewEvent(5, 0, time.Date(2026, 6, 20, 14, 04, 04, 0, time.UTC), "Test4", "Test4"),
			},
			url:        "http://localhost:8080/events_for_day?user_id=5&date=2026-05-12",
			want:       "{\"result\":[{\"user_id\":5,\"id\":0,\"date\":\"2026-05-12T14:04:04Z\",\"title\":\"Test1\",\"description\":\"Test1\"}]}\n",
			wantStatus: http.StatusOK,
		},
		{
			name:   "Not Found Path",
			method: "GET",
			url:    "http://localhost:8080/wrong_path",
			events: []*models.Event{
				models.NewEvent(5, 0, time.Date(2026, 5, 12, 14, 04, 04, 0, time.UTC), "Test1", "Test1"),
				models.NewEvent(5, 0, time.Date(2026, 5, 18, 14, 04, 04, 0, time.UTC), "Test2", "Test2"),
				models.NewEvent(5, 0, time.Date(2026, 5, 30, 14, 04, 04, 0, time.UTC), "Test3", "Test3"),
				models.NewEvent(5, 0, time.Date(2026, 6, 20, 14, 04, 04, 0, time.UTC), "Test4", "Test4"),
			},
			want:       "{\"error\":\"path /wrong_path not found\"}\n",
			wantStatus: http.StatusNotFound,
		},
		{
			name:   "Empty Event For Day",
			method: "GET",
			events: []*models.Event{
				models.NewEvent(5, 0, time.Date(2026, 5, 12, 14, 04, 04, 0, time.UTC), "Test1", "Test1"),
				models.NewEvent(5, 1, time.Date(2026, 5, 18, 14, 04, 04, 0, time.UTC), "Test2", "Test2"),
				models.NewEvent(5, 2, time.Date(2026, 5, 30, 14, 04, 04, 0, time.UTC), "Test3", "Test3"),
				models.NewEvent(5, 3, time.Date(2026, 6, 20, 14, 04, 04, 0, time.UTC), "Test4", "Test4"),
			},
			url:        "http://localhost:8080/events_for_day?user_id=5&date=2026-05-13",
			want:       "{\"result\":[]}\n",
			wantStatus: http.StatusOK,
		},

		{
			name:       "UserID Not Found",
			method:     "GET",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 13, 14, 04, 04, 0, time.UTC), "Test1", "Test1")},
			url:        "http://localhost:8080/events_for_day?user_id=1&date=2026-05-13",
			want:       "{\"error\":\"user_id 1 not found\"}\n",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "Method",
			method:     "POST",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 13, 14, 04, 04, 0, time.UTC), "Test1", "Test1")},
			url:        "http://localhost:8080/events_for_day?user_id=5&date=2026-05-13",
			want:       "{\"error\":\"method not allowed: bad method POST, method must be GET\"}\n",
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:       "Invalid UserID",
			method:     "GET",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 13, 14, 04, 04, 0, time.UTC), "Test1", "Test1")},
			url:        "http://localhost:8080/events_for_day?user_id=one&date=2026-05-13",
			want:       "{\"error\":\"bad request: invalid user_id: use only numbers\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "No UserID",
			method:     "GET",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 13, 14, 04, 04, 0, time.UTC), "Test1", "Test1")},
			url:        "http://localhost:8080/events_for_day?date=2026-05-13",
			want:       "{\"error\":\"bad request: empty parameter: user_id\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "No Date",
			method:     "GET",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 13, 14, 04, 04, 0, time.UTC), "Test1", "Test1")},
			url:        "http://localhost:8080/events_for_day?user_id=5",
			want:       "{\"error\":\"bad request: empty parameter: date\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Format Date",
			method:     "GET",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 13, 14, 04, 04, 0, time.UTC), "Test1", "Test1")},
			url:        "http://localhost:8080/events_for_day?user_id=5&date=2026.05.13",
			want:       "{\"error\":\"bad request: invalid date format: correct format 2006-01-02\"}\n",
			wantStatus: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Error(err)
				return
			}
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			responseRecorder := httptest.NewRecorder()
			service := service.New(data.New())
			for _, event := range tt.events {
				service.Create(event)
			}
			handler := New(service).InitRouter()
			handler.ServeHTTP(responseRecorder, request)
			if responseRecorder.Code != tt.wantStatus {
				t.Errorf("status: got %v want %v", responseRecorder.Code, tt.wantStatus)
			}
			if responseRecorder.Body.String() != tt.want {
				t.Errorf("result: got %v want %v", responseRecorder.Body.String(), tt.want)
			}
		})
	}
}

func TestHandlerEventsForWeek(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		url        string
		events     []*models.Event
		want       string
		wantStatus int
	}{
		{
			name:   "OK Event For Week",
			method: "GET",
			events: []*models.Event{
				models.NewEvent(5, 0, time.Date(2026, 5, 12, 14, 04, 04, 0, time.UTC), "Test1", "Test1"),
				models.NewEvent(5, 0, time.Date(2026, 5, 18, 14, 04, 04, 0, time.UTC), "Test2", "Test2"),
				models.NewEvent(5, 0, time.Date(2026, 5, 30, 14, 04, 04, 0, time.UTC), "Test3", "Test3"),
				models.NewEvent(5, 0, time.Date(2026, 6, 20, 14, 04, 04, 0, time.UTC), "Test4", "Test4"),
			},
			url: "http://localhost:8080/events_for_week?user_id=5&date=2026-05-12",
			want: "{\"result\":[" +
				"{\"user_id\":5,\"id\":0,\"date\":\"2026-05-12T14:04:04Z\",\"title\":\"Test1\",\"description\":\"Test1\"}," +
				"{\"user_id\":5,\"id\":1,\"date\":\"2026-05-18T14:04:04Z\",\"title\":\"Test2\",\"description\":\"Test2\"}" +
				"]}\n",
			wantStatus: http.StatusOK,
		},
		{
			name:   "Not Found Path",
			method: "GET",
			url:    "http://localhost:8080/wrong_path",
			events: []*models.Event{
				models.NewEvent(5, 0, time.Date(2026, 5, 12, 14, 04, 04, 0, time.UTC), "Test1", "Test1"),
				models.NewEvent(5, 0, time.Date(2026, 5, 18, 14, 04, 04, 0, time.UTC), "Test2", "Test2"),
				models.NewEvent(5, 0, time.Date(2026, 5, 30, 14, 04, 04, 0, time.UTC), "Test3", "Test3"),
				models.NewEvent(5, 0, time.Date(2026, 6, 20, 14, 04, 04, 0, time.UTC), "Test4", "Test4"),
			},
			want:       "{\"error\":\"path /wrong_path not found\"}\n",
			wantStatus: http.StatusNotFound,
		},
		{
			name:   "Empty Event For Week",
			method: "GET",
			events: []*models.Event{
				models.NewEvent(5, 0, time.Date(2026, 5, 12, 14, 04, 04, 0, time.UTC), "Test1", "Test1"),
				models.NewEvent(5, 1, time.Date(2026, 5, 18, 14, 04, 04, 0, time.UTC), "Test2", "Test2"),
				models.NewEvent(5, 2, time.Date(2026, 5, 30, 14, 04, 04, 0, time.UTC), "Test3", "Test3"),
				models.NewEvent(5, 3, time.Date(2026, 6, 20, 14, 04, 04, 0, time.UTC), "Test4", "Test4"),
			},
			url:        "http://localhost:8080/events_for_week?user_id=5&date=2026-07-13",
			want:       "{\"result\":[]}\n",
			wantStatus: http.StatusOK,
		},
		{
			name:       "UserID Not Found",
			method:     "GET",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 13, 14, 04, 04, 0, time.UTC), "Test1", "Test1")},
			url:        "http://localhost:8080/events_for_week?user_id=1&date=2026-05-13",
			want:       "{\"error\":\"user_id 1 not found\"}\n",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "Method",
			method:     "POST",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 13, 14, 04, 04, 0, time.UTC), "Test1", "Test1")},
			url:        "http://localhost:8080/events_for_week?user_id=5&date=2026-05-13",
			want:       "{\"error\":\"method not allowed: bad method POST, method must be GET\"}\n",
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:       "Invalid UserID",
			method:     "GET",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 13, 14, 04, 04, 0, time.UTC), "Test1", "Test1")},
			url:        "http://localhost:8080/events_for_week?user_id=one&date=2026-05-13",
			want:       "{\"error\":\"bad request: invalid user_id: use only numbers\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "No UserID",
			method:     "GET",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 13, 14, 04, 04, 0, time.UTC), "Test1", "Test1")},
			url:        "http://localhost:8080/events_for_week?date=2026-05-13",
			want:       "{\"error\":\"bad request: empty parameter: user_id\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "No Date",
			method:     "GET",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 13, 14, 04, 04, 0, time.UTC), "Test1", "Test1")},
			url:        "http://localhost:8080/events_for_week?user_id=5",
			want:       "{\"error\":\"bad request: empty parameter: date\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Format Date",
			method:     "GET",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 13, 14, 04, 04, 0, time.UTC), "Test1", "Test1")},
			url:        "http://localhost:8080/events_for_week?user_id=5&date=2026.05.13",
			want:       "{\"error\":\"bad request: invalid date format: correct format 2006-01-02\"}\n",
			wantStatus: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Error(err)
				return
			}
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			responseRecorder := httptest.NewRecorder()
			service := service.New(data.New())
			for _, event := range tt.events {
				service.Create(event)
			}
			handler := New(service).InitRouter()
			handler.ServeHTTP(responseRecorder, request)
			if responseRecorder.Code != tt.wantStatus {
				t.Errorf("status: got %v want %v", responseRecorder.Code, tt.wantStatus)
			}
			if responseRecorder.Body.String() != tt.want {
				t.Errorf("result: got %v want %v", responseRecorder.Body.String(), tt.want)
			}
		})
	}
}

func TestHandlerEventsForMonth(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		url        string
		events     []*models.Event
		want       string
		wantStatus int
	}{
		{
			name:   "OK Event For Month",
			method: "GET",
			events: []*models.Event{
				models.NewEvent(5, 0, time.Date(2026, 5, 12, 14, 04, 04, 0, time.UTC), "Test1", "Test1"),
				models.NewEvent(5, 0, time.Date(2026, 5, 18, 14, 04, 04, 0, time.UTC), "Test2", "Test2"),
				models.NewEvent(5, 0, time.Date(2026, 5, 30, 14, 04, 04, 0, time.UTC), "Test3", "Test3"),
				models.NewEvent(5, 0, time.Date(2026, 6, 20, 14, 04, 04, 0, time.UTC), "Test4", "Test4"),
			},
			url: "http://localhost:8080/events_for_month?user_id=5&date=2026-05-12",
			want: "{\"result\":[" +
				"{\"user_id\":5,\"id\":0,\"date\":\"2026-05-12T14:04:04Z\",\"title\":\"Test1\",\"description\":\"Test1\"}," +
				"{\"user_id\":5,\"id\":1,\"date\":\"2026-05-18T14:04:04Z\",\"title\":\"Test2\",\"description\":\"Test2\"}," +
				"{\"user_id\":5,\"id\":2,\"date\":\"2026-05-30T14:04:04Z\",\"title\":\"Test3\",\"description\":\"Test3\"}" +
				"]}\n",
			wantStatus: http.StatusOK,
		},
		{
			name:   "Not Found Path",
			method: "GET",
			url:    "http://localhost:8080/wrong_path",
			events: []*models.Event{
				models.NewEvent(5, 0, time.Date(2026, 5, 12, 14, 04, 04, 0, time.UTC), "Test1", "Test1"),
				models.NewEvent(5, 0, time.Date(2026, 5, 18, 14, 04, 04, 0, time.UTC), "Test2", "Test2"),
				models.NewEvent(5, 0, time.Date(2026, 5, 30, 14, 04, 04, 0, time.UTC), "Test3", "Test3"),
				models.NewEvent(5, 0, time.Date(2026, 6, 20, 14, 04, 04, 0, time.UTC), "Test4", "Test4"),
			},
			want:       "{\"error\":\"path /wrong_path not found\"}\n",
			wantStatus: http.StatusNotFound,
		},
		{
			name:   "Empty Event For Month",
			method: "GET",
			events: []*models.Event{
				models.NewEvent(5, 0, time.Date(2026, 5, 12, 14, 04, 04, 0, time.UTC), "Test1", "Test1"),
				models.NewEvent(5, 1, time.Date(2026, 5, 18, 14, 04, 04, 0, time.UTC), "Test2", "Test2"),
				models.NewEvent(5, 2, time.Date(2026, 5, 30, 14, 04, 04, 0, time.UTC), "Test3", "Test3"),
				models.NewEvent(5, 3, time.Date(2026, 6, 20, 14, 04, 04, 0, time.UTC), "Test4", "Test4"),
			},
			url:        "http://localhost:8080/events_for_month?user_id=5&date=2026-07-13",
			want:       "{\"result\":[]}\n",
			wantStatus: http.StatusOK,
		},
		{
			name:       "UserID Not Found",
			method:     "GET",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 13, 14, 04, 04, 0, time.UTC), "Test1", "Test1")},
			url:        "http://localhost:8080/events_for_month?user_id=1&date=2026-05-13",
			want:       "{\"error\":\"user_id 1 not found\"}\n",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "Method",
			method:     "POST",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 13, 14, 04, 04, 0, time.UTC), "Test1", "Test1")},
			url:        "http://localhost:8080/events_for_month?user_id=5&date=2026-05-13",
			want:       "{\"error\":\"method not allowed: bad method POST, method must be GET\"}\n",
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:       "Invalid UserID",
			method:     "GET",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 13, 14, 04, 04, 0, time.UTC), "Test1", "Test1")},
			url:        "http://localhost:8080/events_for_month?user_id=one&date=2026-05-13",
			want:       "{\"error\":\"bad request: invalid user_id: use only numbers\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "No UserID",
			method:     "GET",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 13, 14, 04, 04, 0, time.UTC), "Test1", "Test1")},
			url:        "http://localhost:8080/events_for_month?date=2026-05-13",
			want:       "{\"error\":\"bad request: empty parameter: user_id\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "No Date",
			method:     "GET",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 13, 14, 04, 04, 0, time.UTC), "Test1", "Test1")},
			url:        "http://localhost:8080/events_for_month?user_id=5",
			want:       "{\"error\":\"bad request: empty parameter: date\"}\n",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Format Date",
			method:     "GET",
			events:     []*models.Event{models.NewEvent(5, 0, time.Date(2026, 5, 13, 14, 04, 04, 0, time.UTC), "Test1", "Test1")},
			url:        "http://localhost:8080/events_for_month?user_id=5&date=2026.05.13",
			want:       "{\"error\":\"bad request: invalid date format: correct format 2006-01-02\"}\n",
			wantStatus: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Error(err)
				return
			}
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			responseRecorder := httptest.NewRecorder()
			service := service.New(data.New())
			for _, event := range tt.events {
				service.Create(event)
			}
			handler := New(service).InitRouter()
			handler.ServeHTTP(responseRecorder, request)
			if responseRecorder.Code != tt.wantStatus {
				t.Errorf("status: got %v want %v", responseRecorder.Code, tt.wantStatus)
			}
			if responseRecorder.Body.String() != tt.want {
				t.Errorf("result: got %v want %v", responseRecorder.Body.String(), tt.want)
			}
		})
	}
}
