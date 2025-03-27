build:
	KO_DOCKER_REPO=ghcr.io/konstfish/hits ko build --bare --platform=all

deploy:
	KO_DOCKER_REPO=ghcr.io/konstfish/hits ko apply -f deployment/

local:
	KO_DOCKER_REPO=ghcr.io/konstfish/hits ko build --bare --platform=all
	docker-compose pull hits
	docker-compose up -d