FROM golang:1.16.4
# RUN mkdir /app
# ADD . /app/
# WORKDIR /app/gql-server
# RUN go build -o main .
# EXPOSE 8080
# CMD ["/app/gql-server/main"]

RUN apt update && apt upgrade -y && \
    apt install -y git \
    make openssh-client

WORKDIR /app 

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD air
