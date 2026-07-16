package handler
package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndexHandlerRendersSuccessfully(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	IndexHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	body := rr.Body.String()
	if !strings.Contains(body, "KELIN SALOM") {
		t.Fatalf("expected rendered page content, got %q", body)
	}
}

func TestPoemHandlerRendersSuccessfully(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/poem/1", nil)
	rr := httptest.NewRecorder()

	PoemHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	body := rr.Body.String()
	if !strings.Contains(body, "<title>") {
		t.Fatalf("expected rendered poem page, got %q", body)
	}
}
