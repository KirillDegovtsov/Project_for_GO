local.run:
	go run ./cmd/app/main.go

docker.run:
	docker compose --env-file ./docker.env up -d
docker.run.db:
	docker compose --env-file ./docker.env up -d postgres