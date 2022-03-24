package penguin_test

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/somen440/penguin"
)

func TestExperimental(t *testing.T) {
	t.Parallel()

	t.Run("[]int", func(t *testing.T) {
		t.Parallel()

		ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		want := []int{10, 8, 6, 4, 2}
		got := penguin.Filter(ints, func(e int) bool { return e%2 == 0 })
		got = penguin.Reverse(got)
		if df := cmp.Diff(got, want); df != "" {
			t.Errorf("diff=%s", df)
		}
	})
	t.Run("[]string", func(t *testing.T) {
		t.Parallel()

		strs := []string{"hoge", "foo", "bar"}
		want := []string{"foo", "hoge"}
		got := penguin.Filter(strs, func(e string) bool { return strings.Contains(e, "o") })
		got = penguin.Reverse(got)
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
			{ID: "hoge_id", Name: "hoge_name"},
		}
		got := penguin.Filter(dd, func(e data) bool { return e.ID != "foo_id" })
		got = penguin.Reverse(got)
		if df := cmp.Diff(got, want); df != "" {
			t.Errorf("diff=%s", df)
		}
	})
}
