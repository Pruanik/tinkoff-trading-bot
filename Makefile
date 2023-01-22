RUN_ARGS := $(wordlist 2, $(words $(MAKECMDGOALS)), $(MAKECMDGOALS))
$(eval $(RUN_ARGS):;@:)

### Building Docker Containers ###
build-all:
	@sh -c "'$(CURDIR)/scripts/build-containers.sh'"

build-webapp:
	@sh -c "'$(CURDIR)/scripts/build-containers.sh' webapplication"

build-tinkoffconnection:
	@sh -c "'$(CURDIR)/scripts/build-containers.sh' tinkoffinvestconnection"

build-migrator:
	@sh -c "'$(CURDIR)/scripts/build-containers.sh' migrator"

build-frontend:
	@sh -c "'$(CURDIR)/scripts/build-containers.sh' frontend"

### Stop Docker Containers ###
stop:
	docker-compose stop

### Up Docker Containers ###
up-postgresql:
	docker-compose up -d postgresql

up-all:
	docker-compose up -d webapplication tinkoffinvestconnection postgresql

### Migration Commands ###

migration-create:
	docker-compose run --rm --entrypoint "" migrator make migration-create $(RUN_ARGS)

migration-up:
	docker-compose run --rm --entrypoint "" migrator make migration-up

migration-down:
	docker-compose run --rm --entrypoint "" migrator make migration-down

### Frontend Commands ###
frontend-install:
	docker-compose run --rm frontend yarn install

frontend-build:
	docker-compose run --rm frontend yarn build

frontend-watch:
	docker-compose run --rm frontend yarn watch

### Up Application Local ###
run-webapp:
	go run cmd/webapplication/main.go

run-tinkoffconnection:
	go run cmd/tinkoffinvestconnection/main.go
