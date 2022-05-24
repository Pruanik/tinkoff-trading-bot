RUN_ARGS := $(wordlist 2, $(words $(MAKECMDGOALS)), $(MAKECMDGOALS))
$(eval $(RUN_ARGS):;@:)

migration-create:
	./wait-for $(DATABASE_HOST):$(DATABASE_PORT) -t 60 -- ./migrate create -ext sql -dir migrations $(RUN_ARGS)

migration-up:
	./wait-for $(DATABASE_HOST):$(DATABASE_PORT) -t 60 -- ./migrate -path ./migrations/ -database $(DATABASE_HOST)://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=disable up $(RUN_ARGS)

migration-down:
	./wait-for $(DATABASE_HOST):$(DATABASE_PORT) -t 60 -- ./migrate -path ./migrations/ -database $(DATABASE_HOST)://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=disable down $(RUN_ARGS)
