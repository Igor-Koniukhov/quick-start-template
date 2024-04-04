

#### Run project
```shell
cp .env.sample .env
```

```shell
docker compose up
```
### Catalog structure for adding ci-cd-pipeline.yml on github
```
.
├── .github
│   └── workflows
│       └── ci-cd-pipeline.yml
├── backend
│   ├── Dockerfile
│   └── ...
├── frontend
│   ├── Dockerfile
│   └── ...
└── ...

```

```shell
mkdir -p .github/workflows
touch .github/workflows/ci-cd-pipeline.yml

```