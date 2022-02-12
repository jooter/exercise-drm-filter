package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrmfilterHandlerErrorGetMethod(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(DrmfilterHandler))
	resp, err := http.Get(srv.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	bs, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, `{
	"error": "un-supported request method"
}
`, string(bs))
}

func TestDrmfilterHandlerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(DrmfilterHandler))
	resp, err := http.Post(srv.URL, "", strings.NewReader(""))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	bs, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, `{
	"error": "Could not decode request"
}
`, string(bs))
}

func TestDrmfilterHandlerSimple(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(DrmfilterHandler))
	resp, err := http.Post(srv.URL, "", strings.NewReader("{}"))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	bs, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, `{
	"response": null
}
`, string(bs))
}

func TestDrmfilterHandlerExample(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(DrmfilterHandler))
	resp, err := http.Post(srv.URL, "", strings.NewReader(test_json_request_1))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	bs, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	// to be improved: struct of elements may be changed, so the test has risk of broken
	assert.Equal(t, test_json_response_1, string(bs))
}
