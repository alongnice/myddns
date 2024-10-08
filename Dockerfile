# 构建阶段
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && make clean build


# 构建结束，构建服务
# FROM golang
FROM alpine
LABEL name=myddns
LABEL url=https://github.com/alongnice/myddns

WORKDIR /app
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
    && apk add --no-cache tzdata
ENV TZ=Asia/Shanghai
# 添加时区
COPY --from=builder /app/myddns /app/myddns
EXPOSE 12138
ENTRYPOINT /app/myddns
CMD ["-l", ":9876", "-f", "300"] 
#可选参数