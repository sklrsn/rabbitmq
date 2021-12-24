.DEFAULT_GOAL: all

MESSAGE ?= "change summary - $(shell git diff --name-only)"

.PHONY: all build up commit down

all: build up

build:
	@docker compose build

up:
	@docker compose up

down:
	@docker compose down

commit:
	@git add . && git commit -am $(MESSAGE) && git push