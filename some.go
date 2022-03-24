package penguin

func Some[S []E, E any](s S, cond func(e E) bool) bool {
	for _, e := range s {
		if cond(e) {
			return true
		}
	}
	return false
}
