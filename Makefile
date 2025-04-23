build:
	docker compose build --no-cache
up:
	docker compose up -d
down:
	docker compose down
restart:
	docker compose down
	docker compose up -d
rebuild:
	docker compose down
	docker compose build --no-cache
	docker compose up -d
start:
	docker compose build --no-cache
	docker compose up -d

install:
	cp .env.example .env && \
	cp ad/.env.example ad/.env && \
	cp tracker/.env.example tracker/.env && \
	start