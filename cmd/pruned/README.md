# cmd/pruned

Notice that `indirect v1.0.0` requires `direct v1.0.0` for `indirect/subpkg.Echo()`, but it is not
imported due to [graph pruning](https://go.dev/ref/mod#graph-pruning).

```
$ go mod graph
github.com/AVM-tiket/AVM-LAB-GOMOD/cmd/pruned github.com/AVM-tiket/AVM-LAB-GOMOD/indirect@v1.0.0
github.com/AVM-tiket/AVM-LAB-GOMOD/cmd/pruned go@1.20
github.com/AVM-tiket/AVM-LAB-GOMOD/indirect@v1.0.0 github.com/AVM-tiket/AVM-LAB-GOMOD/direct@v1.0.0
```
