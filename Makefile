build:
	docker build --tag house_market:latest .

run:
	docker run -it --env PROMETHEUS_FQDN="http://host.docker.internal:9091" --network host house_market:latest
