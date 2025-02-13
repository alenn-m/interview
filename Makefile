include .env
export

migrate:
	goose -dir ./svc/migrations postgres "${LOCAL_DB}" up

rollback:
	goose -dir ./svc/migrations postgres "${LOCAL_DB}" down

new-migration: ## New migration (make name=add-some-table new-migration)
	$(eval MIGRATION_FILE := $(shell date +"svc/migrations/%Y%m%d%H%M%S_$(name).sql"))
	touch $(MIGRATION_FILE)
	echo "-- +goose Up" >> $(MIGRATION_FILE)
	echo "\n-- +goose Down" >> $(MIGRATION_FILE)

setup-migrate: ## Install the migrate tool
	go get github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
