goose -dir "./migrations" postgres "user=username password=password host=localhost port=5432 dbname=best_chat sslmode=disable" up 

#### Run project
```shell
cp .env.sample .env
```

```shell
docker compose up
```
