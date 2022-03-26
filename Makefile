# ==============================================================================
# Docker compose commands

develop:
	echo "Starting docker environment"
	docker-compose -f docker-compose.dev.yml up --build

docker_delve:
	echo "Starting docker debug environment"
	docker-compose -f docker-compose.delve.yml up --build

prod:
	echo "Starting docker prod environment"
	docker-compose -f docker-compose.prod.yml up --build

local:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml up --build