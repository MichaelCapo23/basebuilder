# basebuilder
This is a golang repo to be used as a "base-builder" for future projects. This repo is fully dockerized and made to provide auth via firebase, along with postgres migrations using goose. 

## Start Using
1. spin up postgres container
```
docker compose up -d postgres
```

2. run the migrations
```
make migrate
```
