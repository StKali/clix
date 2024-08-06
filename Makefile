PROGRAM=clix
VERSION=v0.1.0
AUTHOR=clarkmonkey@163.com
COMMIT_ID=$(shell git rev-parse HEAD)
CHANGELOG=$(shell git log -1 --pretty=%s)
LICENSE=MIT
DATE=$(shell date '+%Y-%m-%d %H:%M:%S')
PRINT=@echo
REMOVE=rm -f

LDFLAGS_PROGRAM=-X 'github.com/stkali/clix/cmd.Program=$(PROGRAM)'
LDFLAGS_VERSION=-X 'github.com/stkali/clix/cmd.Version=$(VERSION)'
LDFLAGS_LICENSE=-X 'github.com/stkali/clix/cmd.License=$(LICENSE)'
LDFLAGS_BUILD=-X 'github.com/stkali/clix/cmd.Build=$(DATE)'
LDFLAGS_AUTHOR=-X 'github.com/stkali/clix/cmd.Author=$(AUTHOR)'
LDFLAGS_COMMIT_ID=-X 'github.com/stkali/clix/cmd.CommitID=$(COMMIT_ID)'
LDFLAGS_CHANGELOG=-X 'github.com/stkali/clix/cmd.ChangeLog=$(CHANGELOG)'
GO_COMPILE_LDFLAGS += $(LDFLAGS_PROGRAM)
GO_COMPILE_LDFLAGS += $(LDFLAGS_VERSION)
GO_COMPILE_LDFLAGS += $(LDFLAGS_LICENSE)
GO_COMPILE_LDFLAGS += $(LDFLAGS_BUILD)
GO_COMPILE_LDFLAGS += $(LDFLAGS_AUTHOR)
GO_COMPILE_LDFLAGS += $(LDFLAGS_COMMIT_ID)
GO_COMPILE_LDFLAGS += $(LDFLAGS_CHANGELOG)

build:
	go build -o bin/$(PROGRAM) -ldflags="-s $(GO_COMPILE_LDFLAGS)" main.go
	$(PRINT) Successfully build $(PROGRAM).

pybuild:
	@./build.py

test:
	go test ./... -v coverprofile=cover.out
	$(PRINT) "Successfully run $(PROGRAM) test."
	go tool cover --html=cover.out -o coverage.html
	$(PRINT) "Successfully render test coverage page: coverage.html"

clean:
	$(REMOVE) cover.out coverage.html
