package controller

import (
		"testing"	
		"github.com/stretchr/testify/assert"
		"net/http"
		"encoding/json"
		"fmt"
	    "bytes"
	    "io"
		)

type nopCloser struct {
    io.Reader
}

func TestDecodeResponse(t *testing.T) {
    data := "{\"description\" : \"This a sample json response\"}"
    response := &http.Response{
        StatusCode:      404,
        Body: nopCloser{bytes.NewBufferString("{\"Error Encountered\" : \"We have encountered an Error\"}")},
    }
    err := decodeResponse(response, data)
    assert.NotNil(t, err, "We are expecting error and got one")
    assert.Equal(t, "Got 404 returned status", err.Error())
}

func (nopCloser) Close() error { return nil }

func TestStub(t *testing.T) {
    assert.True(t, true, "This is good. Canary test passing")
}

func decodeResponse(resp *http.Response, data interface{}) error {
    if resp.StatusCode != 200 {
        return fmt.Errorf("Got %d returned status", resp.StatusCode)
    }
    return json.NewDecoder(resp.Body).Decode(data)
}

func TestHttpRequest(t *testing.T) {
	_, err := http.NewRequest("GET","/",nil)
	if err != nil {
		t.Fatal("Creating GET to / request failed!")
	}

	t.Log("Test passed")
}