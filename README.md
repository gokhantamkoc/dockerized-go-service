# Dockerized Go Service

## Commands to Build and Run dockerized-go-service

Syntax:
```
$ docker image build /path/to -t imageName:version -t imageName:revision
$ docker run -d --name container_name image_name
```

```
$ docker image build . -t softwaresanctuary/myservice:1.0.0 -t softwaresanctuary/myservice:latest 
$ docker run -d -p 9000:8080 --name myservice-inst1 softwaresanctuary/myservice:latest 
```

## Best Commands for Cloud Environments

From github URL:

```
$ docker build https://github.com/docker/rootfs.git#container:docker
```

From tarball:

```
$ docker build http://server/context.tar.gz
```