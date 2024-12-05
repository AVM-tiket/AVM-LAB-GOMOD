module github.com/AVM-tiket/AVM-LAB-GOMOD/cmd/main

go 1.20

require (
	github.com/AVM-tiket/AVM-LAB-GOMOD/a v1.2.0
	github.com/AVM-tiket/AVM-LAB-GOMOD/b/v2 v2.1.0
	github.com/AVM-tiket/AVM-LAB-GOMOD/pkg/c v1.1.0
	github.com/AVM-tiket/AVM-LAB-GOMOD/pkg/d v0.0.0-00010101000000-000000000000
)

replace github.com/AVM-tiket/AVM-LAB-GOMOD/pkg/d => ../../pkg/d
