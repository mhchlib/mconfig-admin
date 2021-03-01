VERSION=$(shell git describe --tags --always --dirty --dirty="")

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mconfig-admin cmd/mconfig-admin/main.go

image: build
	docker build -t dockerhcy/mconfig-admin:${VERSION}   .

push: image
	docker push dockerhcy/mconfig-admin:${VERSION}

clean:
	-rm mconfig-admin

.PHONY: clean