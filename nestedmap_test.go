package nestedmap

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type data struct {
	k1 int
	k2 string
	v  int
}

func TestNestedMapInnerValue(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		values []data
		k1     int
		k2     string
		ok     bool
		want   int
	}{
		"can get inner value when existing data": {
			values: []data{
				{
					k1: 1,
					k2: "a",
					v:  10,
				},
				{
					k1: 1,
					k2: "b",
					v:  11,
				},
				{
					k1: 2,
					k2: "a",
					v:  20,
				},
			},
			k1:   1,
			k2:   "a",
			ok:   true,
			want: 10,
		},
		"cannot get inner value when not existing data": {
			values: []data{},
			k1:     1,
			k2:     "a",
			ok:     false,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			nm := NewNestedMap[int, string, int]()
			for _, data := range tt.values {
				nm.Set(data.k1, data.k2, data.v)
			}

			got, ok := nm.GetInner(tt.k1, tt.k2)
			if ok != tt.ok {
				t.Fatalf("got: %v, want: %v", ok, tt.ok)
				return
			}
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestNestedMapOuterValue(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		values []data
		k1     int
		ok     bool
		want   map[string]int
	}{
		"can get outer value when existing data": {
			values: []data{
				{
					k1: 1,
					k2: "a",
					v:  10,
				},
				{
					k1: 1,
					k2: "b",
					v:  11,
				},
				{
					k1: 2,
					k2: "a",
					v:  20,
				},
			},
			k1:   1,
			ok:   true,
			want: map[string]int{"a": 10, "b": 11},
		},
		"cannot get outer value when not existing data": {
			values: []data{},
			k1:     1,
			ok:     false,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			nm := NewNestedMap[int, string, int]()
			for _, data := range tt.values {
				nm.Set(data.k1, data.k2, data.v)
			}

			got, ok := nm.GetOuter(tt.k1)
			if ok != tt.ok {
				t.Fatalf("got: %v, want: %v", ok, tt.ok)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("diff: %s", diff)
			}
		})
	}
}
