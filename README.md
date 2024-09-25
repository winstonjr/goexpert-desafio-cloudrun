# Desafio Cloud Run - Sem VS Code

### Para rodar local
1. Baixar repositório
2. Criar API Key para consultar temperatura no site: `https://www.weatherapi.com/`
3. Criar um arquivo .env na mesma pasta que o `docker-compose.yaml`
4. Editar o arquivo .env e colocar como chave `WEATHER_API_KEY=` e como valor a chave criada no passo 2
5. Rodar `docker compose up` na pasta com o arquivo `docker-compose.yaml`

### Gerando a imagem do docker para o google cloud run

```shell
docker build -t gcr.io/YOUR_PROJECT_ID/desafio-cloudrun:latest .
```
```shell
gcloud authentication login
```
```shell
gcloud builds submit --tag gcr.io/YOUR_PROJECT_ID/desafio-cloudrun
```
```shell
gcloud run deploy desafio-cloudrun \
    --image gcr.io/YOUR_PROJECT_ID/desafio-cloudrun \
    --platform managed \
    --region us-central1 \
    --set-env-vars WEATHER_API_KEY="YOUR_WEATHER_API_KEY_HERE" \
    --allow-unauthenticated
```
