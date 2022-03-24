package penguin

func GroupBy[S []E, E any, K comparable](s S, key func(e E) K) map[K]S {
	ret := map[K]S{}
	for _, e := range s {
		k := key(e)
		ret[k] = append(ret[k], e)
	}
	return ret
}
