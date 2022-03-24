package penguin

func Uniq[S []E, E any](s S) S {
	ret := S{}
	eM := map[any]struct{}{}
	for _, e := range s {
		if _, ok := eM[e]; ok {
			continue
		}
		eM[e] = struct{}{}
		ret = append(ret, e)
	}
	return ret
}
