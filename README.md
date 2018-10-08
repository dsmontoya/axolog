# axolog
Axolog follows the logs of all the containers on a host and ship them to a target. It automatically detects when a container is started or stopped.

**Note:** Axolog is not production ready. Use it carefully.

# Usage
Run the container image specifying the target URI (only tcp and udp supported). It's important to mount the host's `docker.sock` file.

``` sh
docker run --name axolog -d --volume=/var/run/docker.sock:/var/run/docker.sock dsmontoya/axolog:0.2.0 axolog udp://localhost:5000
```
