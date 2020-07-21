APP_EXECUTABLE?=./bin/otusdemo
RELEASE?=0.4
MIGRATIONS_IMAGENAME?=arahna/otusdemo-migrations:release-$(RELEASE)
IMAGENAME?=arahna/otusdemo:release-$(RELEASE)
LOCUST_IMAGENAME?=arahna/otusdemo-locust:latest

.PHONY: clean
clean:
	rm -f ${APP_EXECUTABLE}

.PHONY: build
build: clean
	docker build -t $(MIGRATIONS_IMAGENAME) -f DockerfileMigrations .
	docker build -t $(IMAGENAME) .

.PHONY: release
release:
	git tag v$(RELEASE)
	git push origin v$(RELEASE)

.PHONY: helm-update-dependencies
helm-update-dependencies:
	helm dependency update ./helm

.PHONY: run
run: build
	helm install otusdemo ./helm

.PHONY: remove
remove:
	helm uninstall otusdemo

.PHONY: run-stresstest
run-stresstest:
	docker build -t ${LOCUST_IMAGENAME} ./locust
	kubectl apply -f stresstest.yaml

.PHONY: stop-stresstest
stop-stresstest:
	kubectl delete -f stresstest.yaml

.PHONY: minikube-run
k8s-run: build
	kubectl apply -f ./k8s/postgres.yaml \
        -f ./k8s/secrets.yaml \
        -f ./k8s/config.yaml \
        -f ./k8s/dbmigrations.yaml \
        -f ./k8s/deployment.yaml \
        -f ./k8s/service.yaml \
        -f ./k8s/ingress.yaml

.PHONY: minikube-remove
k8s-remove:
	kubectl delete -f ./k8s/