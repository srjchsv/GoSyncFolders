testgo:
	@go test -cover ./...

benchgo:
	@cd ./pkg/utils/ && go test -bench=. -benchmem -run=^$
	@rm -d -r ./test



