package penguin

func Reverse[S []E, E any](s S) S {
	idx := 0
	ret := make(S, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		ret[idx] = s[i]
		idx++
	}
	return ret
}
