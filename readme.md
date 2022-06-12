# gosyncfolders
### A Go program that scans and syncs source and destination folders

----------------
### To run program first you need to clone or download the repo
### and then install all dependencies with:
### `go mod tidy`
### Finally install the program using:
### `go install `
#### or you can execute `export GO111MODULE=off`
#### then `go get github.com/srjchsv/gosyncfolders`
#### and finally  `go install github.com/srjchsv/gosyncfolders`


### And then let GO sync folders for you =) run :
### `gosyncfolders go src/ dst/`
`src/` path to a source directory

`dst/` path to a destination directory

### or use  `gosyncfolders go -h` for help


#### _when running program log.txt with logging will be created in current directory_

------------------------
## To run tests or benchmarks execute commands in the root gosyncfolders directory :
### - Tests `make testgo`
### - Benchmarks `make benchgo`
### - Benchstat `make benchstatgo`
