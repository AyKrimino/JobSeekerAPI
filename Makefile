build:
	@go build -o bin/JobSeekerAPI cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/JobSeekerAPI

migrations:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

migrate-force:
	@if [ -z "$(version)" ]; then echo "Error: Please provide a version using 'make migrate-force version=<migration_version>'"; exit 1; fi
	@migrate -path cmd/migrate/migrations -database "mysql://admin:admin@tcp(localhost:3306)/JobSeeker" force $(version)