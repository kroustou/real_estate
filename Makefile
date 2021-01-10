build:
	docker build --target mac --tag raspberrypi.local:5000/real_estate:mac .
#	docker build --target arm --tag raspberrypi.local:5000/real_estate:arm .

run:
	docker run -it --env PROMETHEUS_FQDN="http://host.docker.internal:9091" --network host raspberrypi.local:5000/real_estate:mac /go/bin/real_estate
