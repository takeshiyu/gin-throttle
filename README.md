# gin-throttle ![MIT license](https://poser.pugx.org/laravel/framework/license.svg)

The `gin-throttle` is a rate limiter for [Gin framework](https://github.com/gin-gonic/gin). 

Original credit to [golang.org/x/time/rate](https://godoc.org/golang.org/x/time/rate), this repo makes it work with `gin.HandlerFunc` in [Gin framework](https://github.com/gin-gonic/gin).

## Usage

Download and install it：

```
go get github.com/gin-gonic/gin
```

Import it in your code：

```
import "github.com/s12i/gin-throttle"
```

## Example

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/s12i/gin-throttle"
)

func main() {
	router := gin.Default()

	maxEventsPerSec := 1000
	maxBurstSize := 20
	router.Use(middleware.Throttle(maxEventsPerSec, maxBurstSize))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
```

## Benchmark

```
❯ wrk -t12 -c400 -d20s http://localhost:8080/ping
Running 20s test @ http://localhost:8080/ping
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    34.19ms   28.57ms 465.71ms   70.47%
    Req/Sec     1.05k   200.17     1.69k    79.12%
  251406 requests in 20.07s, 22.54MB read
  Non-2xx or 3xx responses: 231620
Requests/sec:  12528.62
Transfer/sec:      1.12MB
```

## License

The `gin-throttle` is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).