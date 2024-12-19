# go-webgate

A small service to route clients to different services. 

### Configuration

Create a configuration file in YAML format. Example below will run a webserver listening on localhost:1251. 
Clients must send a GET request with a header `X-Webgate-Request` containing value of `secret` from config.
```
---
rest:
  hostname: localhost
  port: 1251
services:
  name: Service 1
  secret: my-very-secret-long-string
  endpoints:
    - key: "ws"
      endpoint: "ws://localhost:12312"
      disabled: false
    - key: "ws-fallback"
      endpoint: "ws://localhost:12313"
      disabled: false
    - key: "ws-old"
      endpoint: "ws://localhost:12311"
      disabled: true
```

### Running

```
go-webgate serve --config=/path/to/config.yaml
```