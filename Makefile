include $(GOPATH)/src/github.com/inoc603/go-make/Makefile

IMAGE := docker.elenet.me/ruijing.zhang/pansible

setup: schema key

proto:
	protoc -I stream --go_out=plugins=grpc:stream stream/stream.proto

runner_docker: build
	docker build -f Dockerfile-runner -t pansible-runner . 

t: runner_docker
	docker run -t -e PANSIBLE_TOKEN=`http post 127.0.0.1:5757/login | jq .token | tr -d '"'` pansible-runner
	# docker run -t -e PANSIBLE_TOKEN=`http post 127.0.0.1:5757/login | jq .token | tr -d '"'` pansible-runner ./pansible
	# docker run -it -e PANSIBLE_TOKEN=`http post 127.0.0.1:5757/login | jq .token | tr -d '"'` pansible-runner sh

schema:
	mysql -h 127.0.0.1 -uroot -ptoor < drop.sql
	mysql -h 127.0.0.1 -uroot -ptoor -Dpansible < create.sql

models:
	sqlboiler mysql -p models -o models --tinyint-as-bool

key:
	mkdir -p tmp
	ssh-keygen -t rsa -b 4096 -C "pansible@ele.me" -f tmp/id -N ""	
