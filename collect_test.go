package collect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToMapOfBool(t *testing.T) {
	t.Run("string slice", func(t *testing.T) {
		s := []string{"a", "b", "c"}

		m := ToMapOfBool(s)

		assert.Equal(t, map[string]bool{
			"a": true,
			"b": true,
			"c": true,
		}, m)
	})

	t.Run("empty string slice", func(t *testing.T) {
		var s []string

		m := ToMapOfBool(s)

		assert.Equal(t, map[string]bool{}, m)
	})

	t.Run("pointer slice", func(t *testing.T) {
		a, b, c := 1, 2, 3

		s := []*int{&a, &b, &c}

		m := ToMapOfBool(s)

		assert.Equal(t, map[*int]bool{
			&a: true,
			&b: true,
			&c: true,
		}, m)
	})
}

func TestToMapOfEmptyStruct(t *testing.T) {
	t.Run("string slice", func(t *testing.T) {
		s := []string{"a", "b", "c"}

		m := ToMapOfEmptyStruct(s)

		assert.Equal(t, map[string]struct{}{
			"a": {},
			"b": {},
			"c": {},
		}, m)
	})

	t.Run("empty string slice", func(t *testing.T) {
		var s []string

		m := ToMapOfEmptyStruct(s)

		assert.Equal(t, map[string]struct{}{}, m)
	})
}

func TestToMapOfBoolFunc(t *testing.T) {
	type outer struct {
		s string
	}

	t.Run("string slice", func(t *testing.T) {
		s := []outer{
			{s: "a"},
			{s: "b"},
			{s: "c"},
		}

		m := ToMapOfBoolFunc(s, func(t outer) string { return t.s })

		assert.Equal(t, map[string]bool{
			"a": true,
			"b": true,
			"c": true,
		}, m)
	})

	t.Run("empty string slice", func(t *testing.T) {
		var s []outer

		m := ToMapOfBoolFunc(s, func(t outer) string { return t.s })

		assert.Equal(t, map[string]bool{}, m)
	})
}

func TestToMapOfEmptyStructFunc(t *testing.T) {
	type outer struct {
		s string
	}

	t.Run("string slice", func(t *testing.T) {
		s := []outer{
			{s: "a"},
			{s: "b"},
			{s: "c"},
		}

		m := ToMapOfEmptyStructFunc(s, func(t outer) string { return t.s })

		assert.Equal(t, map[string]struct{}{
			"a": {},
			"b": {},
			"c": {},
		}, m)
	})

	t.Run("empty string slice", func(t *testing.T) {
		var s []outer

		m := ToMapOfEmptyStructFunc(s, func(t outer) string { return t.s })

		assert.Equal(t, map[string]struct{}{}, m)
	})
}

func TestToMap(t *testing.T) {
	type entity struct {
		id   int
		text string
	}

	s := []entity{
		{id: 1, text: "a"},
		{id: 2, text: "b"},
		{id: 3, text: "c"},
	}

	// Using a single kvFunc is chosen over two separate functions because anonymous
	// function declaration in Go is very verbose. The goal was to make it fit in one
	// line while keeping it readable.
	m := ToMap(s, func(e entity) (int, string) { return e.id, e.text })

	assert.Equal(t, map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}, m)
}
