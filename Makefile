DB_CONTAINER_ID:=$(shell docker container ls --all --quiet --filter "name=database")

.PHONY: local-run
# Run the local docker-compose
local-run: 
	docker-compose up --build

.PHONY: local-backup
# Backup the local docker database
local-backup: 
	@echo "Backing up database container $(DB_CONTAINER_ID)"
	docker exec $(DB_CONTAINER_ID) /usr/bin/mysqldump -u root --password=root directus > cms/db/schema.sql
	@echo "Backup complete. File saved to cms/db/schema.sql"

.PHONY: local-stop
# Stop the local docker-compose
local-stop: 
	docker-compose down

.PHONY: load-test
# Loadtest the targets
load-test: 
	@echo "Load testing the website"
	k6 run loadtest.js

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help