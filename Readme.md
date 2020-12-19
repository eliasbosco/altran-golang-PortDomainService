# Altran Golang PortDomainService

Following the requirements asked, this service is responsible to receive all request from the [ClientAPI](https://github.com/eliasbosco/altran-golang-ClientAPI) via [gRPC](https://github.com/grpc/grpc-go/blob/d79063fdde284ef7722591e56c72143eea59c256/examples/features/debugging/server/main.go#L45) and the ports data enveloped by protobuf binary protocol. The main routine put 3 or more gRPC servers up and it should be increased, depending on how much listeners the user need, inserting more addresses in .env file.

## Containers

All the service is dockerized and provided with docker compose script to facilitate the deploy step, all the commands are described below:

To bring the things on:

```
$ docker-compose up -d --build
```

After it's up, access:

(to response with port.json file results)

- [http://localhost:10000/ports?skip=0&limit=100](http://localhost:10000/ports?skip=0&limit=100)

(to response with sqlite results)

- http://localhost:10000/ports-db?port_id=ABCD
- http://localhost:10000/ports-db?skip=0&limit=100

To shut the service off:

```
$ docker-compose down
```

Notice: every environments variables are set on &quot;.env&quot; file, at the project root.

## Resources

All the computational resource are limited to 50M memory, even lessly as mentioned on the pre-requisits.
