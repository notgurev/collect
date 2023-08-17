# collect

collect is a little Go library which allows to conveniently convert slices into maps in certain ways.

A slice `[]T` can be converted into:

1. `map[T]bool` with optional key func
2. `map[T]struct{}` with optional key func (more memory-efficient than `map[T]bool` on large maps, but
   less convenient syntactically)
3. `map[T]V` with a key-value func

This package does not contain any common "utility" or "helper" functions for maps and slices themselves, since it has a
different purpose, only the functions described above.

## Documentation

For full documentation see [pkg.go.dev](https://pkg.go.dev/github.com/notgurev/collect).

## Difference between `map[T]bool` and `map[T]struct{}`

`map[T]struct{}` uses slightly less memory compared to `map[T]bool`, but less readable and convenient to use.
No benchmarks yet, but here's
[an article](https://itnext.io/set-in-go-map-bool-and-map-struct-performance-comparison-5315b4b107b) on this topic,
which states the following:

> map[]struct{} is 5% faster in time and 10% less memory
> consumption comparing to map[]bool when it comes to a big Set. 

Using `map[T]struct{}`: 

```go
m := map[T]struct{}{}

if _, ok := m["key"]; ok {
	// ...
}
```

Using `map[T]bool`: 

```go
m := map[T]bool{}

if m["key"] {
	// ...
}
```

## Examples

Converting a string slice into a `map[string]struct{}`:

```go
s := []string{"a", "b", "c"}

m := ToMapOfEmptyStruct(s)

assert.Equal(t, map[string]struct{}{
    "a": {},
    "b": {},
    "c": {},
}, m)
```

Converting a slice of structs into a `map[string]bool` with a custom key function:

```go
type outer struct {
    s string
}

s := []outer{
    {s: "a"},
    {s: "b"},
    {s: "c"},
}

m := ToMapOfBoolFunc(s, func (t outer) string { return t.s })

assert.Equal(t, map[string]bool{
    "a": true,
    "b": true,
    "c": true,
}, m)
```

## Motivation and alternatives

I wasn't able to find an existing library which provides this, yet I needed it for almost every project I worked on.
So I wrote it myself. 

See also:

- Standard library since Go 1.21 has `slices` and `maps` packages with generics, but `slices` does not support
  converting slices to maps in any way, and `maps` only collects keys and values of maps to a slice.
- https://github.com/spf13/cast has a number of ToStringMap functions, but they don't take slices as input, only support
  string keys, and they don't allow to specify a key function.

## TODO

- test coverage
- benchmarks to compare `map[T]bool` with `map[T]struct{}`