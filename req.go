package goshikimori

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/heycatch/goshikimori/concat"
)

// Return the date as bytes.
//
// If the context time is exceeded returns -1.
func sendRequest(req *http.Request) ([]byte, int, error) {
	var client = &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return data, resp.StatusCode, err
}

// Normal GET request with User-Agent only.
func NewGetRequestWithCancel(application, search string, number time.Duration) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), number*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, search, nil)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)

	data, status, err := sendRequest(req)
	if err != nil {
		return nil, status, err
	}

	return data, status, nil
}

// For certain GET requests where a Bearer is needed.
func NewGetRequestWithCancelAndBearer(application, accessToken, search string, number time.Duration) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), number*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, search, nil)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concat.Bearer(accessToken))

	data, status, err := sendRequest(req)
	if err != nil {
		return nil, status, err
	}

	return data, status, nil
}

// To work correctly with the POST method,
// make sure that your application has all the necessary permissions.
func NewPostRequestWithCancel(application, accessToken, search string, number time.Duration) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), number*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, search, nil)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concat.Bearer(accessToken))
	req.Header.Set("Content-Type", "application/json")

	data, status, err := sendRequest(req)
	if err != nil {
		return nil, status, err
	}

	return data, status, nil
}

// GraphQL: POST request.
// For GraphQL you only need User-Agent at POST request.
func NewGraphQLPostRequestWithCancel(application, search string, number time.Duration) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), number*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, search, nil)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Set("Content-Type", "application/json")

	data, status, err := sendRequest(req)
	if err != nil {
		return nil, status, err
	}

	return data, status, nil
}

// Reorder: POST request. To work correctly with the POST method,
// make sure that your application has all the necessary permissions.
func NewReorderPostRequestWithCancel(application, accessToken, search string, position int, number time.Duration) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), number*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, search,
		bytes.NewBuffer(concat.DataBuffer(
			[]string{"{\"new_index\": ", "\"", strconv.Itoa(position), "\"", "}"},
		)),
	)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concat.Bearer(accessToken))
	req.Header.Set("Content-Type", "application/json")

	data, status, err := sendRequest(req)
	if err != nil {
		return nil, status, err
	}

	return data, status, nil
}

// Mark order messages: POST request. To work correctly with the POST method,
// make sure that your application has all the necessary permissions.
func NewMarkReadPostRequestWithCancel(application, accessToken, search, ids string, is_read int, number time.Duration) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), number*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, search,
		bytes.NewBuffer(concat.DataBuffer([]string{
			"{\"ids\": ", "\"", ids, "\"", ", ", "\"is_read\": ",
			"\"", strconv.Itoa(is_read), "\"", "}",
		})),
	)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concat.Bearer(accessToken))
	req.Header.Set("Content-Type", "application/json")

	data, status, err := sendRequest(req)
	if err != nil {
		return nil, status, err
	}

	return data, status, nil
}

// Read/Delete all messages: POST request. To work correctly with the POST method,
// make sure that your application has all the necessary permissions.
func NewReadDeleteAllPostRequestWithCancel(application, accessToken, search, name string, number time.Duration) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), number*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, search,
		bytes.NewBuffer(concat.DataCopy(
			33+len(name),
			[]string{"{\"frontend\": ", "\"false\", ", "\"type\": ", "\"", name, "\"", "}"},
		)),
	)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concat.Bearer(accessToken))
	req.Header.Set("Content-Type", "application/json")

	data, status, err := sendRequest(req)
	if err != nil {
		return nil, status, err
	}

	return data, status, nil
}

// Send message: POST request. To work correctly with the POST method,
// make sure that your application has all the necessary permissions.
func NewSendMessagePostRequestWithCancel(application, accessToken, search, body string, from_id, to_id int, number time.Duration) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), number*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, search,
		bytes.NewBuffer(concat.DataBuffer([]string{
			"{\"frontend\": \"false\", \"message\": {\"body\": \"", body,
			"\", \"from_id\": \"", strconv.Itoa(from_id),
			"\", \"kind\": \"Private\", \"to_id\": \"", strconv.Itoa(to_id), "\"}}",
		})),
	)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concat.Bearer(accessToken))
	req.Header.Set("Content-Type", "application/json")

	data, status, err := sendRequest(req)
	if err != nil {
		return nil, status, err
	}

	return data, status, nil
}

// Change message. To work correctly with the PUT method,
// make sure that your application has all the necessary permissions.
func NewChangeMessagePutRequestWithCancel(application, accessToken, search, body string, number time.Duration) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), number*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPut, search,
		bytes.NewBuffer(concat.DataCopy(
			46+len(body),
			[]string{"{\"frontend\": \"false\", \"message\": {\"body\": \"", body, "\"}}"},
		)),
	)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concat.Bearer(accessToken))
	req.Header.Set("Content-Type", "application/json")

	data, status, err := sendRequest(req)
	if err != nil {
		return nil, status, err
	}

	return data, status, nil
}

// Delete message. To work correctly with the DELETE method,
// make sure that your application has all the necessary permissions.
func NewDeleteMessageDeleteRequestWithCancel(application, accessToken, search string, number time.Duration) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), number*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, search, nil)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concat.Bearer(accessToken))
	req.Header.Set("Content-Type", "application/json")

	data, status, err := sendRequest(req)
	if err != nil {
		return nil, status, err
	}

	return data, status, nil
}

// To work correctly with the DELETE method,
// make sure that your application has all the necessary permissions.
func NewDeleteRequestWithCancel(application, accessToken, search string, number time.Duration) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), number*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, search, nil)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concat.Bearer(accessToken))

	data, status, err := sendRequest(req)
	if err != nil {
		return nil, status, err
	}

	return data, status, nil
}
