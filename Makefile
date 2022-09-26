.PHONY: build start stop logs

IMAGE_NAME = spinitron-proxy-image
CONTAINER_NAME = spinitron-proxy-container

stop:
	(docker stop $(CONTAINER_NAME)) || true

build:
	docker build --tag $(IMAGE_NAME) .

start: stop	build
	docker run --env SPINITRON_API_KEY=$$SPINITRON_API_KEY  --rm -d -p 8080:8080 --name $(CONTAINER_NAME) $(IMAGE_NAME)

logs:
	docker logs -f $(CONTAINER_NAME)
