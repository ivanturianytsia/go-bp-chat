FROM alpine

RUN apk --update upgrade && \
    apk --update add ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

RUN mkdir /chat
RUN mkdir /chat/avatars

WORKDIR /chat

ADD ./bin/chat_alpine /chat/chat
ADD /templates/ /chat/templates/
ADD /assets/ /chat/assets/

CMD ["/chat/chat","-addr=:8080"]
