module github.com/mchmarny/knative-ws-example

require (
	github.com/cloudevents/sdk-go v0.0.0-20190102195109-feec6e002535
	github.com/google/uuid v1.1.0
	github.com/mchmarny/myevents v0.0.0-20190129001603-4253a49ac80f
	golang.org/x/net v0.0.0-20190213061140-3a22650c66bd
)

// TODO: Remove when PR #35 and #37 land
replace github.com/cloudevents/sdk-go => github.com/n3wscott/sdk-go v0.0.0-20190301221545-90692f88c5ae
