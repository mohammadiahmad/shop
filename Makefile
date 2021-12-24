compile:


build:


up:
	docker-compose -f deployments/docker-compose.yml up -d

down:
	docker-compose -f deployments/docker-compose.yml down