docker-dev:
	@docker compose -f docker-compose.dev.yaml up --build

docker-prod:
	@docker compose -f docker-compose.prod.yaml up --build
