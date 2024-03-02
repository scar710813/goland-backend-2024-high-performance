TAG=$(shell git log -1 --format=%h)

build:
	docker build -t rinha-de-backend-2024-q1 ./docker/
login:
	docker login
tag: login
	docker tag rinha-de-backend-2024-q1 kleytonsolinho/rinha-de-backend-2024-q1:$(TAG)
push: tag
	docker push kleytonsolinho/rinha-de-backend-2024-q1:$(TAG)