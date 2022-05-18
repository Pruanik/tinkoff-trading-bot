RUN_ARGS := $(wordlist 2, $(words $(MAKECMDGOALS)), $(MAKECMDGOALS))
$(eval $(RUN_ARGS):;@:)

migration-create:
	./migrate create -ext sql -dir migrations $(RUN_ARGS)

migration-up:
	./migrate -path ./migrations/ -database postgres://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=disable up $(RUN_ARGS)

migration-down:
	./migrate -path ./migrations/ -database postgres://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=disable down $(RUN_ARGS)
