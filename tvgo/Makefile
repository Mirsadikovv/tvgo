DOCKER_DEV = docker compose -f docker-compose.dev.yml --env-file .env.dev
DOCKER_PROD = docker compose -f docker-compose.yml

go_run:
	env $$(grep -v '^#' .env.dev | xargs) go run main.go

swag_init:
	swag init -pd --ot json -g main.go

swag_gen:
	swag init -g main.go -o docs

docker_up:
	$(DOCKER_DEV) up --build --force-recreate -d

docker_down:
	$(DOCKER_DEV) down -v

docker_logs:
	$(DOCKER_DEV) logs -f

docker_config:
	$(DOCKER_DEV) config

docker_prod_up:
	$(DOCKER_PROD) up --build --force-recreate -d

docker_prod_down:
	$(DOCKER_PROD) down -v

docker_prod_logs:
	$(DOCKER_PROD) logs -f

docker_prod_config:
	$(DOCKER_PROD) config
