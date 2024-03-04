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

## How to deploy

The only prerequisite here is having an API key. Container services are supported by most cloud providers these days. The memory and CPU requirements are extremely minimal.

### AWS Lightsail

1. Create a new container service (the cheapest option is fine)
1. Create a new deployment
1. Set the image to `docker.io/wcbn/spinitron-proxy:latest`
1. Set an environment variable named `SPINITRON_API_KEY` with the value of your API key
1. Set the port to 8080, HTTP
1. Use the container name as the public endpoint
1. Set the health check to `/`
1. Save and deploy

## How to Develop

### Requirements

- Go 1.19

1. Make changes to `main.go`
1. Run `SPINITRON_API_KEY=XXX go run main.go`
1. Make [some requests](https://spinitron.github.io/v2api/) e.g. `curl "localhost:8080/api/spins"`

### Release

1. `docker login`
1. `make build`
1. `make push`
1. See new version here: https://hub.docker.com/repository/docker/wcbn/spinitron-proxy

# Cache strategy

## Individual resources

- Key is a `(collection-name, id)` pair formatted as a URL path string e.g. `"/spins/1"`
- Value is a `byte[]` JSON document
- Query parameters are ignored
- TTL of 3 minutes

## Collections

- Key is the substring of the request URL composed of the path and the query string e.g. `"/spins?page=1"`
- Value is a `byte[]` JSON document
- Maximum TTL depends on the collection. See `cache.go` for details.
- As soon as one cache for a collection expires, all caches for that collection are invalidated e.g. When `/spins?page=1` expires, `/spins?page=3` is also invalidated (and vice-versa).
