container_test
==============

This is a simple HTTP server written in Go returning the hostname and IP of the
host the server runs on. It's meant to be running in a Docker container and
return the name and the IP of the container for testing purposes.


Usage
-----

Build container:

```shell
docker build -t container_test .
```

Run container:

```shell
docker run --rm --publish 8080:8080 --name container_test container_test
```

Test container:

```
curl localhost:8080
```


Author
------

Jiri Tyr


License
-------

MIT
