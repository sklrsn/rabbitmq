.DEFAULT_GOAL: all

all: clean install package assembly
.PHONY: clean install package assembly

clean:
	@mvn clean

install:
	@mvn install

package:
	@mvn package

assembly:
	@mvn assembly:single

run:
	@ java -jar target/rabbitmq-consumer-1.0-SNAPSHOT-jar-with-dependencies.jar

