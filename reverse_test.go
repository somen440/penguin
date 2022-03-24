package penguin_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/somen440/penguin"
)

func TestReverse(t *testing.T) {
	t.Parallel()

	t.Run("[]int", func(t *testing.T) {
		t.Parallel()

		ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		want := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		got := penguin.Reverse(ints)
		if df := cmp.Diff(got, want); df != "" {
			t.Errorf("diff=%s", df)
		}
	})
	t.Run("[]string", func(t *testing.T) {
		t.Parallel()

		strs := []string{"hoge", "foo", "bar"}
		want := []string{"bar", "foo", "hoge"}
		got := penguin.Reverse(strs)
		if df := cmp.Diff(got, want); df != "" {
			t.Errorf("diff=%s", df)
		}
	})
	t.Run("custom struct", func(t *testing.T) {
		t.Parallel()

		type data struct {
			ID   string
			Name string
		}
		dd := []data{
			{ID: "hoge_id", Name: "hoge_name"},
			{ID: "foo_id", Name: "foo_name"},
			{ID: "bar_id", Name: "bar_name"},
		}
		want := []data{
			{ID: "bar_id", Name: "bar_name"},
			{ID: "foo_id", Name: "foo_name"},
			{ID: "hoge_id", Name: "hoge_name"},
		}
		got := penguin.Reverse(dd)
		if df := cmp.Diff(got, want); df != "" {
			t.Errorf("diff=%s", df)
		}
	})
}
