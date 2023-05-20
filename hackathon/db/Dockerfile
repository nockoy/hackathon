FROM golang:1.20.2 as build
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian11
COPY --from=build /go/bin/app /
CMD ["/app"]

#docker build -t hackathon hackathon
#docker run -p 8000:8000 hackathon