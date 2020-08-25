FROM golang
WORKDIR /go/src/github.com/stevenandrewcarter/terradex
COPY . .
# RUN go test -json ./...
RUN env GOOS=linux GOARCH=amd64 go build -v -o /output/terradex github.com/stevenandrewcarter/terradex/cmd/terradex

FROM centos
COPY --from=0 /output/terradex /app/terradex
# Add the config volume link
EXPOSE 8080
WORKDIR "app"
CMD ["./terradex", "server"]
