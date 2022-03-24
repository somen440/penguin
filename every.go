package penguin

func Every[S []E, E any](s S, cond func(e E) bool) bool {
	for _, e := range s {
		if !cond(e) {
			return false
		}
	}
	return true
}
