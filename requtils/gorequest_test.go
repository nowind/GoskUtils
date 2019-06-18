package requtils

import (
	"net/http"
	"testing"
)

func TestResponseWrap(t *testing.T) {
	t.Log(ResponseWrap(new(http.Response)))
}