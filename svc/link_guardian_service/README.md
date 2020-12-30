# Link Check Service

The link check microservice. It uses a multi-stage [Dockerfile](Dockerfile) to generate a lean and mean image from SCRATCH that just includes the Go binary. The system has a CI/CD pipeline, but you also


## Build Docker image

```
$ docker build . -t vivsoft-platform/vivsoft-link-guardian:${VERSION}
```

## Push to Registry

This will go by default to DockerHub. Ensure you're logged in:

```
$ docker login
```

Then push the image:

```
$ docker push vivsoft-platform/vivsoft-link-guardian:${VERSION}
```

## Deploy to active Kubernetes cluster

To push a local cluster (microK8s, k3d, kind) ensure your kubectl is pointed to the right cluster (check the cluster context configuration) and type:

```
$ kubectl apply -f k8s
```






