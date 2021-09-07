migration-create:
	cd migration && go run . create $(name) sql

migrate:
	cd migration && go run . up