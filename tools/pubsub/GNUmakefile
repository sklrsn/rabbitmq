.DEFAULT_GOAL=all

all: compile
.PHONY: all compile produce consume

compile:
	@go build -o bin/pubsub .

produce: compile
	@./bin/pubsub --exchange=logs --frequency=3 --routingkey=9 --workers=4 \
		--host=localhost --secret=guest --user=guest --mode=produce

consume: compile
	@./bin/pubsub --queue=logs.01 --frequency=3  --mode=consume --workers=4 \
		--host=localhost --secret=guest --user=guest

clean:
	@rm -rf bin/