module main

replace (
	bean => ./bean
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.36.0

	controller => ./controller

	db => ./db

	golang.org/x/build => github.com/golang/build v0.0.0-20190307215223-c78805dbabc8

	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190228161510-8dd112bcdc25

	golang.org/x/exp => github.com/golang/exp v0.0.0-20190306152737-a1d7652674e8

	golang.org/x/lint => github.com/golang/lint v0.0.0-20190301231843-5614ed5bae6f

	golang.org/x/net => github.com/golang/net v0.0.0-20190301231341-16b79f2e4e95

	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190226205417-e64efc72b421

	golang.org/x/perf => github.com/golang/perf v0.0.0-20190306144031-151b6387e3f2

	golang.org/x/sync => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6

	golang.org/x/sys => github.com/golang/sys v0.0.0-20190308023053-584f3b12f43e

	golang.org/x/text => github.com/golang/text v0.3.0

	golang.org/x/time => github.com/golang/time v0.0.0-20181108054448-85acf8d2951c

	golang.org/x/tools => github.com/golang/tools v0.0.0-20190307163923-6a08e3108db3

	google.golang.org/api => github.com/googleapis/google-api-go-client v0.1.0

	google.golang.org/appengine => github.com/golang/appengine v1.4.0

	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190307195333-5fe7a883aa19

	google.golang.org/grpc => github.com/grpc/grpc-go v1.19.0

	handlers => ./handlers

	router => ./router

)

require (
	bean v0.0.0
	controller v0.0.0
	db v0.0.0
	github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/golang/protobuf v1.3.0 // indirect
	github.com/json-iterator/go v1.1.5 // indirect
	github.com/mattn/go-isatty v0.0.6 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43 // indirect
	gopkg.in/gcfg.v1 v1.2.3
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
	handlers v0.0.0
	router v0.0.0
)
