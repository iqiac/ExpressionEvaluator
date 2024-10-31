FROM ibmjava:latest AS dev

RUN apt update && \
  apt install -y \
  maven

RUN apt upgrade -y

WORKDIR /app
