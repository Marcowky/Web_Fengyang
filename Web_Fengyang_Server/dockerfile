FROM golang

WORKDIR "/app"

COPY . .

ENV GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN go build -o serve

CMD [ "./serve" ]