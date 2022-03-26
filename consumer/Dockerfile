FROM maven:3.8-openjdk-17 as maven

WORKDIR /workspace

ADD . /workspace/

RUN mvn clean compile install package assembly:single

RUN mkdir -p /opt/rabbitmq/

RUN cp target/rabbitmq-consumer-1.0-jar-with-dependencies.jar /opt/rabbitmq/rmq-consumer-1.0.jar

FROM openjdk:17-bullseye

COPY --from=maven /opt/rabbitmq/rmq-consumer-1.0.jar /opt/rabbitmq/rmq-consumer-1.0.jar

CMD [ "java", "-jar", "/opt/rabbitmq/rmq-consumer-1.0.jar" ]