package penguin

func Flatten[G []S, S []E, E any](g G) S {
	ret := S{}
	for _, s := range g {
		for _, e := range s {
			ret = append(ret, e)
		}
	}
	return ret
}
