FROM golang:1.17 as base

WORKDIR /app
COPY go.mod .
COPY go.sum .

FROM base as source
COPY ./*.go /app/

RUN ls /app/
RUN go build -o freego_api .
EXPOSE 1379
CMD [ "./freego_api" ]