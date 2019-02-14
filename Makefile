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
		--project ${GCP_PROJECT} \
		--tag gcr.io/${GCP_PROJECT}/kws:latest

# DEPLOYMENT

secrets:
	kubectl create secret generic kws \
		--from-literal=KNOWN_PUBLISHER_TOKEN=${KNOWN_PUBLISHER_TOKEN}

secrets-clean:
	kubectl delete secret kws

deployment:
	kubectl apply -f deployments/service.yaml

nodeployment:
	kubectl delete -f deployments/service.yaml

# DEMO

event:
	curl -H "Content-Type: application/json" \
		 -X POST --data "{ \
			\"specversion\": \"0.2\", \
			\"type\": \"tech.knative.event.write\", \
			\"source\": \"https://knative.tech/test\", \
			\"id\": \"id-0000-1111-2222-3333-4444\", \
			\"time\": \"2019-01-11T17:31:00Z\", \
			\"contenttype\": \"text/plain\", \
			\"data\": \"My message content\" \
		}" \
		http://localhost:8080/?token=${KNOWN_PUBLISHER_TOKEN}
