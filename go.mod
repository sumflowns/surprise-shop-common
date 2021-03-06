module github.com/sumflowns/toyo-src/book-ticket-common

go 1.13

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/gin-gonic/gin v1.5.0
	github.com/go-log/log v0.1.0
	github.com/go-redis/redis v6.15.6+incompatible // indirect
	github.com/goinggo/mapstructure v0.0.0-20140717182941-194205d9b4a9
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.11 // indirect
	github.com/kataras/golog v0.0.9 // indirect
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.17.1
	github.com/micro/micro v1.17.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/prometheus/client_golang v1.1.0
	github.com/rs/cors v1.7.0
	github.com/satori/go.uuid v1.2.0
	github.com/uber/jaeger-client-go v2.20.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
)

replace github.com/sumflowns/toyo-src/book-ticket-common => ../book-ticket-commo
