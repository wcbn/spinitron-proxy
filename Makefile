.PHONY: build start stop logs push

IMAGE_NAME = wcbn/spinitron-proxy
CONTAINER_NAME = spinitron-proxy-container

stop:
	(docker stop $(CONTAINER_NAME)) || true

build:
	docker build --platform=linux/amd64 --tag $(IMAGE_NAME) .

start: stop	build
	docker run --platform=linux/amd64 --env SPINITRON_API_KEY=$$SPINITRON_API_KEY  --rm -d -p 8080:8080 --name $(CONTAINER_NAME) $(IMAGE_NAME)

logs:
	docker logs -f $(CONTAINER_NAME)

push:
	docker push $(IMAGE_NAME)
