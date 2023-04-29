FROM mariadb:latest
ADD ./skyshi.sql /etc/mysql/skyshi.sql
ENV MYSQL_ROOT_PASSWORD=M@utauaja982
ENV MYSQL_DATABASE=skyshi
RUN sed -i 's/MYSQL_DATABASE/'skyshi'/g' /etc/mysql/skyshi.sql
RUN cp /etc/mysql/skyshi.sql /docker-entrypoint-initdb.d
CMD ["mysqld"]

FROM golang:1.20
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /usr/local/bin/app ./
CMD ["app"]

ENV MYSQL_HOST=$MYSQL_HOST
ENV MYSQL_USER=$MYSQL_USER
ENV MYSQL_PASSWORD=$MYSQL_PASSWORD
ENV MYSQL_DBNAME=$MYSQL_DBNAME