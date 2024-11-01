FROM golang:1.23

WORKDIR /app

ENV PORT 8080
ENV USER imageuser

RUN useradd -m $USER && \
    chown -R $USER:$USER /app  && \
    chown -R $USER:$USER /go


COPY go.mod go.sum ./
RUN go mod download

RUN mkdir -p /data && chown -R $USER:$USER /data

EXPOSE $PORT

USER $USER
ENTRYPOINT ["go", "run", "main.go"]
