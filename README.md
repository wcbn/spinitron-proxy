# Spinitron proxy server

We need a proxy in order to meet the terms of service for Spinitron's API (no credentials stored on client).

## Getting Started

1. Run Docker
2. Make changes to `main.go`
3. Run `SPINITRON_API_KEY=XXX make start`

This will stop the existing server (if any), rebuild a new one, and run the new one.

At this point, you should be able to make [some requests](https://spinitron.github.io/v2api/)

```
curl "localhost:8080/api/v1/spins"
```

4. If you are done, go to step 5. Otherwise, go to step 2.
5. Run `make stop` to turn it off.

## Makefile

Feel free to run any of the target in the `Makefile`.

Pro tip: `make logs` will watch the docker logs and can be super helpful!
