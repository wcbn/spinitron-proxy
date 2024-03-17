# Spinitron proxy server

Developers using the Spinitron API must adhere to the following terms of service:

> Two rules that we hope you will follow in your use of the Spinitron API can impact design of your app.
>
> First, as a client of Spinitron you have an API key that allows access to the API. You may not delegate your access to third parties except for development purposes. In particular, you may not delegate your access to clients of your web or mobile apps. For example, a web or mobile client of yours may fetch data from your servers but they may not use your API key to access the Spinitron API. In other words, don’t build Spinitron API clients into your client scripts or mobile apps.
>
> Second, you should use a cache in the implementation of your web servers or mobile app back-end servers. For example, say you display part of the program schedule on a certain web page. It’s not ok if for every request of that page from visitors to your site, your server makes a request to the Spinitron API. The program schedule doesn’t change rapidly so you can assume that data you fetched a minute ago is still valid. So you should cache data you fetch from Spinitron for a while in case you need the same info again soon. Cacheing is good for your website visitors (faster page loads), reduces load on your and Spinitron’s servers, reduces Internet traffic (and therefore even reduces energy waste a little). How you implement the cache is up to you. Good cache implementations take into account the specific design of the app and how users are expected to behave.

With that in mind, this little server...

- forwards requests from a client (e.g. a mobile app) to Spinitron with the developer's API key
- is read-only i.e. it only services GET requests
- includes an in-memory cache mechanism optimized for https://github.com/dctalbot/spinitron-mobile-app

## Cache strategy

### Individual resources

- When selecting an endpoint with an ID value e.g. `/spins/1`
- Query parameters are ignored
- TTL of 3 minutes

### Collections

- When selecting an endpoint that returns a list e.g. `/spins?`, `/spins?page=1`
- Query parameters are not ignored
- TTL depends on the collection. See comments in `cache.go` for details
- Upon expiration, all caches for the same collection are invalidated e.g. When `/spins?page=1` expires, `/spins?page=3` is also invalidated (and vice-versa).

## How to deploy

This software is distributed as container images which are hosted [on GitHub](https://github.com/wcbn/spinitron-proxy/pkgs/container/spinitron-proxy/versions?filters%5Bversion_type%5D=tagged).

The following architectures are supported: `linux/amd64`, `linux/arm/v7`, `linux/arm64`, `linux/ppc64le`, and `linux/s390x`.

Container-based services are supported by most cloud providers. The memory and CPU requirements are extremely minimal, so just pick the cheapest option.

### AWS Lightsail

1. Create a new container service
1. Create a new deployment
1. Set the image to `ghcr.io/wcbn/spinitron-proxy:latest`
1. Set an environment variable named `SPINITRON_API_KEY` with the value of your API key
1. Set the port to 8080, HTTP
1. Use the container name as the public endpoint
1. Set the health check to `/`
1. Save and deploy

## Related Projects

- https://github.com/dctalbot/react-spinitron
- https://github.com/dctalbot/spinitron-mobile-app

## How to Develop

### Requirements

- Go (version specified in `go.mod`)
- Spinitron API key

1. Make changes to `main.go`
1. Run `SPINITRON_API_KEY=XXX go run main.go`
1. Make [some requests](https://spinitron.github.io/v2api/) e.g. `curl "localhost:8080/api/spins"`

### Release

1. `docker login`
1. `make build`
1. `SPINITRON_API_KEY=XXX make start`, do some smoke testing e.g. `curl "localhost:8080/api/spins"`
1. `make push`
1. See new version here: https://hub.docker.com/repository/docker/wcbn/spinitron-proxy
