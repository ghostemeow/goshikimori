package request

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ghostemeow/goshikimori/internal/concatenation"
)

func TestNewGetRequestWithCancel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if header := r.Header.Get("User-Agent"); header != "test-app" {
			t.Errorf("invalid User-Agent: %s", header)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok": true}`))
	}))
	defer server.Close()

	data, status, err := NewGetRequestWithCancel("test-app", server.URL, 5*time.Second)
	if err != nil {
		t.Fatalf("expected nil, got %s", err.Error())
	}
	if status != http.StatusOK {
		t.Errorf("expected 200, got %d", status)
	}
	if !strings.Contains(string(data), `"ok": true`) {
		t.Errorf("invalid response body: %s", string(data))
	}
}

func TestNewGetRequestWithCancelAndBearer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if header := r.Header.Get("User-Agent"); header != "test-app" {
			t.Errorf("invalid User-Agent: %s", header)
		}
		if auth := r.Header.Get("Authorization"); !strings.HasPrefix(auth, "Bearer ") {
			t.Errorf("missing/invalid Authorization: %s", auth)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok": true}`))
	}))
	defer server.Close()

	data, status, err := NewGetRequestWithCancelAndBearer("test-app", "test-token", server.URL, 5*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if status != http.StatusOK {
		t.Errorf("expected 200, got %d", status)
	}
	if !strings.Contains(string(data), `"ok": true`) {
		t.Errorf("invalid response body: %s", string(data))
	}
}

func TestNewPostRequestWithCancel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if header := r.Header.Get("User-Agent"); header != "test-app" {
			t.Errorf("invalid User-Agent: %s", header)
		}
		if auth := r.Header.Get("Authorization"); !strings.HasPrefix(auth, "Bearer ") {
			t.Errorf("missing/invalid Authorization: %s", auth)
		}
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	_, status, err := NewPostRequestWithCancel("test-app", "test-token", server.URL, 5*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if status != http.StatusCreated {
		t.Errorf("expected 201, got %d", status)
	}
}

func TestNewGraphQLPostRequestWithCancel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, received %s", r.Method)
		}
		if header := r.Header.Get("Content-Type"); header != "application/json" {
			t.Errorf("invalid Content-Type: %s", header)
		}
		body, _ := io.ReadAll(r.Body)
		if string(body) != `{"query": "{animes(search: \"initial d\"){id}}"}` {
			t.Errorf("invalid request body: %s", string(body))
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data": {}}`))
	}))
	defer server.Close()

	data, status, err := NewGraphQLPostRequestWithCancel("test-app", server.URL, `{animes(search: "initial d"){id}}`, 5*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if status != 200 {
		t.Errorf("expected 200, got %d", status)
	}
	if !strings.Contains(string(data), `{"data": {}}`) {
		t.Errorf("invalid response body: %s", string(data))
	}
}

func TestNewReorderPostRequestWithCancel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, received %s", r.Method)
		}
		if header := r.Header.Get("User-Agent"); header != "test-app" {
			t.Errorf("invalid User-Agent: %s", header)
		}
		if auth := r.Header.Get("Authorization"); !strings.HasPrefix(auth, "Bearer ") {
			t.Errorf("missing/invalid Authorization: %s", auth)
		}
		body, _ := io.ReadAll(r.Body)
		var payload struct {
			NewIndex string `json:"new_index"`
		}
		if err := json.Unmarshal(body, &payload); err != nil {
			t.Errorf("invalid json: %s", err.Error())
		}
		if payload.NewIndex != "7" {
			t.Errorf("expected new_index=7, got %s", payload.NewIndex)
		}
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	_, status, err := NewReorderPostRequestWithCancel("test-app", "test-token", server.URL, 7, 5*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if status != 201 {
		t.Errorf("expected 201, got %d", status)
	}
}

func TestNewSendMessagePostRequestWithCancel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, received %s", r.Method)
		}
		if header := r.Header.Get("User-Agent"); header != "test-app" {
			t.Errorf("invalid User-Agent: %s", header)
		}
		if auth := r.Header.Get("Authorization"); !strings.HasPrefix(auth, "Bearer ") {
			t.Errorf("missing/invalid Authorization: %s", auth)
		}
		body, _ := io.ReadAll(r.Body)
		var payload struct {
			Frontend string `json:"frontend"`
			Message  struct {
				Body   string `json:"body"`
				FromID string `json:"from_id"`
				Kind   string `json:"kind"`
				ToID   string `json:"to_id"`
			} `json:"message"`
		}
		if err := json.Unmarshal(body, &payload); err != nil {
			t.Errorf("invalid json: %s", err.Error())
		}
		if payload.Message.Body != "test message" {
			t.Errorf("wrong body: %s", payload.Message.Body)
		}
		if payload.Message.Kind != "Private" {
			t.Error("kind should be Private")
		}
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	_, status, err := NewSendMessagePostRequestWithCancel(
		"test-app", "test-token", server.URL,
		"test message", 1, 2, 5*time.Second,
	)
	if err != nil {
		t.Fatal(err)
	}
	if status != 201 {
		t.Errorf("expected 201, got %d", status)
	}
}

