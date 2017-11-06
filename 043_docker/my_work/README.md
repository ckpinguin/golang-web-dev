# Go & Network stuff

## Run container (no port configured, go standard is 8080)

``` shell
docker run  -dit --rm --name mygoapp mygo
```

## Find out IP address of the container

``` shell
docker inspect container_name | grep IPAddress
```


## Use a port mapping (preferred way)

``` shell
docker run -d -p 8080:80 my-app
```

## Expose the container's port to a locahost's port (old way?)

``` bash
 iptables -t nat -A  DOCKER -p tcp --dport 8001 -j DNAT --to-destination 172.17.0.19:8000
```

## Cleanup containers

``` bash
docker ps -a |grep -v mygo |awk '{ print $1 }' | xargs sudo docker
docker rmi $(docker images -f dangling=true -q
```