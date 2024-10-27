docker build . -t go-api
docker run -it -p 5050:5050 go-api

docker ps -a

