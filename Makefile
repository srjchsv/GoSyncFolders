testgo:
	@go test -cover ./...
	@rm -d -r ./test

benchgo:
	@cd ./pkg/utils/ && go test -bench=. -benchmem -benchtime=5s > result.txt
	@rm -d -r ./test

benchstatgo:
	@cd ./pkg/utils &&  benchstat -sort -name result.txt


