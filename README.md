# Spinitron proxy server

We need a proxy in order to meet the terms of service for Spinitron's API:

> Two rules that we hope you will follow in your use of the Spinitron API can impact design of your app.
>
> First, as a client of Spinitron you have an API key that allows access to the API. You may not delegate your access to third parties except for development purposes. In particular, you may not delegate your access to clients of your web or mobile apps. For example, a web or mobile client of yours may fetch data from your servers but they may not use your API key to access the Spinitron API. In other words, don’t build Spinitron API clients into your client scripts or mobile apps.
>
> Second, you should use a cache in the implementation of your web servers or mobile app back-end servers. For example, say you display part of the program schedule on a certain web page. It’s not ok if for every request of that page from visitors to your site, your server makes a request to the Spinitron API. The program schedule doesn’t change rapidly so you can assume that data you fetched a minute ago is still valid. So you should cache data you fetch from Spinitron for a while in case you need the same info again soon. Cacheing is good for your website visitors (faster page loads), reduces load on your and Spinitron’s servers, reduces Internet traffic (and therefore even reduces energy waste a little). How you implement the cache is up to you. Good cache implementations take into account the specific design of the app and how users are expected to behave.

## Getting Started

1. Run Docker
2. Make changes to `main.go`
3. Run `SPINITRON_API_KEY=XXX make start`

This will stop the existing server (if any), rebuild a new one, and run the new one.

At this point, you should be able to make [some requests](https://spinitron.github.io/v2api/)

```
curl "localhost:8080/api/spins"
```

4. If you are done, go to step 5. Otherwise, go to step 2.
5. Run `make stop` to turn it off.

## Makefile

Feel free to run any of the target in the `Makefile`.

Pro tip: `make logs` will watch the docker logs and can be super helpful!
