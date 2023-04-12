VERSION := "0.0.1"

build:
	go build -o utils -ldflags "-X github.com/msknapp/command-line-utilities/cmd.version=$(VERSION)" main.go
	[ ! -d "$${GOBIN}" ] || mv utils "$${GOBIN}/utils"