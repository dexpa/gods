package slice_test

import (
	"testing"
	"github.com/cheebo/gods/slice"
	"github.com/stretchr/testify/assert"
)

func TestStringSlice_Contains(t *testing.T) {
	asserts := assert.New(t)

	ss := slice.StringSlice{"one", "two", "three"}

	asserts.Equal(0, ss.Contains("one"))
	asserts.Equal(1, ss.Contains("two"))
	asserts.Equal(2, ss.Contains("three"))
	asserts.Equal(-1, ss.Contains("four"))
}

func TestIntSlice_Contains(t *testing.T) {
	asserts := assert.New(t)

	ss := slice.IntSlice{1,2,3}

	asserts.Equal(0, ss.Contains(1))
	asserts.Equal(1, ss.Contains(2))
	asserts.Equal(2, ss.Contains(3))
	asserts.Equal(-1, ss.Contains(4))
}

func TestFloat64Slice_Contains(t *testing.T) {
	asserts := assert.New(t)

	ss := slice.IntSlice{10.0,20.0,30.0}

	asserts.Equal(0, ss.Contains(10.0))
	asserts.Equal(1, ss.Contains(20.0))
	asserts.Equal(2, ss.Contains(30.0))
	asserts.Equal(-1, ss.Contains(40.0))
}
