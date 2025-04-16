# Secure NATS Setup with JetStream and JWT Authentication

This guide describes how to set up a secure [NATS](https://nats.io) server with [JetStream](https://docs.nats.io/nats-concepts/jetstream) and full JWT-based authentication using Docker Compose, including operator/account/user configuration and client `.creds` usage.

---

## 📁 Directory Structure

```
nats-service/
├── docker-compose.yml
├── nats-server.conf
├── config/
│   ├── operator.jwt
│   └── jwt/
│       ├── <account-public-key>.jwt
│       └── <user-public-key>.jwt
├── jetstream-data/           # JetStream persistence
├── gateway-user.creds        # Copied to gateway-service (not in container)

gateway-service/
├── Dockerfile
├── .env
├── main.go
├── gateway-user.creds        # For local dev (not baked into image)
```

---

## 🔧 Prerequisites

- Docker & Docker Compose
- Homebrew (for CLI tools on macOS)

Install the necessary tools:

```bash
brew install nats-io/natscli/nsc
brew install nats-io/nats-tools/nats
```

---

## 🚀 Step-by-step: Create Operator, Account and User

### 1. Create workspace folder
```bash
mkdir -p ~/nats-auth && cd ~/nats-auth
```

### 2. Create operator
```bash
nsc add operator my-operator
```

### 3. Create account
```bash
nsc add account microservices
```

### 4. Create user for your gateway service
```bash
nsc add user -a microservices gateway-user
```

### 5. Generate `.creds` file
```bash
nsc generate creds -a microservices -n gateway-user -o ./gateway-user.creds
```

### 6. Export JWTs for the server

Set context:
```bash
nsc env -o my-operator
nsc env -a microservices
```

Then export:
```bash
nsc describe operator --raw > ./config/operator.jwt
nsc describe account --raw > ./config/jwt/<account-public-key>.jwt
nsc describe user --raw > ./config/jwt/<user-public-key>.jwt
```

You can find the public keys with:
```bash
nsc describe account
nsc describe user
```

---

## 🐳 Docker Compose Setup

### docker-compose.yml

```yaml
version: '3.9'

services:
  nats:
    image: nats:latest
    container_name: nats-secure
    ports:
      - "4222:4222"
      - "8222:8222"
    volumes:
      - ./nats-server.conf:/etc/nats/nats-server.conf
      - ./config/operator.jwt:/config/operator.jwt
      - ./config/jwt:/config/jwt
      - ./jetstream-data:/data/jetstream
    command: ["-js", "-c", "/etc/nats/nats-server.conf"]
```

### nats-server.conf

```hcl
operator: /config/operator.jwt

resolver: {
  type: full
  dir: "/config/jwt"
}

system_account: <account-public-key>

http: 8222

jetstream {
  store_dir: "/data/jetstream"
  max_mem_store: 1Gb
  max_file_store: 10Gb
}
```

Replace `<account-public-key>` with the real key from `nsc describe account`.

---

## ✅ Start the NATS container

```bash
docker-compose up -d
```

Check logs:

```bash
docker logs nats-secure
```

Look for: `Server is ready`, `JetStream enabled`, and no auth errors.

---

## 🧪 Test the credentials

```bash
nats --creds ./gateway-user.creds pub test.subject "Hello NATS"
```

Expected result:
```bash
Published 13 bytes to "test.subject"
```

---

## 🔐 Usage in Gateway Service (Go example)

```go
nats.Connect(
  os.Getenv("NATS_URL"),
  nats.UserCredentials(os.Getenv("NATS_CREDS")),
)
```

`.env` example:

```env
NATS_URL=nats://nats:4222
NATS_CREDS=/run/secrets/nats_user.creds
```

`docker-compose.yml` (for local dev):

```yaml
volumes:
  - ./secrets/nats_user.creds:/run/secrets/nats_user.creds
```

---

## ☸️ Kubernetes (Minikube or Production)

In production, create a Kubernetes Secret:

```bash
kubectl create secret generic gateway-creds   --from-file=gateway_user.creds=./gateway-user.creds
```

And mount it:

```yaml
volumeMounts:
  - name: creds
    mountPath: /etc/nats/creds
    readOnly: true
volumes:
  - name: creds
    secret:
      secretName: gateway-creds
```

Then your app reads:

```env
NATS_CREDS=/etc/nats/creds/gateway_user.creds
```

---

## ✅ Done!

You now have a secure, JWT-authenticated NATS server with JetStream and Go clients using credentials. 🎉

Let me know if you want the Kubernetes deployment YAMLs or automation scripts!