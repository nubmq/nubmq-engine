all: build

build:
	go run connectionHandler.go init.go hasher.go setter.go getter.go minimal_engine.go resizer.go ShardUtils.go

test:
	go test