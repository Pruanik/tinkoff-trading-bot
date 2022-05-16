build-all:
	@sh -c "'$(CURDIR)/scripts/build.sh'"

build-webapp:
	@sh -c "'$(CURDIR)/scripts/build.sh' webapplication"

up:
	docker-compose up -d