.DEFAULT_GOAL=all

all: compile
.PHONY: all compile produce consume

compile:
	@go build -o bin/pubsub .

produce:
	@./bin/pubsub --exchange=logs --frequency=3 --routing_key=# --mode=produce

consume:
	@./bin/pubsub --queue=logs.01 --frequency=3  --mode=consume