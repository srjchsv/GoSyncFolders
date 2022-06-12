syncgo:
	go run cmd/app/main.go ../SRC ../DEST

gosyncgo:
	go run cmd/app/main.go ../SRC ../DEST  &


testgo:
	@go test -cover ./...
	@rm -d -r ./test


benchgo:
	@cd ./pkg/utils/ && go test -bench=. -benchmem -benchtime=5s > result.txt
	@rm -d -r ./test

benchstatgo:
	@cd ./pkg/utils &&  benchstat -sort -name result.txt


