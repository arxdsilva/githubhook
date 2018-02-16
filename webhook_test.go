package webhook_test

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/arxdsilva/webhook"
)

const testSecret = "foobar"

func expectErrorMessage(t *testing.T, msg string, err error) {
	if err == nil || err.Error() != msg {
		t.Error(fmt.Sprintf("Expected '%s', got %s", msg, err))
	}
}

func expectNewError(t *testing.T, msg string, r *http.Request) {
	_, err := webhook.New(r)
	expectErrorMessage(t, msg, err)
}

func expectParseError(t *testing.T, msg string, r *http.Request) {
	_, err := webhook.Parse([]byte(testSecret), r)
	expectErrorMessage(t, msg, err)
}

func signature(body string) string {
	dst := make([]byte, 40)
	computed := hmac.New(sha1.New, []byte(testSecret))
	computed.Write([]byte(body))
	hex.Encode(dst, computed.Sum(nil))
	return "sha1=" + string(dst)
}

func TestNonPost(t *testing.T) {
	r, _ := http.NewRequest("GET", "/path", nil)
	expectNewError(t, "Unknown method!", r)
}

func TestMissingSignature(t *testing.T) {
	r, _ := http.NewRequest("POST", "/path", nil)
	expectNewError(t, "No signature!", r)
}

func TestInvalidSignature(t *testing.T) {
	r, _ := http.NewRequest("POST", "/path", strings.NewReader("..."))
	r.Header.Add("x-hub-signature", "bogus signature")
	expectParseError(t, "Invalid signature", r)
}

func TestValidSignature(t *testing.T) {
	body := "{}"
	r, _ := http.NewRequest("POST", "/path", strings.NewReader(body))
	r.Header.Add("x-hub-signature", signature(body))
	if _, err := webhook.Parse([]byte(testSecret), r); err != nil {
		t.Error(fmt.Sprintf("Unexpected error '%s'", err))
	}
}
