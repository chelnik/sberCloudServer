all: restart


build:
	docker compose up -d

restart:
	docker compose down
	docker rmi sbercloudserver-api
	docker compose up -d
