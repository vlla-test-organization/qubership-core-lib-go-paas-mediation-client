package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getTestRegistration() *Registration {
	return &Registration{Name: "regName",
		Service: "regService",
		Route:   "Host registration1"}
}

func getTestRouter(reg *Registration) *Route {
	metadata := Metadata{Kind: "Route", Name: reg.Name, Namespace: "local"}
	routeSpec := RouteSpec{
		Service: Target{Name: reg.Service},
		Host:    reg.getHost(),
	}
	return &Route{Metadata: metadata, Spec: routeSpec}
}

func TestGetHost(t *testing.T) {
	registration := &Registration{Route: "Host !!!1599"}
	host := registration.getHost()
	assert.Equal(t, "1599", host)
	registration.Route = "Just a small test"
	host = registration.getHost()
	assert.NotEqual(t, registration.Route, host)
}

func TestToRoute(t *testing.T) {
	registration := getTestRegistration()
	testRouter := getTestRouter(registration)
	router := registration.ToRoute()
	assert.Equal(t, testRouter, router)
}

func TestRegistrations_ToRoutes(t *testing.T) {
	registrations := Registrations{{Route: "Host !!!filter1599"}, {Route: "filter98"}, {Route: "Host test"}}
	testRoutes := &[]Route{}
	*testRoutes = append(*testRoutes, *registrations[0].ToRoute(), *registrations[1].ToRoute())
	routes := registrations.ToRoutes("filter")
	assert.Equal(t, testRoutes, routes)
}
