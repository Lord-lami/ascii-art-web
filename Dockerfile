FROM golang:1.25

WORKDIR /app

COPY banners/ ./

COPY . .

EXPOSE 8080

RUN go build -v -o bin/ascii-art-web

CMD [ "./bin/ascii-art-web" ]