package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/knadh/koanf/providers/confmap"
	bgStateMonitor "github.com/netcracker/qubership-core-lib-go-bg-state-monitor/v2"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/types"
	"github.com/netcracker/qubership-core-lib-go/v3/configloader"
	"github.com/netcracker/qubership-core-lib-go/v3/serviceloader"
	"github.com/netcracker/qubership-core-lib-go/v3/security"
	"github.com/stretchr/testify/require"
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

func TestCreatePlatformService_Consul_Retries(t *testing.T) {
	assertions := require.New(t)
	attempts := 3
	ctx, cancel := context.WithCancel(context.Background())
	ts := createTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/v1/acl/login" {
			if attempts > 0 {
				// imitate like policy is not configured yet
				w.WriteHeader(http.StatusForbidden)
				_, _ = w.Write(nil)
			} else {
				responseMap := map[string]any{
					"SecretID":       "test-secret-id",
					"ExpirationTime": time.Now().Add(time.Hour).Format(time.RFC3339),
				}
				responseBody, _ := json.Marshal(responseMap)
				_, _ = w.Write(responseBody)
			}
			attempts--
		} else if r.Method == "GET" && strings.HasPrefix(r.URL.Path, fmt.Sprintf("/v1/kv/"+bgStateMonitor.BgStateConsulPathNew, testNamespace)) &&
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

	configloader.Init(&configloader.PropertySource{
		Provider: configloader.AsPropertyProvider(confmap.Provider(map[string]any{"identity.provider.url": ts.URL}, ".")),
		Parser:   nil,
	})

	consulRetryDuration = time.Millisecond
	consulRetries = attempts - 1

	builder := NewPlatformClientBuilder().
		WithPlatformType(types.Kubernetes).
		WithClients(prepareFakeClients()).
		WithNamespace(testNamespace).
		WithConsul(true, ts.URL)

	// first attempt must fail because all retries has failed
	client, err := createPlatformService(builder)
	assertions.NotNil(err)
	assertions.Nil(client)

	// second attempt must succeed because request to log in must succeed
	client, err = createPlatformService(builder)
	assertions.NoError(err)
	assertions.NotNil(client)

	cancel()
}

func createTestServer(fn func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(fn))
}
