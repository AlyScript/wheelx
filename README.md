# wheelx (nginx's little cousin)

Run with

`go run main.go`

By default, the host `192.168.1.192` is used. This likely will not work for you, so you should use your own. To do this, you can use the `-i` flag

`go run main.go -i <your_host>`

By default, wheelx uses port 8888. To change this, you should use the `-p` option

`go run main.go -p <port>`
