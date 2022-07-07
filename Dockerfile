FROM golang:1.18-alpine

WORKDIR /app

COPY ./GoBlogCore .

RUN go mod download

RUN go build -o /fresh-blog-core

ENV HTTP_PORT 5000
ENV MONGO_URI uri
ENV DB_NAME dbname

EXPOSE 5000

CMD [ "/fresh-blog-core" ]
