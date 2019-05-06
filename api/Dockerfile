FROM golang:latest

WORKDIR /go/src/api
COPY . .
#ENV GO111MODULE=on
#ENV GIT_TERMINAL_PROMPT=1
RUN go get github.com/gravityblast/fresh
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql

CMD ["fresh"]

