up:
	docker-compose up -d

down:
	@docker ps -aq | xargs docker rm -f || true
	if [ -d "data" ]; then \
		sudo rm -r data; \
	fi

status:
	docker container ls

stoped:
	docker container ls -a

run: 
	cd server && make run

lint-server:
	cd server && make lint

lint-client:
	cd client && make lint