module github.com/GrooveCommunity/zabbix-proxy-paygo

go 1.16

require (
	github.com/GrooveCommunity/glib-cloud-storage v0.0.0
	github.com/GrooveCommunity/glib-noc-event-structs v0.0.0
	github.com/gorilla/mux v1.8.0
)

replace (
    github.com/GrooveCommunity/glib-noc-event-structs v0.0.0 => /go/src/github.com/GrooveCommunity/glib-noc-event-structs
    github.com/GrooveCommunity/glib-cloud-storage v0.0.0 => /go/src/github.com/GrooveCommunity/glib-cloud-storage
)
