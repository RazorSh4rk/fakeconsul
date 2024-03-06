# API Call Collection Documentation

## Description

This image should really be used for a backend for your library instead of
having to mock everything, but just in case, here are the implemented endpoints.

To use it in a docker compose file:
```yaml
  consul:
    image: ghcr.io/razorsh4rk/fakeconsul:master
    ports:
      - "8500:8500"
    container_name: consul
    restart: always
```

If you want to use these endpoints and you also use Insomnia, you can
find the exported workspace near the code. This documentation was
generated from it with chatgpt.

## Requests

### Delete all keys under a main key

Description: Deletes all keys under a specified main key in the Consul key-value store.

Method: DELETE

URL: `localhost:8500/v1/kv/mainkey`

Body: None

---
### Delete a specific key

Description: Deletes a specific key in the Consul key-value store.

Method: DELETE

URL: `localhost:8500/v1/kv/mainkey/subkey`

Body: None

---
### Get multiple values

Description: Retrieves multiple values under a specified main key in the Consul key-value store.

Method: GET

URL: `localhost:8500/v1/kv/mainkey`

Parameters:
    `recurse: true`

---
### Get a single value

Description: Retrieves a single value for a specific key in the Consul key-value store.

Method: GET

URL: `localhost:8500/v1/kv/mainkey/subkey`

Body: None

---
### Add or update a value

Description: Adds or updates a value for a specific key in the Consul key-value store.

Method: PUT

URL: `localhost:8500/v1/kv/mainkey/subkey`

Headers:
Content-Type: text/plain

Body: value
