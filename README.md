### Usage

[Extract image metadata to JSON format using ImageMagick (screencast)](https://sysadmincasts.com/episodes/53-extracting-image-metadata-to-json-using-imagemagick)

![screenshot](https://raw.githubusercontent.com/jweissig/alpine-imagemagick-detect/master/templates/screenshot.png)

##### Build go app for linux (from mac)

```sh
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o web ./main.go
```

##### Build/push container

```sh
docker build -t jweissig/alpine-imagemagick-detect .
docker push jweissig/alpine-imagemagick-detect
```

##### Run container

https://hub.docker.com/r/jweissig/alpine-imagemagick-detect

```sh
docker pull jweissig/alpine-imagemagick-detect
docker run -it -p 5000:5000 --rm jweissig/alpine-imagemagick-detect
```

##### Connect to localhost

http://localhost:5000/

##### Debugging

```sh
docker ps | grep imagemagick
docker exec -it <CONTAINER ID> bash
convert example.png json:
```
