build:
	docker compose run -d db
run:
	docker compose build client && docker compose run client