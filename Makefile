


.PHONY: run
run:
	go run -tags wireinject cmd/main.go cmd/wire_gen.go
	
.PHONY: build
build:
	CGO_ENABLED=1 go build -tags wireinject -o service ./cmd/main.go ./cmd/wire_gen.go

.PHONY: api
api:
	bug generate

