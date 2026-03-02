# Go Lang Practice

A hands-on collection of Go examples covering fundamentals through advanced concurrency patterns — built as a progressive learning path.

---

## Topics

| # | Topic | Description |
|---|-------|-------------|
| 01 | [Hello World](1_hello_world/) | Minimal Go program with `fmt.Println` |
| 02 | [Simple Values](2_simple_values/) | Basic types — integers, floats, booleans, strings |
| 03 | [Variables](3_variables/) | `var`, `:=`, type inference, zero values, scope |
| 04 | [Constants](4_constants/) | `const`, grouped constants, immutability rules |
| 05 | [For Loops](5_for/) | Classic for, while-style, infinite loop, `break`/`continue`, `range` |
| 06 | [If / Else](6_if_else/) | Conditionals, compound conditions, init statements |
| 07 | [Switch](7_switch/) | Expression & expressionless switch, multi-value cases, type switch |
| 08 | [Arrays](8_arrays/) | Fixed-size arrays, indexing, iteration, 2D arrays |
| 09 | [Slices](9_slices/) | Dynamic slices, `append`, `copy`, `make`, `slices` package |
| 10 | [Maps](10_maps/) | Key-value pairs, comma-ok idiom, `delete`, `maps` package |
| 11 | [Range](11_range/) | `range` over arrays, slices, maps, and strings (runes) |
| 12 | [Functions](12_functions/) | Multiple returns, named returns, higher-order & anonymous functions |
| 13 | [Variadic Functions](13_variadic_functions/) | `...` syntax, spreading slices, `...interface{}` |
| 14 | [Closures](14_closures/) | Capturing variables, counter closures, multiplier factories |
| 15 | [Pointers](15_pointers/) | `&`, `*`, pointer-to-pointer, pass-by-reference |
| 16 | [Structs](16_structs/) | Fields, methods, pointer receivers, embedding, factory functions |
| 17 | [Interfaces](17_interfaces/) | Implicit implementation, polymorphism, `Stringer`, type assertions |
| 18 | [Enums](18_enums/) | `iota` for integer enums, typed string constants |
| 19 | [Generics](19_generics/) | Type parameters, constraints, generic structs (e.g. `stack[T]`) |
| 20 | [Goroutines](20_goroutines/) | `go` keyword, `sync.WaitGroup`, closure pitfalls |
| 21 | [Channels](21_channels/) | Buffered & unbuffered, `range`, `close`, done-channels, pipelines |
| 22 | [Mutex](22_mutex/) | `sync.Mutex`, race condition prevention, `defer Unlock` |
| 23 | [Files](23_files/) | Read, write, copy, `os.Stat`, `bufio`, directory listing |
| 24 | [Packages](24_packages/) | Module organisation, sub-packages, external dependencies |
| 25 | [Patterns](25_patterns/) | Fan-out/fan-in, worker pools, error handling in concurrency |

---

## Running Examples

```bash
# Run any example
cd <folder>
go run <file>.go

# e.g.
cd 20_goroutines
go run goroutine.go
```

---

## Open-Source Projects to Refer

| Project | Repository |
|---------|------------|
| traQ | https://github.com/traPtitech/traQ |
| Harness | https://github.com/harness/harness |
| Incubator Answer | https://github.com/shuashuai/incubator-answer |
| Focalboard | https://github.com/mattermost-community/focalboard |
| Fider | https://github.com/getfider/fider |