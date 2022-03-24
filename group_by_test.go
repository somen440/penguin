package penguin_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/somen440/penguin"
)

func TestGroupBy(t *testing.T) {
	t.Parallel()

	type data struct {
		ID    string
		Name  string
		Group string
	}
	dd := []data{
		{ID: "hoge_id", Name: "hoge_name", Group: "group_1"},
		{ID: "foo_id", Name: "foo_id", Group: "group_1"},
		{ID: "bar_id", Name: "bar_id", Group: "group_2"},
	}
	want := map[string][]data{
		"group_1": {
			{ID: "hoge_id", Name: "hoge_name", Group: "group_1"},
			{ID: "foo_id", Name: "foo_id", Group: "group_1"},
		},
		"group_2": {
			{ID: "bar_id", Name: "bar_id", Group: "group_2"},
		},
	}
	got := penguin.GroupBy(dd, func(e data) string { return e.Group })
	if df := cmp.Diff(got, want); df != "" {
		t.Errorf("diff=%s", df)
	}
}
