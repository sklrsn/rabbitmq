.DEFAULT_GOAL: all

MESSAGE ?= "change summary - $(shell git diff --name-only)"

.PHONY: all build up commit

all: build up commit

build:
	@docker compose build

up:
	@docker compose up

commit:
	@git add . && git commit -am $(MESSAGE) && git push
