FROM golang:1.22.2 AS BuildStage

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o rgr ./cmd

FROM alpine:latest 

COPY --from=BuildStage /app/rgr /app/rgr

COPY --from=BuildStage /app/base.json base.json

RUN apk --no-cache add ca-certificates tzdata

ENTRYPOINT [ "/app/rgr" ]



