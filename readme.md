# API Gateway

A lightweight, configurable API Gateway written in Go.

## Requirements

- Go 1.24+
- Docker (optional)

## Installation

Clone the repository:

```bash
git clone https://github.com/abdul-ghaffar01/api-gateway.git
cd api-gateway
```

Install dependencies:

```bash
go mod tidy
```

---

## Configuration

Create a configuration file named `config.yaml`.

Example:

```yaml
server:
  port: 8080
  auth_endpoint: https://example.com/auth

routes:
  - path: /users
    upstream: http://user-service:8080

  - path: /orders
    upstream: http://order-service:8080
```

---

## Running

### Development

```bash
go run . --config=config.yaml
```

---

### Build

```bash
go build -o api-gateway
```

Run:

```bash
./api-gateway --config=config.yaml
```

---

### Docker

Build the image:

```bash
docker build -t api-gateway .
```

Run the container:

```bash
docker run \
    -p 8080:8080 \
    -v $(pwd)/config.yaml:/config/config.yaml \
    api-gateway \
    --config=/config/config.yaml
```

---

## Configuration Priority

The gateway loads configuration in the following order:

1. Command-line flag (`--config`)
2. Environment variable (`CONFIG_PATH`)
3. Default (`config.yaml`)

Example:

```bash
./api-gateway --config=prod.yaml
```

or

```bash
CONFIG_PATH=prod.yaml ./api-gateway
```