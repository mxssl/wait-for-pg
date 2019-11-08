# wait-for-pg

## install

```sh
wget https://github.com/mxssl/wait-for-pg/releases/download/v0.0.1/wait-for-pg -O /usr/local/bin/wait-for-pg

chmod +x /usr/local/bin/wait-for-pg
```

## usage

```sh
./wait-for-pg check \
  --host postgres.domain.com \
  --port 5432 \
  --user pguser \
  --password pgpass \
  --dbname dbname \
  --retry 10 \
  --sleep 2
```

- If PG is ready then app returns exit code 0
- If PG isn't ready then app returns exit code 1
