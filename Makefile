GOOS?=linux
GOARCH?=amd64
APP_EXECUTABLE?=./bin/otusdemo
RELEASE?=0.2
IMAGENAME?=arahna/otusdemo:release-$(RELEASE)

.PHONY: clean
clean:
	rm -f ${APP_EXECUTABLE}

.PHONY: build
build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
	    -ldflags="-w -s" \
	    -o ${APP_EXECUTABLE} ./cmd

.PHONY: release
release:
	git tag v$(RELEASE)
	git push origin v$(RELEASE)

.PHONY: container
container: build
	docker build -t $(IMAGENAME) .

.PHONY: minikube
minikube-run: container
	kubectl apply -f ./k8s/postgres.yaml \
        -f ./k8s/secrets.yaml \
        -f ./k8s/config.yaml \
        -f ./k8s/initdb.yaml \
        -f ./k8s/deployment.yaml \
        -f ./k8s/service.yaml \
        -f ./k8s/ingress.yaml

.PHONY: minikube-clean
minikube-clean:
	kubectl delete -f ./k8s/