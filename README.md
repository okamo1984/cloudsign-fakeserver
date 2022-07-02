# Go Echo API Server for cloudsignfakeserver

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This server was generated by the [openapi-generator]
(https://openapi-generator.tech) project.
By using the [OpenAPI-Spec](https://github.com/OAI/OpenAPI-Specification) from a remote server, you can easily generate a server stub.
-

To see how to make this your own, look here:

[README](https://openapi-generator.tech)

- API version: 0.19.0
- Build date: 2022-07-01T23:33:47.129973+09:00[Asia/Tokyo]

### Running the server

To run the server, follow these simple steps:

```
go mod download
go build -o app
```

To run the server in a docker container
```
docker build --network=host -t cloudsignfakeserver .
```

Once the image is built, just run
```
docker run --rm -it cloudsignfakeserver
```

### Known Issue

TBA