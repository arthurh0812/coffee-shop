FROM golang:1.16.5

ENV GO111MODULE=on

WORKDIR /app

RUN useradd -m arthur && groupadd docker && usermod -a -G docker arthur

COPY --chown=arthur:docker go.mod .
COPY --chown=arthur:docker go.sum .

RUN go mod vendor && go mod verify

COPY --chown=arthur:docker . .

ENV PORT=8080
ENV HOST=localhost

RUN go build -o bin .

CMD [ "/app/bin" ]