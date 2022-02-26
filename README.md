# REST API

This is a simple REST API server that serves each endpoint with a known message as a response.

# Usage

Changing the message by initializing the environment variable `DEFSTR`

``` export DEFSTR=<desire-messege> ```

Running server:
``` make ```

Sending http request:
```
curl -X POST \
  http://localhost:3000/ \
  -H 'content-type: text/html' \
  -d <body>
```

**It was designed to only receive HTTP requests with the content-type text/html**