# wait-for-pg

Simple app that checks if PostgreSQL database is ready or not.

## Install and usage

### Option 1: binary

```sh
wget https://github.com/mxssl/wait-for-pg/releases/download/v1.0.3/wait-for-pg-linux-amd64.tar.gz
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
  mxssl/wait-for-pg:v1.0.3 \
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
