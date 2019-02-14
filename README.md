# myevents

Knative Cloud Events Gateway and Viewer

## Demo

https://myevents.default.knative.tech/

## Setup

Setup assumes you already have `gcloud` installed. If not, see [Installing Google Cloud SDK](https://cloud.google.com/sdk/install)


> This readme is still a bit of work in progress so if you are finding something missing do take a look at the [Makefile](https://github.com/mchmarny/myevents/blob/master/Makefile)

### Knative URL

To avoid the kind of chicken and an egg situation we are going to first define the `URL` that your application will have when you publish it on Knative. Knative uses convention to build serving URL by combining the deployment name (e.g. `myevents`), namespace name (e.g. `default`), and the pre-configured domain name (e.g. `knative.tech`). The resulting URL, assuming you already configured SSL, should look something like this:

```shell
https://myevents.default.knative.tech
```

### Google OAuth Credentials

In your Google Cloud Platform (GCP) project console navigate to the Credentials section. You can use the search bar, just type `Credentials` and select the option with "API & Services". To create new OAuth credentials:

* Click “Create credentials” and select “OAuth client ID”
* Select "Web application"
* Add authorized redirect URL at the bottom using the fully qualified domain we defined above and appending the `callback` path:
 * `https://myevents.default.knative.tech/auth/callback`
* Click create and copy both `client id` and `client secret`
* CLICK `OK` to save

For ease of use, export the copied client `id` as `MYEVENTS_OAUTH_CLIENT_ID` and `secret` as `MYEVENTS_OAUTH_CLIENT_SECRET` in your environment variables (e.g. ~/.bashrc or ~/.profile)

> You will also have to verify the domain ownership. More on that [here](https://support.google.com/cloud/answer/6158849?hl=en#authorized-domains)


### App Deployment

> TODO: Provide public image deployment instructions

To deploy the `myevents` are are going to:

* [Build the image](#build-the-image)
* [Configure Knative](#configure-knative)
* [Deploy Service](#deploy-service)

#### Build the image

Quickest way to build your service image is through [GCP Build](https://cloud.google.com/cloud-build/). Just submit the build request from within the `myevents` directory:

```shell
gcloud builds submit \
    --project ${GCP_PROJECT} \
	--tag gcr.io/${GCP_PROJECT}/myevents:latest
```

The build service is pretty verbose in output but eventually you should see something like this

```shell
ID           CREATE_TIME          DURATION  SOURCE                                   IMAGES                      STATUS
6905dd3a...  2018-12-23T03:48...  1M43S     gs://PROJECT_cloudbuild/source/15...tgz  gcr.io/PROJECT/myevents SUCCESS
```

Copy the image URI from `IMAGE` column (e.g. `gcr.io/PROJECT/myevents`).

#### Configure Knative

Before we can deploy that service to Knative, we just need to create Kubernetes secrets and update the `deploy/server.yaml` file

```shell
kubectl create secret generic myevents \
    --from-literal=OAUTH_CLIENT_ID=$MYEVENTS_OAUTH_CLIENT_ID \
    --from-literal=OAUTH_CLIENT_SECRET=$MYEVENTS_OAUTH_CLIENT_SECRET
```

Now in the `deploy/server.yaml` file update the `GCP_PROJECT_ID`

```yaml
    - name: GCP_PROJECT_ID
      value: "enter your project ID here"
```

And the external URL of your which we defined at the begining of this readme in [###knative-url] section.

```yaml
    - name: EXTERNAL_URL
      value: "https://myevents.default.DOMAIN.COM"
```

#### Deploy Service

Once done updating our service manifest (`deploy/server.yaml`) you are ready to deploy it.

```shell
kubectl apply -f deployments/service.yaml
```

The response should be

```shell
service.serving.knative.dev "myevents" configured
```

To check if the service was deployed successfully you can check the status using `kubectl get pods` command. The response should look something like this (e.g. Ready `3/3` and Status `Running`).

```shell
NAME                                          READY     STATUS    RESTARTS   AGE
myevents-00002-deployment-5645f48b4d-mb24j     3/3       Running   0          4h
```

You should be able to test the app now in browser using the `URL` you defined above.

## Disclaimer

This is my personal project and it does not represent my employer. I take no responsibility for issues caused by this code. I do my best to ensure that everything works, but if something goes wrong, my apologies is all you will get.

