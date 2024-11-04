go_run:
	env $$(grep -v '^#' .env.dev | xargs) go run main.go

swag_init:
	swag init -pd --ot json -g main.go

DOCKER_VAR = docker compose -f docker-compose.dev.yml --env-file .env.dev

docker_up:
	$(DOCKER_VAR) up --build --force-recreate -d

docker_down:
	$(DOCKER_VAR) down -v

docker_config:
	$(DOCKER_VAR) config