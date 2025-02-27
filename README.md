# kids-api

## Pré-req
brew install docker-buildx
## Criar link simbólico
mkdir -p ~/.docker/cli-plugins
ln -sfn $(brew --prefix)/opt/docker-buildx/bin/docker-buildx ~/.docker/cli-plugins/docker-buildx

## Verificar
docker buildx version

## Inicializar o projeto

go mod init github.com/eduardohass/kids-api
go get github.com/auth0/go-jwt-middleware@v2.1.0+incompatible
go get github.com/form3tech-oss/jwt-go@v3.2.5+incompatible
go get github.com/gorilla/mux@v1.8.1
go get github.com/jmoiron/sqlx@v1.3.5
go get github.com/lib/pq@v1.10.9
go mod tidy




# Construir e executar com Docker
make build
docker run -p 8080:8080 \
  -e DATABASE_URL="postgres://user:pass@host:5432/db" \
  -e AUTH0_DOMAIN="your-domain.auth0.com" \
  -e AUTH0_AUDIENCE="your-audience" \
  kids-api:latest

# Executar testes na API
curl -X POST http://localhost:8080/api/v1/children -H "Authorization: Bearer YOUR_TOKEN" -d @child.json

# Construir e publicar imagem
make build push

# Implantar em staging
make deploy-staging

# Executar migrações
make migrate

# Implantar em produção
make deploy-prod

# Fazer deployment com backup
make deploy-staging-with-backup

# Listar backups
kubectl exec -n kids-app-staging deploy/postgres -- ls /backup

# Recuperar um backup específico
kubectl cp kids-app-staging/$(kubectl get pods -n kids-app-staging -l app=postgres -o jsonpath='{.items[0].metadata.name}'):/backup/2024-01-01-1200.dump ./restore.dump