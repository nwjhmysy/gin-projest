generate-api:
	docker-compose -f docker-compose.api.yml run --rm openapi-generator-cli

build-linux-amd64:
	docker build --platform=linux/amd64 -t yinsiyu/gin-project .

launch-app:
	docker-compose -f docker-compose.app.yml up -d

image-push:
	docker push yinsiyu/gin-project