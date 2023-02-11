dc := docker compose -f ./docker-compose.yml

up:
	$(dc) up -d

down:
	$(dc) down

restart:
	$(dc) restart

reup:
	@make down
	@make up

rm:
	$(dc) down --rmi all

logs:
	$(dc) logs -f

app:
	$(dc) exec app /bin/sh

test:
	$(dc) exec app go test -v ./...

format:
	$(dc) exec app go fmt ./...

.PHONY:	up down restart reup rm logs app test format

#------------------------------------------------------------------------------

.DEFAULT_GOAL=build

PROJECT?=$(shell basename $(shell pwd))

GO_BUILD_DIR=build
.PHONY: build

build:
	make clean
	mkdir -p $(GO_BUILD_DIR)
	go build -o $(GO_BUILD_DIR) -v

run:
	make clean
	make build
	cd $(GO_BUILD_DIR)
	./$(GO_BUILD_DIR)/$(PROJECT)

clean:
	rm -rf build

#------------------------------------------------------------------------------

precommit:
	echo "#!/bin/sh\n\
	files=\$$(git diff --cached --name-only | grep -E '\\.(go)$$')\n\
	\n\
	if [ -n \"\$$files\" ]; then\n\
		echo \"Formatting Go files...\"\n\
		gofmt -s -w \$$files\n\
		git add \$$files\n\
	fi" > .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit