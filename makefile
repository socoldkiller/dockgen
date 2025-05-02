
APP_NAME := dockgen

GRAMMAR_DIR := grammar
GEN_DIR := parser
CMD_DIR := cmd
PKG_DIR := ir

ANTLR_GRAMMAR := $(PWD)/pkg/analysis/ir/antlr4/Bash.g4

BUILD_DIR := bin
BIN := $(BUILD_DIR)/$(APP_NAME)

DEBUG_CONTAINER = debug-container

GO := go

CONTAINER_NAME = dockgen-container
.PHONY: all antlr build clean run

all: antlr build


antlr:
	cd $(PWD)/pkg/analysis/types/antlr4 && \
	antlr -Dlanguage=Go -visitor Bash.g4


build-debugger-container:
	docker build -t $(CONTAINER_NAME) -f $(DEBUG_CONTAINER)/Dockerfile $(DEBUG_CONTAINER)


debugger: build-debugger-container
	docker run --privileged -itd -v $(PWD):/app  --net host --name $(CONTAINER_NAME) $(CONTAINER_NAME)
	docker exec -it --workdir /app $(CONTAINER_NAME) /bin/sh


build:
	@mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BIN) .

release:
	@mkdir -p $(BUILD_DIR)
	$(GO) build -ldflags '-w -s' -gcflags '-l' -o $(BIN) .


run:
	$(GO) build



clean:
	@rm -rf $(BUILD_DIR) $(GEN_DIR)/*.go
	@docker rm -f $(CONTAINER_NAME)

