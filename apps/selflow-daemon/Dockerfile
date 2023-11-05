###############################
# DOCKER START STAGE
###############################
FROM golang:1.21-bookworm
WORKDIR /usr/src/goapp/
USER ${USER}
COPY ./go.mod .
COPY ./go.sum .

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
# DOCKER INSTALL
###############################
RUN go mod download \
  && go mod tidy \
  && go mod verify

###############################
# DOCKER BUILD STAGE
###############################
COPY . .
RUN go build -o main ./apps/selflow-daemon

###############################
# DOCKER FINAL STAGE
###############################
EXPOSE 3000
CMD ["./main"]
