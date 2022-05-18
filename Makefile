RUN_ARGS := $(wordlist 2, $(words $(MAKECMDGOALS)), $(MAKECMDGOALS))
$(eval $(RUN_ARGS):;@:)

build-all:
	@sh -c "'$(CURDIR)/scripts/build.sh'"

build-webapp:
	@sh -c "'$(CURDIR)/scripts/build.sh' webapplication"

build-migrator:
	@sh -c "'$(CURDIR)/scripts/build.sh' migrator"

up:
	docker-compose up -d

migration-create:
	docker-compose run --rm --entrypoint "" migrator make migration-create $(RUN_ARGS)

migration-up:
	docker-compose run --rm --entrypoint "" migrator make migration-up

migration-down:
	docker-compose run --rm --entrypoint "" migrator make migration-down