func TestNewDeleteRequestWithCancel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if header := r.Header.Get("User-Agent"); header != "test-app" {
			t.Errorf("invalid User-Agent: %s", header)
		}
		if auth := r.Header.Get("Authorization"); !strings.HasPrefix(auth, "Bearer ") {
			t.Errorf("missing/invalid Authorization: %s", auth)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	data, status, err := NewDeleteRequestWithCancel("test-app", "test-token", server.URL, 5*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if status != 204 {
		t.Errorf("expected 204, got %d", status)
	}
	if len(data) != 0 {
		t.Error("expected empty body on 204")
	}
}

func TestNewMarkReadPostRequestWithCancel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if header := r.Header.Get("User-Agent"); header != "test-app" {
			t.Errorf("invalid User-Agent: %s", header)
		}
		if auth := r.Header.Get("Authorization"); !strings.HasPrefix(auth, "Bearer ") {
			t.Errorf("missing/invalid Authorization: %s", auth)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()

	data, status, err := NewMarkReadPostRequestWithCancel(
		"test-app", "test-token", server.URL,
		concatenation.IdsToString([]int{123, 456, 789}), 1337, 5*time.Second,
	)
	if err != nil {
		t.Fatal(err)
	}
	if status != 200 {
		t.Errorf("expected 200, got %d", status)
	}
	if !strings.Contains(string(data), `"success": true`) {
		t.Errorf("invalid response body: %s", string(data))
	}
}

func TestNewReadDeleteAllPostRequestWithCancel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if header := r.Header.Get("User-Agent"); header != "test-app" {
			t.Errorf("invalid User-Agent: %s", header)
		}
		if auth := r.Header.Get("Authorization"); !strings.HasPrefix(auth, "Bearer ") {
			t.Errorf("missing/invalid Authorization: %s", auth)
		}
		body, _ := io.ReadAll(r.Body)
		var payload struct {
			Frontend string `json:"frontend"`
			Type     string `json:"type"`
		}
		if err := json.Unmarshal(body, &payload); err != nil {
			t.Errorf("invalid json: %s", err.Error())
		}
		if payload.Frontend != "false" {
			t.Errorf("expected frontend=\"false\", got %s", payload.Frontend)
		}
		if payload.Type != "inbox" {
			t.Errorf("expected type=\"inbox\", got %s", payload.Type)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	_, status, err := NewReadDeleteAllPostRequestWithCancel(
		"test-app", "test-token", server.URL,
		"inbox", 5*time.Second,
	)
	if err != nil {
		t.Fatal(err)
	}
	if status != 200 {
		t.Errorf("expected 200, got %d", status)
	}
}

func TestNewChangeMessagePutRequestWithCancel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if header := r.Header.Get("User-Agent"); header != "test-app" {
			t.Errorf("invalid User-Agent: %s", header)
		}
		if auth := r.Header.Get("Authorization"); !strings.HasPrefix(auth, "Bearer ") {
			t.Errorf("missing/invalid Authorization: %s", auth)
		}
		body, _ := io.ReadAll(r.Body)
		var payload struct {
			Frontend string `json:"frontend"`
			Message  struct {
				Body string `json:"body"`
			} `json:"message"`
		}
		if err := json.Unmarshal(body, &payload); err != nil {
			t.Errorf("invalid json: %s", err.Error())
		}
		if payload.Frontend != "false" {
			t.Errorf("expected frontend=\"false\", got %s", payload.Frontend)
		}
		if payload.Message.Body != "updated message text her" {
			t.Errorf("wrong message body: %s", payload.Message.Body)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	_, status, err := NewChangeMessagePutRequestWithCancel(
		"test-app", "test-token", server.URL,
		"updated message text her", 5*time.Second,
	)
	if err != nil {
		t.Fatal(err)
	}
	if status != 200 {
		t.Errorf("expected 200, got %d", status)
	}
}

func TestNewDeleteMessageDeleteRequestWithCancel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if header := r.Header.Get("User-Agent"); header != "test-app" {
			t.Errorf("invalid User-Agent: %s", header)
		}
		if auth := r.Header.Get("Authorization"); !strings.HasPrefix(auth, "Bearer ") {
			t.Errorf("missing/invalid Authorization: %s", auth)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	data, status, err := NewDeleteMessageDeleteRequestWithCancel("test-app", "test-token", server.URL, 5*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if status != 204 {
		t.Errorf("expected 204, got %d", status)
	}
	if len(data) != 0 {
		t.Error("expected empty body on 204")
	}
}

func TestNewCreateTopicPostRequestWithCancel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if header := r.Header.Get("User-Agent"); header != "test-app" {
			t.Errorf("invalid User-Agent: %s", header)
		}
		if auth := r.Header.Get("Authorization"); !strings.HasPrefix(auth, "Bearer ") {
			t.Errorf("missing/invalid Authorization: %s", auth)
		}
		if header := r.Header.Get("Content-Type"); header != "application/json" {
			t.Errorf("invalid Content-Type: %s", header)
		}
		body, _ := io.ReadAll(r.Body)
		var payload struct {
			Topic struct {
				Body     string `json:"body"`
				ForumID  int    `json:"forum_id"`
				LinkedID int    `json:"linked_id,omitempty"`
				Title    string `json:"title"`
				Type     string `json:"type"`
				UserID   int    `json:"user_id"`
			} `json:"topic"`
		}
		if err := json.Unmarshal(body, &payload); err != nil {
			t.Errorf("invalid json: %s", err.Error())
		}
		if payload.Topic.Body != "test body" {
			t.Errorf("wrong body: %s", payload.Topic.Body)
		}
		if payload.Topic.ForumID != 1 {
			t.Errorf("expected forum_id=1, got %d", payload.Topic.ForumID)
		}
		if payload.Topic.Title != "test title" {
			t.Errorf("wrong title: %s", payload.Topic.Title)
		}
		if payload.Topic.Type != "Topic" {
			t.Errorf("expected type=Topic, got %s", payload.Topic.Type)
		}
		if payload.Topic.UserID != 23456791 {
			t.Errorf("expected user_id=23456791, got %d", payload.Topic.UserID)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"id": 270126, "topic_title": "test title", "body": "test body"}`))
	}))
	defer server.Close()

	data, status, err := NewCreateTopicPostRequestWithCancel(
		"test-app", "test-token", server.URL,
		[]byte(`{"topic":{"body":"test body","forum_id":1,"title":"test title","type":"Topic","user_id":23456791}}`),
		5*time.Second,
	)
	if err != nil {
		t.Fatal(err)
	}
	if status != 201 {
		t.Errorf("expected 201, got %d", status)
	}
	if !strings.Contains(string(data), `"id": 270126`) {
		t.Errorf("invalid response body: %s", string(data))
	}
}

func TestNewUpdateTopicPatchRequestWithCancel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("expected PATCH, got %s", r.Method)
		}
		if header := r.Header.Get("User-Agent"); header != "test-app" {
			t.Errorf("invalid User-Agent: %s", header)
		}
		if auth := r.Header.Get("Authorization"); !strings.HasPrefix(auth, "Bearer ") {
			t.Errorf("missing/invalid Authorization: %s", auth)
		}
		if header := r.Header.Get("Content-Type"); header != "application/json" {
			t.Errorf("invalid Content-Type: %s", header)
		}
		body, _ := io.ReadAll(r.Body)
		var payload struct {
			Topic struct {
				Body string `json:"body"`
			} `json:"topic"`
		}
		if err := json.Unmarshal(body, &payload); err != nil {
			t.Errorf("invalid json: %s", err.Error())
		}
		if payload.Topic.Body != "blablalbla" {
			t.Errorf("wrong body: %s", payload.Topic.Body)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id": 270122, "body": "blablalbla"}`))
	}))
	defer server.Close()

	data, status, err := NewUpdateTopicPatchRequestWithCancel(
		"test-app", "test-token", server.URL,
		[]byte(`{"topic":{"body":"blablalbla"}}`),
		5*time.Second,
	)
	if err != nil {
		t.Fatal(err)
	}
	if status != 200 {
		t.Errorf("expected 200, got %d", status)
	}
	if !strings.Contains(string(data), `"id": 270122`) {
		t.Errorf("invalid response body: %s", string(data))
	}
}
