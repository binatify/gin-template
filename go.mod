module github.com/binatify/gin-template

require (
	github.com/atarantini/ginrequestid v0.0.0-20180307004245-6d9eee666701
	github.com/binatify/gin-template/base/context v0.0.0
	github.com/binatify/gin-template/base/errors v0.0.0
	github.com/binatify/gin-template/base/logger v0.0.0
	github.com/binatify/gin-template/base/model v0.0.0
	github.com/binatify/gin-template/base/runmode v0.0.0
	github.com/binatify/gin-template/base/runmodegin v0.0.0
	github.com/gin-gonic/contrib v0.0.0-20181101072842-54170a7b0b4b // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/golib/assert v0.0.0-20170825111607-0306abba9bd3
	github.com/gorilla/sessions v1.1.3 // indirect
	github.com/json-iterator/go v1.1.5 // indirect
	github.com/kidstuff/mongostore v0.0.0-20181113001930-e650cd85ee4b // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b // indirect
	github.com/sirupsen/logrus v1.2.0
)

replace (
	github.com/binatify/gin-template/base/context => ./base/context
	github.com/binatify/gin-template/base/errors => ./base/errors
	github.com/binatify/gin-template/base/logger => ./base/logger
	github.com/binatify/gin-template/base/model => ./base/model
	github.com/binatify/gin-template/base/runmode => ./base/runmode
	github.com/binatify/gin-template/base/runmodegin => ./base/runmodegin
)
