package openshiftV3

import (
	"context"

	openshift_v1 "github.com/openshift/api/route/v1"
	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/filter"
	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/service/internal/kubernetes"
)

func (os *OpenshiftV3Client) GetRoute(ctx context.Context, resourceName string, namespace string) (*entity.Route, error) {
	if result, err := os.Kubernetes.GetRoute(ctx, resourceName, namespace); err != nil {
		return nil, err
	} else if result != nil {
		return result, nil
	}
	return kubernetes.GetWrapper(ctx, resourceName, namespace, os.RouteV1Client.Routes(namespace).Get,
		os.Cache.Ingresses, entity.RouteFromOsRoute)
}

func (os *OpenshiftV3Client) GetRouteList(ctx context.Context, namespace string, filter filter.Meta) ([]entity.Route, error) {
	result, err := os.Kubernetes.GetRouteList(ctx, namespace, filter)
	if err != nil {
		return nil, err
	}
	result2, err := kubernetes.ListWrapper(ctx, filter, os.RouteV1Client.Routes(namespace).List, os.Cache.Ingresses,
		func(listObj *openshift_v1.RouteList) (result []entity.Route) {
			for _, item := range listObj.Items {
				route := entity.RouteFromOsRoute(&item)
				if route != nil {
					result = append(result, *route)
				}
			}
			return
		})
	if err != nil {
		return nil, err
	}
	return append(result, result2...), nil
}
