# wheelx (nginx's little cousin)

```yaml
wheelx is a very simple HTTP server which you can deploy on almost any device. It is extremely lightweight and insecure (so i recommend that you take great care with this) by design.
```

Run with

```go
go run main.go
```

By default, the host `192.168.1.192` is used. This likely will not work for you, so you should use your own. To do this, you can use the `-i` option

`go run main.go -i <your_host>`

By default, wheelx uses port 8888. To change this (you'll probably want to do this), you should use the `-p` option

`go run main.go -p <port>`

wheelx benefits from several helpful endpoints which provide useful data.

the `stats` endpoint will show the number of requests that the server has received. Note that this value is reset with the server.
