# Debug only

In response to https://stackoverflow.com/questions/66930097/race-with-mutex-corrupt-data-in-map and https://github.com/darmiel/yaxc/issues/8 .

See `internal/server/cache_test.go` in combination with `handlePostAnywhereWithHash` in `internal/server/server.go` and https://docs.gofiber.io/#zero-allocation

> Because fiber is optimized for high-performance, values returned from fiber.Ctx are not immutable by default and will be re-used across requests. As a rule of thumb, you must only use context values within the handler, and you must not keep any references. As soon as you return from the handler, any values you have obtained from the context will be re-used in future requests and will change below your feet. Here is an example:

# YAxC
Yet Another Cross Clipboard
> Allan, please add details!

## Demo
https://youtu.be/OVpH70byKRQ

## Server

### Set Data
Just make a POST request to any path:

**POST** `/hi`
```
Hello World!
```

#### TTL
By default, the data is kept for 5 minutes. This TTL can be changed via the `ttl`-parameter.

**POST** `/hi?ttl=1m30s`
```
Hello World!
```

#### Encryption
By default, the data is not encrypted. 
**It is not recommended to encrypt the data on server side. The data should always be encrypted on the client side.**

However, if this is not possible, the `secret`-parameter can be used to specify a password with which the data should be encrypted.

**POST** `/hi?secret=s3cr3tp455w0rd`
```
Hello World!
```
**Produces:**
```
gwttKS3Q2l0+YR+jQF/02u3fNVmMIcVOTNSGD5vWfrYTtH8adt8r
```

### Get Data
**GET** `/hi`
```
Hello World!
```

#### Encryption
If the data has been encrypted and should be decrypted on the server side (**which is not recommended**), the "password" can be passed via the `secret`-parameter.
**GET** `/hi`
```
gwttKS3Q2l0+YR+jQF/02u3fNVmMIcVOTNSGD5vWfrYTtH8adt8r
```

**GET** `/hi?secret=s3cr3tp455w0rd`
```
Hello World!
```


### CLI
```bash
Run the YAxC server

Usage:
  yaxc serve [flags]

Flags:
  -b, --bind string                 Bind-Address (default ":1332")
  -t, --default-ttl duration        Default TTL (default 1m0s)
  -e, --enable-encryption           Enable Encryption (default true)
  -h, --help                        help for serve
  -x, --max-body-length int         Max Body Length (default 1024)
  -s, --max-ttl duration            Max TTL (default 5m0s)
  -l, --min-ttl duration            Min TTL (default 5s)
      --proxy-header string         Proxy Header
  -r, --redis-addr string           Redis Address
      --redis-db int                Redis Database
      --redis-pass string           Redis Password
      --redis-prefix-hash string    Redis Prefix (Hash) (default "yaxc::hash::")
      --redis-prefix-value string   Redis Prefix (Value) (default "yaxc::val::")

Global Flags:
      --config string   config file (default is $HOME/.yaxc.yaml)
      --server string   URL of API-Server

```

## Client
### Watch
```bash
Watch Clipboard

Usage:
  yaxc watch [flags]

Flags:
  -a, --anywhere string     Path (Anywhere)
  -h, --help                help for watch
      --ignore-client       Ignore Client Updates
      --ignore-server       Ignore Server Updates
  -s, --passphrase string   Encryption Key

Global Flags:
      --config string   config file (default is $HOME/.yaxc.yaml)
      --server string   URL of API-Server (default "https://yaxc.d2a.io")
```
