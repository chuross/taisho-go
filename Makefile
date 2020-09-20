.PHONY: local

C=""

local:
	docker-compose -f deployment/local/docker-compose.yml ${C}