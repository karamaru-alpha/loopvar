## loopvar
loopvar is a linter that detects places where loop variables are copied.

cf. [Fixing For Loops in Go 1.22](https://go.dev/blog/loopvar-preview)

## Rquirements

GO_VERSION >= 1.22 or GOEXPERIMENT=loopvar

## Example
```go
for i, v := range []int{1, 2, 3} {
    i := i // The loop variable "i" should not be copied (GO_VERSION >= 1.22 or GOEXPERIMENT=loopvar)
    v := v // The loop variable "v" should not be copied (GO_VERSION >= 1.22 or GOEXPERIMENT=loopvar)
    _, _ = i, v
}

for i := 1; i <= 3; i++ {
    i := i // The loop variable "i" should not be copied (GO_VERSION >= 1.22 or GOEXPERIMENT=loopvar)
    _ = i
}
```

## Install
```bash
go install github.com/karamaru-alpha/loopvar/cmd/loopvar
go vet -vettool=`which loopvar` ./...
```
