package test

const (
	TestConfig = `
http {
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent"'

}
`
	TestLogs = `
61.165.46.190 - - [26/Dec/2020:06:00:53 +0800] "GET /ouc/dealer/ping_session.jsp HTTP/1.1" 200 98 "-" "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0)"
180.171.59.71 - - [26/Dec/2020:06:00:53 +0800] "GET /ouc/dealer/ping_session.jsp HTTP/1.1" 200 98 "-" "Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko"
157.55.39.150 - - [26/Dec/2020:06:00:54 +0800] "GET /carinfo_143812 HTTP/1.1" 302 0 "-" "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)"
203.208.60.101 - - [26/Dec/2020:06:00:55 +0800] "GET /car_1198367.html HTTP/1.1" 200 12011 "-" "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"
42.120.161.124 - - [26/Dec/2020:06:00:55 +0800] "GET /carinfo_155205 HTTP/1.1" 302 0 "-" "YisouSpider"
106.11.153.52 - - [26/Dec/2020:06:00:55 +0800] "GET / HTTP/1.1" 200 76935 "-" "YisouSpider"
`
)
