FROM alpine
LABEL authors="socoldkiller"


RUN apk update && \
    apk add curl

CMD sh