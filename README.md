# svang api

## Building the application

The following assumes that you are using a MacBook. You may to tweak this for Linux and Windows.

1. Clone the repo
2. `brew install go`
3. `brew install cobra`
4. `git checkout develop`
2. `brew install goreleaser`
5. `mkdir -p svang/api`
6. `cobra init svang/api --pkg-name github.com/svang/api`
7. `cd svang/api`
8. Create run and version command `cobra add run` and `cobra add version`

Define module path
```
go mod init github.com/svang/api
go install
go run main.go
go run main.go run --> run command
```

Make sure your `PATH` env variable has `go/bin` included (`export PATH=$PATH:$(go env GOPATH)/bin)`)


3. `make install` to test the tool on your MacBook or Windows laptop.
4. `make build-linux` to build the GO binary for Linux.
5. `make send-linux` to scp the tool to your target RHEL/CentOS machine. You need to modift the Makefile.
6. `upload.sh` - used by the developers of this tool to release the binary to GitHub.

