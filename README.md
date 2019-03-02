# knative-ws-example

Simple Knative example consisting of Web UI which uses WebSockets to display POST'ed Cloud Events.

![Knative WebSocket](/../master/static/img/after.png?raw=true "Knative WebSocket")

## Prerequisites

> Note, this sample requires Knative `v0.4` or better

Additionally, if you are want to build the demo you will need `gcloud` installed. See [Installing Google Cloud SDK](https://cloud.google.com/sdk/install)

## Setup

In configuring the The Knative WebSocket example we will deploy pre-built image of this application and use `curl` to post messages to that application, which if everything is configured correctly, will show on in the UI.

> Note, if you prefer to build your own image please see `Build` section bellow

### Token

To prevent just anyone publishing to your application we are going to configure `token` which will be shared between the application and all permitted clients. Let's start by creating a token

```shell
KNOWN_PUBLISHER_TOKEN=$(uuidgen)
echo $KNOWN_PUBLISHER_TOKEN
```

### Namespace

If you don't already have a `demo` namespace created

```shell
kubectl create ns demo
```

### Secret

Before deploying our application we will load the above token into the Knative cluster as a secret so that our application can use it


```shell
kubectl create secret generic kws -n demo \
	--from-literal=KNOWN_PUBLISHER_TOKEN=${KNOWN_PUBLISHER_TOKEN}
```

### Service

Once our namespace and secret is created, you can apply the  (`deployment/service.yaml`) to deploy the application

```shell
kubectl -n demo apply -f \
    https://raw.githubusercontent.com/mchmarny/knative-ws-example/master/deployment/service.yaml
```

If everything worked correctly you should be able to see the `kws` service listed as running

```shell
kubectl get pods -n demo
```

```shell
NAME                                      READY     STATUS      RESTARTS   AGE
kws-00001-deployment-74c64b8d6b-6dqds     3/3       Running     1          1m
```

## Demo

### UI

First navigate to the newly deployed application

http://kws.demo.YOUR-DOAIN.com/

The status message on the top fo the screen should say `Opening Connection`

### Client

Now that the application is deployed, you can use `curl` to post to its endpoint and the messages should show on the UI. For ease of demonstration we are going to post a simple text message with hard-coded values.

```shell
curl -H "Content-Type: text/plain" \
     -H "ce-specversion: 0.2" \
     -H "ce-type: github.com.mchmarny.knative-ws-example.message" \
     -H "ce-source: https://github.com/mchmarny/knative-ws-example" \
     -H "ce-id: $(uuidgen)" \
     -H "ce-time: $(date +%Y-%m-%dT%H:%M:%S:%Z)" \
     -H "ce-token: $(KNOWN_PUBLISHER_TOKEN)" \
     -X POST --data "My sample message" \
    https://kws.demo.knative.tech/v1/event
```

Every time you run that command, a new message should be added to the UI. Go ahead, change the `data` value from `This is my sample message` to something else and post it.

## Build

> Note, do this only if you want to build from source, otherwise use the pre-built image

Quickest way to build your service image is through [GCP Build](https://cloud.google.com/cloud-build/). Just submit the build request using the bellow command

```shell
gcloud builds submit \
    --project ${YOUR_GCP_PROJECT} \
	--tag gcr.io/${YOUR_GCP_PROJECT}/kws
```

The build service is pretty verbose in output but eventually you should see something like this

```shell
ID           CREATE_TIME          DURATION  SOURCE                                   IMAGES               STATUS
6905dd3a...  2018-12-23T03:48...  1M43S     gs://PROJECT_cloudbuild/source/15...tgz  gcr.io/PROJECT/kws   SUCCESS
```

Copy the image URI from `IMAGE` column (e.g. `gcr.io/PROJECT/kws`).

### Configure Knative

Now in the `deployment/service.yaml` file update the `image` URI to value of `IMAGE` column above

## Disclaimer

This is my personal project and it does not represent my employer. I take no responsibility for issues caused by this code. I do my best to ensure that everything works, but if something goes wrong, my apologies is all you will get.
