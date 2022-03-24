package penguin_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/somen440/penguin"
)

func TestReducer(t *testing.T) {
	t.Parallel()

	t.Run("int slice", func(t *testing.T) {
		t.Parallel()

		ints := []int{1, 2, 3, 4, 5}
		want := 15
		got := penguin.Reduce(ints, func(prev, current int) int { return prev + current })
		if df := cmp.Diff(got, want); df != "" {
			t.Errorf("diff=%s", df)
		}
	})
	t.Run("string slice", func(t *testing.T) {
		t.Parallel()

		strs := []string{"h", "e", "l", "l", "o"}
		want := "hello"
		got := penguin.Reduce(strs, func(prev, current string) string { return prev + current })
		if df := cmp.Diff(got, want); df != "" {
			t.Errorf("diff=%s", df)
		}
	})
}
