FROM --platform=linux/amd64 alpine

ENV GIN_MODE=release
ENV WORK_DIR /work
ENV APP_BIN albumsvr

RUN apk add --no-cache libc6-compat curl tzdata
ENV TZ=Asia/Shanghai

RUN mkdir -p $WORK_DIR
COPY ./$APP_BIN $WORK_DIR
WORKDIR $WORK_DIR

CMD ./$APP_BIN