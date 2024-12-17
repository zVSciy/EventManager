# Guide - EventManager - Ticket Service
This README.md contains all essential steps for provisioning and testing the Ticket service from the EventManager project. There is also an additional guide for Docker Rootless, as the Ticket service can also be used in rootless container environments.

## Requirements
* Ubuntu 24.04 Server (recommended)
* Docker Runtime (more security with Docker Rootless)
* Source code of Ticket service

## Ubuntu Server
As no GUI is needed for running the Ticket service, only a Ubuntu in the server version is necessary. As every step of the installation is not covered in this README.md, please have a look at the official step-by-step guide: https://ubuntu.com/tutorials/install-ubuntu-server#1-overview

The following resources are recommended for the VM:
* **Processors:** 4
* **Memory (RAM):** 4 GB
* **Storage:** min. 20 GB

## Running Rootless Docker
To start off, if you have not installed the Docker Runtime yet, mind Docker's installation guide: https://docs.docker.com/engine/install/ubuntu/

The following commands for rootless Docker originate from Docker's official documentation, see: https://docs.docker.com/engine/security/rootless/#install. Keep in mind that the given commands should be executed with the user that is responsible for Docker later on.

1. Install `uidmap`
```bash
sudo apt-get install -y uidmap
``` 
2. Shutdown Docker's system deamon
```bash
sudo systemctl disable --now docker.service docker.socket
```

3. Install `Dockerd-Rootless-Setup`
```bash
dockerd-rootless-setuptool.sh install
```

4. Define the variable `DOCKER_HOST`
```bash
export DOCKER_HOST=unix://$XDG_RUNTIME_DIR/docker.sock
```

The Docker engine is now placed in the current user's directory which means that the Docker runtime uses user permissions instead of `root` ones.

## Provisioning the Ticket service
At this point, the Ticket service may be provisioned with the last steps provided.

1. Clone the source code from GitHub
```bash
git clone https://github.com/zVSciy/EventManager.git
```

2. Switch into the `Tickets` folder and run the containers with `docker-compose.yml`
```bash
cd ./EventManager/Tickets
docker compose up [parameters]
```
If you are not familiar with Docker, here are some parameters that might be helpful.
* **-d** &rarr; Detached mode; Output of containers is not shown in the command line which causes the CLI to be further accessible.
* **--build** &rarr; All containers, including their environments, are rebuild.

## Testing
> [!NOTE]
> As no integration with other services has been done so far, keep in mind that given endpoints might need some further adjustments when integrated with other services - it might be possible that endpoints are also added or completely changed, regarding how the integration is done (CRUD functions seem to be enough to start off). The same goes for the GUI - every functionality might not be practical yet.

Currently, the FastAPI application of the Ticket service is automatically tested with the `src_fastapi/test_app.py` file which includes some unit testing. This python script is executed when provisioning the API container - the output of the unit test is shown in the CLI if the parameter `-d` is NOT in use.

However, testing processes on the graphical user interface must be performed manually.

---

**Author:** kingdanxi\
**Year:** 2024-2025\