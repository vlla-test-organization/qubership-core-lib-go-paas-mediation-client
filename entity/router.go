package entity

import (
	"regexp"
	"strings"
)

// todo should it be removed?
type Registration struct {
	Name    string `json:"name"`
	Service string `json:"service"`
	Route   string `json:"rule"`
}

type Registrations []Registration

func (r *Registrations) ToRoutes(hostFilter string) *[]Route {
	var routes []Route

	for _, reg := range *r {
		route := *reg.ToRoute()
		if strings.Contains(route.Spec.Host, hostFilter) {
			routes = append(routes, route)
		}
	}

	return &routes
}

func (r *Registration) ToRoute() *Route {
	metadata := Metadata{Kind: "Route", Name: r.Name, Namespace: "local", Annotations: *new(map[string]string), Labels: *new(map[string]string)}

	target := Target{Name: r.Service}
	routeSpec := RouteSpec{
		Service: target,
		Host:    r.getHost(),
	}

	return &Route{Spec: routeSpec, Metadata: metadata}
}

func (r *Registration) getHost() string {
	re := regexp.MustCompile(`[a-zA-Z0-9.-]+`)
	return re.FindString(strings.TrimPrefix(r.Route, "Host"))
}
