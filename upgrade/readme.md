# GO Crop HTTP Server + Redis 

## Description.

This is an Golang HTTP Server that is having 2 endpoints:

1. Check the status (/status)
2. Receive a image, crop it and upload the url of the modified image to a redis db.

## How to run it

### Docker redis + go in local ( TestMode )



### Docker-compose ( "Prod" Mode )

```bash
docker-compose build
docker-compose up
```

## How to test it.

1.
2.

## Why I choose the libraries, the workflow, etc.

* fmt:
* log:
* net/http:
* go-redis/redis:
* gorilla/mux:
* os:
* io:
* desintegration/imaging:
* image + image/color:
* math/rand:
* strconv:


### TODO

1. Upload the images to another directory + show static files .
2. TLS. (https://github.com/denji/golang-tls  This example will be enough)
3. Crop better the image (accept all image formats).
4. Implement it in K8s.
