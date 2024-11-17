FROM ibmjava:latest AS java-dev

RUN apt update && apt install -y \
  maven

RUN apt upgrade -y



FROM mcr.microsoft.com/dotnet/sdk:latest AS cs-dev

RUN apt update
Run apt upgrade -y
