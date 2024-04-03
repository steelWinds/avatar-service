build:
	docker build -t identavatar . && docker run --name identavatar -d -p 3180:3180 identavatar

up:
	docker compose -f docker/docker-compose.dev.yml up -d

test:
	go test -v ./...
