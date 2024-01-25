build-linux-amd64:
	docker build --platform=linux/amd64 -t yinsiyu/gin-project .

launch-app:
	docker-compose up -d

image-push:
	docker push yinsiyu/gin-project