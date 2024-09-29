# go-rest-api-frameworks

### Benchmarking results

**gin:** 

	wrk -t10 -c500 -d45s http://localhost:8080/ping
	Running 45s: 10 threads and 500 connections

	Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    30.07ms  142.52ms   1.59s    97.62%
    Req/Sec     6.48k     1.04k   25.45k    77.34%

  	2813442 requests in 45.10s, 378.32MB read
  	Socket errors: connect 0, read 654, write 8, timeout 0

	Requests/sec: 62385.67
	Transfer/sec: 8.39MB

**fasthttp:**

	wrk -t10 -c500 -d45s http://localhost:8080/ping

	Running 45s: 10 threads and 500 connections
  	
	Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     7.65ms   18.20ms 252.65ms   90.17%
    Req/Sec    19.38k     8.75k   72.70k    74.21%
  	
	8535571 requests in 45.08s, 1.14GB read
  	Socket errors: connect 0, read 631, write 0, timeout 0

	Requests/sec: 189354.56
	Transfer/sec: 26.00MB

**fiber:**

	wrk -t10 -c500 -d45s http://localhost:8080/ping

	Running 45s: 10 threads and 500 connections
  
  	Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     9.62ms   20.18ms 300.25ms   88.78%
    Req/Sec    18.59k    11.27k   76.69k    70.75%
  
  	8120495 requests in 45.09s, 0.95GB read
  	Socket errors: connect 0, read 566, write 0, timeout 0
	
	Requests/sec: 180092.97
	Transfer/sec: 21.64MB

**echo:**

	wrk -t10 -c500 -d45s http://localhost:8080/ping

	Running 45s: 10 threads and 500 connections
  
  	Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.39ms    4.67ms  88.50ms   84.38%
    Req/Sec     9.88k     2.65k   29.60k    73.13%
  
  	4419550 requests in 45.09s, 535.28MB read
  	Socket errors: connect 0, read 664, write 0, timeout 0

	Requests/sec:  98025.04
	Transfer/sec:     11.87MB

### Summary
	gin:      62385.67 requests/s & 8.39MB Transfer/s
	echo:     98025.04 requests/s & 11.87MB Transfer/s
	fiber: 	  180092.97 requests/s & 21.64MB Transfer/s
	fasthttp: 189354.56 requests/s & 26.00MB Transfer/s
