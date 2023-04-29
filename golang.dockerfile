FROM golang:1.20

# init database
WORKDIR /usr/src/app
COPY . .
RUN go mod download && go mod verify

RUN cd migrate
RUN go mod download && go mod verify

# init aplikasi
RUN cd /usr/src/app
RUN go build -v -o /usr/local/bin/app ./
EXPOSE 3030

ENTRYPOINT [ "/bin/sh", "/usr/src/app/entrypoint.sh" ]