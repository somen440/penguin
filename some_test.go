package penguin_test

import (
	"strings"
	"testing"

	"github.com/somen440/penguin"
)

func TestSome(t *testing.T) {
	t.Parallel()

	t.Run("int slice", func(t *testing.T) {
		t.Parallel()

		type args struct {
			ints []int
		}
		type want struct {
			ret bool
		}
		tests := []struct {
			title string
			args  args
			want  want
		}{
			{
				title: "is some",
				args: args{
					ints: []int{2, 3, 4, 6},
				},
				want: want{
					ret: true,
				},
			},
			{
				title: "is not some",
				args: args{
					ints: []int{1, 3, 5},
				},
				want: want{
					ret: false,
				},
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.title, func(t *testing.T) {
				t.Parallel()

				got := penguin.Some(tt.args.ints, func(e int) bool { return e%2 == 0 })
				if got != tt.want.ret {
					t.Errorf("want=%+v, got=%+v", tt.want.ret, got)
				}
			})
		}
	})
	t.Run("string slice", func(t *testing.T) {
		t.Parallel()

		type args struct {
			strs []string
		}
		type want struct {
			ret bool
		}
		tests := []struct {
			title string
			args  args
			want  want
		}{
			{
				title: "is some",
				args: args{
					strs: []string{"hoge", "foo", "bar"},
				},
				want: want{
					ret: true,
				},
			},
			{
				title: "is not some",
				args: args{
					strs: []string{"bar", "bar", "bar"},
				},
				want: want{
					ret: false,
				},
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.title, func(t *testing.T) {
				t.Parallel()

				got := penguin.Some(tt.args.strs, func(e string) bool { return strings.Contains(e, "o") })
				if got != tt.want.ret {
					t.Errorf("want=%+v, got=%+v", tt.want.ret, got)
				}
			})
		}
	})
	t.Run("custom struct slice", func(t *testing.T) {
		t.Parallel()

		type data struct {
			ID   string
			Name string
		}

		type args struct {
			dd []data
		}
		type want struct {
			ret bool
		}
		tests := []struct {
			title string
			args  args
			want  want
		}{
			{
				title: "is some",
				args: args{
					dd: []data{
						{ID: "hoge_id", Name: "hoge_name"},
						{ID: "foo_id", Name: "foo_name"},
						{ID: "bar_id", Name: "bar_name"},
					},
				},
				want: want{
					ret: true,
				},
			},
			{
				title: "is not some",
				args: args{
					dd: []data{
						{ID: "bar_id", Name: "bar_name"},
						{ID: "bar_id", Name: "bar_name"},
						{ID: "bar_id", Name: "bar_name"},
					},
				},
				want: want{
					ret: false,
				},
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.title, func(t *testing.T) {
				t.Parallel()

				got := penguin.Some(tt.args.dd, func(e data) bool { return e.ID != "bar_id" })
				if got != tt.want.ret {
					t.Errorf("want=%+v, got=%+v", tt.want.ret, got)
				}
			})
		}
	})
}
