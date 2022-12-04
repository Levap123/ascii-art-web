#build
FROM golang:1.18.3 AS build

WORKDIR /app

COPY . .

RUN GOOS=linux go build -o main ./cmd/


#run 
FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=build /app/main .
COPY --from=build /app/templates templates
COPY --from=build /app/static static
COPY --from=build /app/src src  

EXPOSE 8080

USER root:root

ENTRYPOINT ["./main"]