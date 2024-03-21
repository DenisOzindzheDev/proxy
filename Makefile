## run: starts demo http services
.PHONY: run-containers
run-containers:
#docker-run

## stop: stops all demo services
.PHONY: stop
stop:
 docker stop server1
 docker stop server2
 docker stop server3