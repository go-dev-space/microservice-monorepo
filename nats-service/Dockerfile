FROM nats:latest

# Copy NATS configuration and JWTs into the image
COPY nats-server.conf /etc/nats/nats-server.conf
COPY config/operator.jwt /config/operator.jwt
COPY config/jwt /config/jwt

# Start NATS with JetStream enabled and the config
CMD ["-js", "-c", "/etc/nats/nats-server.conf"]