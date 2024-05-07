SRC_DIR=src
FRONTEND_DIR=$(SRC_DIR)/frontend/infrastructure/discord
DB_NAME=papers.db
MIGRATIONS_DIR=migrations

MIGRATION_FILES=$(shell find $(MIGRATIONS_DIR) -name '*.sql' | sort)

.PHONY: run migrate lint

run:
	cd $(FRONTEND_DIR) && deno run --allow-net --allow-read --allow-env --allow-write bot.ts

migrate:
	@$(foreach file,$(MIGRATION_FILES),sqlite3 $(DB_NAME) < $(file);)

lint:
	deno lint