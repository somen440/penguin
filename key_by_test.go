package penguin_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/somen440/penguin"
)

func TestKeyBy(t *testing.T) {
	t.Parallel()

	type data struct {
		ID   string
		Name string
	}
	dd := []data{
		{ID: "hoge_id", Name: "hoge_name"},
		{ID: "foo_id", Name: "foo_id"},
		{ID: "bar_id", Name: "bar_id"},
	}
	want := map[string]data{
		"hoge_id": {ID: "hoge_id", Name: "hoge_name"},
		"foo_id":  {ID: "foo_id", Name: "foo_id"},
		"bar_id":  {ID: "bar_id", Name: "bar_id"},
	}
	got := penguin.KeyBy(dd, func(e data) string { return e.ID })
	if df := cmp.Diff(got, want); df != "" {
		t.Errorf("diff=%s", df)
	}
}
