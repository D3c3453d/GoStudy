db:
	docker compose up -d db
server:
	docker compose build server && docker compose up server
down:
	docker compose down