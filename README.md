Fooling with go

footnotes

```
docker build --target dev . -t go
docker run -it -v ${PWD}:/work go sh
go version
```

```
docker build . -t videos
docker run -it videos get --help
```
