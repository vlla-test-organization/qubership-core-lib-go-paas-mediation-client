package filter

type Meta struct {
	Labels        map[string]string
	Annotations   map[string]string
	WatchBookmark bool
}

func (f Meta) GetAnnotations() map[string]string {
	return f.Annotations
}

func (f Meta) GetLabels() map[string]string {
	return f.Labels
}
