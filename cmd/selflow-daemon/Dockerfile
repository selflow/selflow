###############################
# DOCKER START STAGE
###############################
FROM golang:1.20-buster
WORKDIR /usr/src/goapp/
USER ${USER}
COPY ./go.mod .
COPY . .

###############################
# DOCKER ENVIRONMENT STAGE
###############################
ENV GO111MODULE="on" \
  GO_GC="off"

###############################
# DOCKER UPGRADE STAGE
###############################
RUN apt-get autoremove \
  && apt-get autoclean \
  && apt-get update --fix-missing \
  && apt-get upgrade -y \
  && apt-get install curl \
  build-essential -y

###############################
# DOCKER INSTALL & BUILD STAGE
###############################
RUN go mod download \
  && go mod tidy \
  && go mod verify \
  && go build -o main ./cmd/selflow-daemon

###############################
# DOCKER FINAL STAGE
###############################
EXPOSE 3000
CMD ["./main"]
