.PHONY: app client service

# DEV
test:
	go test ./... -v

test-event:
	go test --run NameOfTest ./pkg/handler/event/ -v

run:
	go run ./cmd/*.go

deps:
	go mod tidy

# BUILD

image:
	gcloud builds submit \
		--project knative-samples \
		--tag gcr.io/knative-samples/kws:latest

# DEPLOYMENT

secrets:
	kubectl create secret generic kws -n demo \
		--from-literal=KNOWN_PUBLISHER_TOKEN=${KNOWN_PUBLISHER_TOKEN}

secrets-clean:
	kubectl delete secret kws

deployment:
	kubectl apply -f deployment/service.yaml

nodeployment:
	kubectl delete -f deployment/service.yaml

# DEMO

event:
	curl -H "Content-Type: application/json" \
		 -X POST --data "{ \
			\"specversion\": \"0.2\", \
			\"type\": \"github.com.mchmarny.knative-ws-example.message\", \
			\"source\": \"https://github.com/mchmarny/knative-ws-example\", \
			\"id\": \"6CC459AE-D75D-4556-8C14-CD1ED5D95AE7\", \
			\"time\": \"2019-02-13T17:31:00Z\", \
			\"contenttype\": \"text/plain\", \
			\"data\": \"This is my sample message\" \
		}" \
		http://localhost:8080/?token=${KNOWN_PUBLISHER_TOKEN}

