# ==============================================================================
# Docker compose commands

develop:
	echo "Starting docker environment"
	docker-compose --env-file ./config/.env up --build

prod:
	echo "Starting docker prod environment"
	docker-compose  --env-file ./config/.prod.env up --build

start-dev:
	echo "Running Development"
	docker-compose --env-file ./config/.env up

genmock:
	echo "Generating interface mock"
	go generate ./...

test:
	echo "Testing ..."
	go test -v ./...

