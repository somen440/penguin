package penguin

func Reduce[S []E, E any](s S, reducer func(prev E, current E) E) E {
	var ret E
	for _, e := range s {
		ret = reducer(ret, e)
	}
	return ret
}
