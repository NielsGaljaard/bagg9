build:
	@for OS in linux ; do \
		for arch in amd64 ; do \
			GOOS=$$OS GOARCH=$$arch go build  -o dist/release/$$arch/$$OS-$$arch-bagg9 cmd/bagg9/main.go ; \
		done \
	done 
binaries:
	@for OS in windows linux ; do \
		for arch in amd64 arm arm64; do \
			GOOS=$$OS GOARCH=$$arch go build  -o dist/release/$$arch/$$OS-$$arch-bagg9 cmd/bagg9/main.go ; \
		done \
	done 
test:
	go test -v ./...
