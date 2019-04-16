FROM golang:latest as build
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app
RUN go build -o main ./snake

FROM iron/go
RUN mkdir /app
COPY --from=build /app/main /app/
CMD ["/app/main"]
