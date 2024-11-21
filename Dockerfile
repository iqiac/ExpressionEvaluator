FROM ibmjava:latest AS java-dev
RUN apt update && apt install -y maven && apt upgrade -y



FROM mcr.microsoft.com/dotnet/sdk:latest AS cs-dev
RUN apt update && apt upgrade -y



FROM golang:latest AS go-dev
RUN apt update && apt upgrade -y