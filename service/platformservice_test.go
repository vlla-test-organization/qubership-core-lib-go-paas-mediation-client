package service

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	bgStateMonitor "github.com/vlla-test-organization/qubership-core-lib-go-bg-state-monitor/v2"
	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/types"
	"github.com/vlla-test-organization/qubership-core-lib-go/v3/security"
	"github.com/vlla-test-organization/qubership-core-lib-go/v3/serviceloader"
)

func init() {
	serviceloader.Register(1, &security.DummyToken{})
}

func TestCreatePlatformService_Consul(t *testing.T) {
	assertions := require.New(t)

	ctx, cancel := context.WithCancel(context.Background())
	ts := createTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" && strings.HasPrefix(r.URL.Path, fmt.Sprintf("/v1/kv/"+bgStateMonitor.BgStateConsulPathNew, testNamespace)) &&
			r.URL.Query().Get("index") == "" {
			w.Header().Add("X-Consul-Index", "1")
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write(nil)
		} else if r.Method == "GET" && strings.HasPrefix(r.URL.Path, fmt.Sprintf("/v1/kv/"+bgStateMonitor.BgStateConsulPath, testNamespace)) &&
			r.URL.Query().Get("index") == "" {
			w.Header().Add("X-Consul-Index", "1")
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write(nil)
		} else if r.Method == "GET" && strings.HasPrefix(r.URL.Path, fmt.Sprintf("/v1/kv/"+bgStateMonitor.BgStateConsulPath, testNamespace)) &&
			r.URL.Query().Get("index") == "1" {
			// wait to imitate long polling request
			select {
			case <-ctx.Done():
				w.Header().Add("X-Consul-Index", "2")
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write(nil)
			}
		} else {
			panic("unexpected consul request, method: " + r.Method + ", url: " + r.URL.Path)
		}
	})
	defer ts.Close()

	builder := NewPlatformClientBuilder().
		WithPlatformType(types.Kubernetes).
		WithClients(prepareFakeClients()).
		WithNamespace(testNamespace).
		WithConsul(true, ts.URL, "test-token")

	client, err := createPlatformService(builder)
	assertions.Nil(err)
	assertions.NotNil(client)
	cancel()
}

func TestCreatePlatformService_Consul_Enabled(t *testing.T) {
	assertions := require.New(t)
	attempts := 3
	ctx, cancel := context.WithCancel(context.Background())
	ts := createTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" && strings.HasPrefix(r.URL.Path, fmt.Sprintf("/v1/kv/"+bgStateMonitor.BgStateConsulPathNew, testNamespace)) &&
			r.URL.Query().Get("index") == "" {
			w.Header().Add("X-Consul-Index", "1")
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write(nil)
		} else if r.Method == "GET" && strings.HasPrefix(r.URL.Path, fmt.Sprintf("/v1/kv/"+bgStateMonitor.BgStateConsulPath, testNamespace)) &&
			r.URL.Query().Get("index") == "" {
			w.Header().Add("X-Consul-Index", "1")
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write(nil)
		} else if r.Method == "GET" && strings.HasPrefix(r.URL.Path, fmt.Sprintf("/v1/kv/"+bgStateMonitor.BgStateConsulPath, testNamespace)) &&
			r.URL.Query().Get("index") == "1" {
			// wait to imitate long polling request
			select {
			case <-ctx.Done():
				w.Header().Add("X-Consul-Index", "1")
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write(nil)
			}
		} else {
			panic("unexpected consul request, method: " + r.Method + ", url: " + r.URL.Path)
		}
	})
	defer ts.Close()

	consulRetryDuration = time.Millisecond
	consulRetries = attempts - 1

	builder := NewPlatformClientBuilder().
		WithPlatformType(types.Kubernetes).
		WithClients(prepareFakeClients()).
		WithNamespace(testNamespace).
		WithConsul(true, ts.URL)

	client, err := createPlatformService(builder)
	assertions.NoError(err)
	assertions.NotNil(client)

	cancel()
}

func createTestServer(fn func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(fn))
}
