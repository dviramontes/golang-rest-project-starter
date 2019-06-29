FROM golang:1.12.6-stretch

RUN go get github.com/cespare/reflex

COPY reflex.conf /reflex.conf

COPY config.yml /config.yml

EXPOSE 3000

ENTRYPOINT ["reflex", "-c", "/reflex.conf"]
