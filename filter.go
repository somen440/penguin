package penguin

func Filter[S []E, E any](s S, cond func(e E) bool) S {
	ret := S{}
	for _, v := range s {
		if !cond(v) {
			continue
		}
		ret = append(ret, v)
	}
	return ret
}
