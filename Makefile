test:
	go clean -testcache
	go test -cover ./...

bench:
	go clean -testcache
	go test -bench=. -benchmem ./...