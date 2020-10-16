build:
	go build -ldflags "-X github.com/svang/svangapi/cmd.gitCommitHash=`git rev-parse HEAD` \
	   -X github.com/svang/svangapi/cmd.buildTime=`date -u '+%Y-%m-%d--%H:%M:%S%p'` \
	   -X github.com/svang/svangapi/cmd.gitBranch=`git branch --show-current` \
	   -X github.com/svang/svangapi/cmd.tagVersion=`git describe --tags --long`" \
	   -o bin/svangapi main.go
build-w-clean: clean build
build-linux: # example: make build-linux DB_PATH=/dir/to/db
	env GOOS=linux GOARCH=amd64 go build -ldflags \
	"-X github.com/svang/svangapi/internal/command.stateDBPathFromEnv=/tmp \
	-X github.com/svang/svangapi/internal/command.logDirPathFromEnv=/var/log/svang/api \
	-X github.com/svang/svangapi/cmd.gitCommitHash=`git rev-parse HEAD` \
	-X github.com/svang/svangapi/cmd.buildTime=`date -u '+%Y-%m-%d--%H:%M:%S%p'` \
	-X github.com/svang/svangapi/cmd.gitBranch=`git branch --show-current` \
	-X github.com/svang/svangapi/cmd.tagVersion=`git describe --tags --long`" \
	-o bin/svangapi main.go
send-linux: install build-linux
	scp bin/svangapi root@naanstop.zinox.com:/usr/local/bin
clean:
	rm -f bin/svang
	rm -rf logs
	rm -f *.log
install: clean
	go install
run-help:
	go run main.go --help
run-server: install
	svangapi server
# save github token in an environment variable export GITHUB_TOKEN="token"
add-tag:
ifeq (,$(findstring v,$(tag)))
	@echo "error : tag needs to be of format v0.x.x. Usage --> make upload tag=v0.x.x"
	@echo
	exit 1
endif
	git fetch
	git tag $(tag)
	git push origin --tags
upload: add-tag install build-linux #make upload tag=v0.x.x, install --> brew install goreleaser
	goreleaser --rm-dist
test:
	go test -v ./...
