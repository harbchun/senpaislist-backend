FROM golang:1.16.4 AS development

RUN apt update && apt upgrade -y && \
    apt install -y git \
    make openssh-client

WORKDIR /app 

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD air

# build stage
FROM golang:alpine AS builder

COPY . /app

WORKDIR /app

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /app/main .

# final stage
FROM golang:alpine as staging

COPY --from=builder /app/migrations /go/migrations
COPY --from=builder /app/seed /go/seed
COPY --from=builder /app/scripts /go/scripts
COPY --from=builder /app/Makefile /go/

COPY --from=builder /dist/main /

# Command to run
ENTRYPOINT ["/main"]

# final stage
FROM golang:alpine as production

COPY --from=builder /app/migrations /go/migrations
COPY --from=builder /app/seed /go/seed
COPY --from=builder /app/scripts /go/scripts
COPY --from=builder /app/Makefile /go/
COPY --from=builder /app/go.mod /go/
COPY --from=builder /app/go.sum /go/

COPY --from=builder /dist/main /

# Command to run
ENTRYPOINT ["/main"]
