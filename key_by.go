package penguin

func KeyBy[S []E, E any, K comparable](s S, key func(e E) K) map[K]E {
	ret := map[K]E{}
	for _, e := range s {
		k := key(e)
		ret[k] = e
	}
	return ret
}
