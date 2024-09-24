# 构建阶段
FROM golang AS builder
WORKDIR /app
COPY . .
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go get -d -v . \
    && go install -v . \
    && go build -v .

# 构建结束，构建服务
# FROM golang
FROM ubuntu
WORKDIR /app
COPY --from=builder /app/myddns /app/myddns
EXPOSE 12138
ENTRYPOINT /app/myddns
LABEL NAME=myddns Version=0.0.1
