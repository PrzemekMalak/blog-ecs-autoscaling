FROM golang:1.15.3-alpine AS build

COPY main.go ./
RUN CGO_ENABLED=0 go build -o /bin/stress

FROM scratch
COPY --from=build /bin/stress /bin/stress

EXPOSE 8080

CMD ["/bin/stress"]