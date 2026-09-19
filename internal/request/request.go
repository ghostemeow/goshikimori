package request

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/ghostemeow/goshikimori/internal/concatenation"
)

// Return the date as bytes.
//
// If the context time is exceeded returns -1.
//
// A non-2xx status code is returned as an error along with the status.
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

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		msg := string(data)
		if len(msg) > 512 {
			msg = msg[:512]
		}
		return data, resp.StatusCode, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, msg)
	}

	return data, resp.StatusCode, nil
}

// Normal GET request with User-Agent only.
func NewGetRequestWithCancel(application, search string, number time.Duration) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), number)
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
	ctx, cancel := context.WithTimeout(context.Background(), number)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, search, nil)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concatenation.Bearer(accessToken))

	data, status, err := sendRequest(req)
	if err != nil {
		return nil, status, err
	}

	return data, status, nil
}

// To work correctly with the POST method,
// make sure that your application has all the necessary permissions.
func NewPostRequestWithCancel(application, accessToken, search string, number time.Duration) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), number)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, search, nil)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concatenation.Bearer(accessToken))
	req.Header.Set("Content-Type", "application/json")

	data, status, err := sendRequest(req)
	if err != nil {
		return nil, status, err
	}

	return data, status, nil
}

// GraphQL: POST request with the query in the JSON body.
// For GraphQL you only need User-Agent at POST request.
func NewGraphQLPostRequestWithCancel(application, search, query string, number time.Duration) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), number)
	defer cancel()

	body := concatenation.DataBuffer([]string{
		"{\"query\": ", strconv.Quote(query), "}",
	})

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, search, bytes.NewBuffer(body))
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
	ctx, cancel := context.WithTimeout(context.Background(), number)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, search,
		bytes.NewBuffer(concatenation.DataBuffer(
			[]string{"{\"new_index\": ", "\"", strconv.Itoa(position), "\"", "}"},
		)),
	)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concatenation.Bearer(accessToken))
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
	ctx, cancel := context.WithTimeout(context.Background(), number)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, search,
		bytes.NewBuffer(concatenation.DataBuffer([]string{
			"{\"ids\": ", "\"", ids, "\"", ", ", "\"is_read\": ",
			"\"", strconv.Itoa(is_read), "\"", "}",
		})),
	)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concatenation.Bearer(accessToken))
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
	ctx, cancel := context.WithTimeout(context.Background(), number)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, search,
		bytes.NewBuffer(concatenation.DataCopy(
			33+len(name),
			[]string{"{\"frontend\": ", "\"false\", ", "\"type\": ", "\"", name, "\"", "}"},
		)),
	)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concatenation.Bearer(accessToken))
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
	ctx, cancel := context.WithTimeout(context.Background(), number)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, search,
		bytes.NewBuffer(concatenation.DataBuffer([]string{
			"{\"frontend\": \"false\", \"message\": {\"body\": \"", body,
			"\", \"from_id\": \"", strconv.Itoa(from_id),
			"\", \"kind\": \"Private\", \"to_id\": \"", strconv.Itoa(to_id), "\"}}",
		})),
	)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concatenation.Bearer(accessToken))
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
	ctx, cancel := context.WithTimeout(context.Background(), number)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPut, search,
		bytes.NewBuffer(concatenation.DataCopy(
			46+len(body),
			[]string{"{\"frontend\": \"false\", \"message\": {\"body\": \"", body, "\"}}"},
		)),
	)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concatenation.Bearer(accessToken))
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
	ctx, cancel := context.WithTimeout(context.Background(), number)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, search, nil)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concatenation.Bearer(accessToken))
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
	ctx, cancel := context.WithTimeout(context.Background(), number)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, search, nil)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concatenation.Bearer(accessToken))

	data, status, err := sendRequest(req)
	if err != nil {
		return nil, status, err
	}

	return data, status, nil
}

// Create topic: POST request. To work correctly with the POST method,
// make sure that your application has the 'topics' oauth scope.
func NewCreateTopicPostRequestWithCancel(application, accessToken, search string, body []byte, number time.Duration) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), number)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, search, bytes.NewBuffer(body))
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concatenation.Bearer(accessToken))
	req.Header.Set("Content-Type", "application/json")

	data, status, err := sendRequest(req)
	if err != nil {
		return nil, status, err
	}

	return data, status, nil
}

// Update topic: PATCH request. To work correctly with the PATCH method,
// make sure that your application has the 'topics' oauth scope.
func NewUpdateTopicPatchRequestWithCancel(application, accessToken, search string, body []byte, number time.Duration) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), number)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, search, bytes.NewBuffer(body))
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concatenation.Bearer(accessToken))
	req.Header.Set("Content-Type", "application/json")

	data, status, err := sendRequest(req)
	if err != nil {
		return nil, status, err
	}

	return data, status, nil
}

// Create comment: POST request. To work correctly with the POST method,
// make sure that your application has the 'comments' oauth scope.
func NewCreateCommentPostRequestWithCancel(application, accessToken, search string, body []byte, number time.Duration) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), number)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, search, bytes.NewBuffer(body))
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concatenation.Bearer(accessToken))
	req.Header.Set("Content-Type", "application/json")

	data, status, err := sendRequest(req)
	if err != nil {
		return nil, status, err
	}

	return data, status, nil
}

// Update comment: PATCH request. To work correctly with the PATCH method,
// make sure that your application has the 'comments' oauth scope.
func NewUpdateCommentPatchRequestWithCancel(application, accessToken, search string, body []byte, number time.Duration) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), number)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, search, bytes.NewBuffer(body))
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("User-Agent", application)
	req.Header.Add("Authorization", concatenation.Bearer(accessToken))
	req.Header.Set("Content-Type", "application/json")

	data, status, err := sendRequest(req)
	if err != nil {
		return nil, status, err
	}

	return data, status, nil
}
