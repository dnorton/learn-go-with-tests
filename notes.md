
lesson: <https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world>

## Running Go without module

```bash
❯ go test
--- FAIL: TestHello (0.00s)
    hello_test.go:10: got "Hello, World!" want "Hello, world."
FAIL
exit status 1
FAIL	github.com/dnorton/learn-go/learn-with-tests/helloworld	0.203s
```

## Benchmarks

```go
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("world")
	}
}
```

to run `go test -bench=.`

You'll get output like this:

```bash
goos: darwin
goarch: amd64
pkg: github.com/dnorton/learn-go/learn-with-tests/helloworld
BenchmarkHello-8    1000000000	         80.282 ns/op
PASS
```

The `80.282 ns/op` means that it took an average of 80.282 nanoseconds to run the function. The `-8` means that it used 8 cores to run the benchmark.
