GitTag=$(shell git describe --tag --always)


LDFLAGS='-X 'github.com/jimyag/cmd-demo/version.version=$(GitTag)' -extldflags '-static''

default:
	go build -ldflags "$(LDFLAGS)" -o bin/cmd-demo .