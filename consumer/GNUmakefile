.DEFAULT_GOAL: all

all: clean compile install package assembly
.PHONY: clean install package assembly compile

clean:
	@mvn clean

compile:
	@mvn compile

install:
	@mvn install

package:
	@mvn package

assembly:
	@mvn assembly:single

run:
	@java -jar target/rabbitmq-consumer-1.0-SNAPSHOT-jar-with-dependencies.jar

docker:
	docker build -t maven-builder:latest .
	docker run maven-builder:latest