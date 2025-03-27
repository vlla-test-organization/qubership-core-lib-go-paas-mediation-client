package mock

import (
	"context"
	"errors"
	v1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	"k8s.io/api/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	"k8s.io/client-go/rest"
)

type FakePods struct {
	Fake *FakeCoreV1
	ns   string
}

func (f FakePods) UpdateResize(ctx context.Context, podName string, pod *v1.Pod, opts metav1.UpdateOptions) (*v1.Pod, error) {
	panic("not implemented")
}

func (f FakePods) Create(ctx context.Context, pod *v1.Pod, opts metav1.CreateOptions) (*v1.Pod, error) {
	panic("not implemented")
}

func (f FakePods) Update(ctx context.Context, pod *v1.Pod, opts metav1.UpdateOptions) (*v1.Pod, error) {
	panic("not implemented")
}

func (f FakePods) UpdateStatus(ctx context.Context, pod *v1.Pod, opts metav1.UpdateOptions) (*v1.Pod, error) {
	panic("not implemented")
}

func (f FakePods) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	panic("not implemented")
}

func (f FakePods) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	panic("not implemented")
}

func (f FakePods) Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Pod, error) {
	return nil, errors.New("error test")
}

func (f FakePods) List(ctx context.Context, opts metav1.ListOptions) (*v1.PodList, error) {
	return nil, errors.New("WebSocket error test")
}

func (f FakePods) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	panic("not implemented")
}

func (f FakePods) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Pod, err error) {
	panic("not implemented")
}

func (f FakePods) Apply(ctx context.Context, pod *corev1.PodApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Pod, err error) {
	panic("not implemented")
}

func (f FakePods) ApplyStatus(ctx context.Context, pod *corev1.PodApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Pod, err error) {
	panic("not implemented")
}

func (f FakePods) UpdateEphemeralContainers(ctx context.Context, podName string, pod *v1.Pod, opts metav1.UpdateOptions) (*v1.Pod, error) {
	panic("not implemented")
}

func (f FakePods) Bind(ctx context.Context, binding *v1.Binding, opts metav1.CreateOptions) error {
	panic("not implemented")
}

func (f FakePods) Evict(ctx context.Context, eviction *v1beta1.Eviction) error {
	panic("not implemented")
}

func (f FakePods) EvictV1(ctx context.Context, eviction *policyv1.Eviction) error {
	panic("not implemented")
}

func (f FakePods) EvictV1beta1(ctx context.Context, eviction *v1beta1.Eviction) error {
	panic("not implemented")
}

func (f FakePods) GetLogs(name string, opts *v1.PodLogOptions) *rest.Request {
	panic("not implemented")
}

func (f FakePods) ProxyGet(scheme, name, port, path string, params map[string]string) rest.ResponseWrapper {
	panic("not implemented")
}
