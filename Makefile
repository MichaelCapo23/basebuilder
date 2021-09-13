migration-create:
	cd migration && go run . create $(name) $(type)

migrate:
	cd migration && go run . up

gentoken:
	cd cmd/gentoken && go run .