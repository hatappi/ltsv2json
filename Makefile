depend:
	go get -u github.com/tcnksm/ghr
	go get -u golang.org/x/crypto/ssh/terminal

build:
	go build -o pkg/ltsv main.go

build_crosscompile:
	./bin/crosscompile

fmt:
	go fmt

run:
	go run main.go

travis_deploy:
	./bin/travis_deploy
