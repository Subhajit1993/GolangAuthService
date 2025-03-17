## Container Deployment

```bash
docker build -t go_auth_service .

docker tag go_auth_service subhajit1993/go_auth_service:1.0.0

docker push subhajit1993/go_auth_service:1.0.0
```

Test the container locally

```bash
docker run -p 8080:8080 go_auth_service
```


<h3>TODOs</h3>

1. Refresh token
2. 