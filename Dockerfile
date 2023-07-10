FROM golang:1.20 as go_builder
RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o go_link_preview_generator .

FROM alpine as go_link_preview_generator
COPY --from=go_builder /app/go_link_preview_generator /go_link_preview_generator

ENTRYPOINT [ "/go_link_preview_generator" ]