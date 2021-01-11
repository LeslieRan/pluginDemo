package pluginDemo_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LeslieRan/pluginDemo"
)

func TestDemo(t *testing.T) {
	cfg := pluginDemo.CreateConfig()
	cfg.Info = "Hello"

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := pluginDemo.New(ctx, next, cfg, "demo-plugin")
	if err != nil {
		t.Fatalf("new plugin:%s\n", err.Error())
	}

	recorder := httptest.NewRecorder()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatalf("new http request:%s\n", err.Error())
	}
	// client := http.DefaultClient
	// resp, err := client.Do(req)
	// if err != nil {
	// 	t.Fatalf("send http request:%s\n", err.Error())
	// }
	// t.Logf("response: %d and %s\n", resp.StatusCode, resp.Status)
	handler.ServeHTTP(recorder, req)
}
