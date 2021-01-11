FROM golang:1.15 as build

WORKDIR /src/app

COPY . .

RUN bin/build

FROM gcr.io/distroless/base
COPY --from=build /bin/app /

EXPOSE 5000

CMD ["/app", "server"]
