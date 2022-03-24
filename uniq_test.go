package penguin_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/somen440/penguin"
)

func TestUniq(t *testing.T) {
	t.Parallel()

	t.Run("int slice", func(t *testing.T) {
		t.Parallel()

		ints := []int{1, 1, 2, 3, 3}
		want := []int{1, 2, 3}
		got := penguin.Uniq(ints)
		if df := cmp.Diff(got, want); df != "" {
			t.Errorf("diff=%s", df)
		}
	})
	t.Run("string slice", func(t *testing.T) {
		t.Parallel()

		strs := []string{"hoge", "foo", "foo", "bar"}
		want := []string{"hoge", "foo", "bar"}
		got := penguin.Uniq(strs)
		if df := cmp.Diff(got, want); df != "" {
			t.Errorf("diff=%s", df)
		}
	})
	t.Run("custom struct slice", func(t *testing.T) {
		t.Parallel()

		type data struct {
			ID   string
			Name string
		}
		dd := []data{
			{ID: "hoge_id", Name: "foo_id"},
			{ID: "hoge_id", Name: "foo_id"},
			{ID: "hoge_id", Name: "foo_id"},
			{ID: "hoge_id", Name: "foo_id"},
			{ID: "hoge_id", Name: "foo_id"},
		}
		want := []data{
			{ID: "hoge_id", Name: "foo_id"},
		}
		got := penguin.Uniq(dd)
		if df := cmp.Diff(got, want); df != "" {
			t.Errorf("diff=%s", df)
		}
	})
}
