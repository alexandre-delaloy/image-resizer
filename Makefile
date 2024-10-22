help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup-env: ## Copy sample files 
	if [ -f api/.env ] ; then mv api/.env api/.env.old ; fi
	cp api/.env.sample api/.env

start: ## Up the docker-compose without cache or orphans
	docker-compose up \
	--detach \
	--build \
	--remove-orphans \
	--force-recreate

restart: ## down, then up the docker-compose
	make stop
	make start

stop: ## Down the docker-compose 
	docker-compose down

logs: ## Display logs of your containers 
	docker-compose logs --follow

lint-app:
	cd app ; \
	npm run lint ; \
	cd -
	
lint-api:
	cd api ; \
	gofmt -s -w -l . ; \
	cd -

start-producer:
	cd api/worker ; \
	go run producer/main.go ; \
	cd -

start-consumer:
	cd api/worker ; \
	go run consumer/main.go ; \
	cd -


.PHONY: help
