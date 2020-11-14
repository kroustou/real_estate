build:
	docker build --tag real_estate:latest .

run:
	docker run -it --env PROMETHEUS_FQDN="http://host.docker.internal:9091" --network host real_estate:latest
