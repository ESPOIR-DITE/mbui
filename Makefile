
docker-build :
	docker build --tag mbui:latest .
	#docker build --tag 0816048957docker/mbui:latest .

docker-run :
	docker create --name mbui p 4000:4000 mbui


