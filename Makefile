docker-build:
	docker build -t ha-dns .
docker-run:
	docker run -it -p 8000:7001 ha-dns