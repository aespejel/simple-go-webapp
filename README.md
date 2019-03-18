# simple-go-webapp
Build it using:
```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
```
Run it with docker using:
```
docker build -t demo-app .
docker run --rm -p 8080:8080 --name demoapp demo-app
```
