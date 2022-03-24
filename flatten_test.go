package penguin_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/somen440/penguin"
)

func TestFlatten(t *testing.T) {
	t.Parallel()

	t.Run("int slice group", func(t *testing.T) {
		t.Parallel()

		intsG := [][]int{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
		}
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := penguin.Flatten(intsG)
		if df := cmp.Diff(got, want); df != "" {
			t.Errorf("diff=%s", df)
		}
	})
	t.Run("string slice group", func(t *testing.T) {
		t.Parallel()

		strsG := [][]string{
			{"hoge"},
			{"foo"},
			{"bar"},
		}
		want := []string{"hoge", "foo", "bar"}
		got := penguin.Flatten(strsG)
		if df := cmp.Diff(got, want); df != "" {
			t.Errorf("diff=%s", df)
		}
	})
	t.Run("custom struct slice group", func(t *testing.T) {
		t.Parallel()

		type data struct {
			ID   string
			Name string
		}
		ddG := [][]data{
			{
				{ID: "hoge_id", Name: "hoge_name"},
				{ID: "foo_id", Name: "foo_name"},
			},
			{
				{ID: "bar_id", Name: "bar_name"},
			},
		}
		want := []data{
			{ID: "hoge_id", Name: "hoge_name"},
			{ID: "foo_id", Name: "foo_name"},
			{ID: "bar_id", Name: "bar_name"},
		}
		got := penguin.Flatten(ddG)
		if df := cmp.Diff(got, want); df != "" {
			t.Errorf("diff=%s", df)
		}
	})
}
