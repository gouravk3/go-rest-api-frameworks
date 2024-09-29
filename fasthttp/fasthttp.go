package main

import (
	"github.com/valyala/fasthttp"
)

func main() {
	fasthttp.ListenAndServe(":8080", func(c *fasthttp.RequestCtx) {
		c.SetContentType("application/json")
		c.SetStatusCode(fasthttp.StatusOK)
		c.SetBody([]byte(`{"message":"pong"}`))
	})
}
