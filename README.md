# wait-for-pg

Simple app that checks if PostgreSQL database is ready or not.

## Install and usage

### Option 1: binary

```sh
wget https://github.com/mxssl/wait-for-pg/releases/download/v1.0.5/wait-for-pg-linux-amd64.tar.gz
tar xvzf wait-for-pg-linux-amd64.tar.gz
mv wait-for-pg /usr/local/bin/wait-for-pg
chmod +x /usr/local/bin/wait-for-pg
rm wait-for-pg-linux-amd64.tar.gz

wait-for-pg check \
  --host postgres.domain.com \
  --port 5432 \
  --user pguser \
  --password pgpass \
  --dbname dbname \
  --sslmode disable \
  --retry 10 \
  --sleep 2
```

### Option 2: docker container

```sh
docker container \
  run \
  --rm \
  mxssl/wait-for-pg:v1.0.5 \
  wait-for-pg check \
    --host postgres.domain.com \
    --port 5432 \
    --user pguser \
    --password pgpass \
    --dbname dbname \
    --sslmode disable \
    --retry 10 \
    --sleep 2
```

- If PG is ready then app returns exit code 0
- If PG isn't ready then app returns exit code 1

### Option 3: Kubernetes init container

Use `wait-for-pg` as an init container to delay your application startup until PostgreSQL is ready.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
spec:
  replicas: 1
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      initContainers:
        - name: wait-for-pg
          image: mxssl/wait-for-pg:v1.0.5
          command:
            - wait-for-pg
            - check
            - --host
            - postgres-service
            - --port
            - "5432"
            - --user
            - pguser
            - --password
            - pgpass
            - --dbname
            - dbname
            - --sslmode
            - disable
            - --retry
            - "30"
            - --sleep
            - "2"
      containers:
        - name: myapp
          image: myapp:latest
```

Using environment variables from a Secret:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: pg-credentials
type: Opaque
stringData:
  PGHOST: postgres-service
  PGPORT: "5432"
  PGUSER: pguser
  PGPASSWORD: pgpass
  PGDATABASE: dbname
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
spec:
  replicas: 1
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      initContainers:
        - name: wait-for-pg
          image: mxssl/wait-for-pg:v1.0.5
          command: ["sh", "-c"]
          args:
            - >
              wait-for-pg check
              --host $(PGHOST)
              --port $(PGPORT)
              --user $(PGUSER)
              --password $(PGPASSWORD)
              --dbname $(PGDATABASE)
              --sslmode disable
              --retry 30
              --sleep 2
          envFrom:
            - secretRef:
                name: pg-credentials
      containers:
        - name: myapp
          image: myapp:latest
          envFrom:
            - secretRef:
                name: pg-credentials
```
