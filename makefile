.PHONY: build-image run test
REGISTRY=
TAG=latest
APP=btcltp

build-image:
	docker build --rm -t ${REGISTRY}/${APP}:${TAG} .

push-image:
	docker push ${REGISTRY}/${APP}:${TAG}

run:
	go run ./main.go

test:
	go test -cover ./...

