# EventManager - Service Integration
This README.md contains all essential steps for provisioning and testing the integration of the Server and Ticket and Event service from the EventManager project. There is also an additional guide for Docker Rootless as all associated containers are useable via rootless Docker.

## Requirements
* Ubuntu 24.04 Server (recommended)
* Docker Runtime (more security with Docker Rootless)
* Source code of Ticket and Event service

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

## Privsioning and testing this integration
### Starting the Container
First of all you must create a .env file in the Event_Tickets folder. You have an example for it there than you can start the container:

```bash
docker compose up --build
```

### Test
1. Go in the Event Service GUI should be http://container-host:8001
2. Klick on the Details button for the first Event (you can also create or update Events under the admin panel in the Event service)
3. Here you can see now the details for the Event and there is a button Buy tickets. Klick on it and it will redirect you to the Tickets service. Important you must change redirect IP address to your container-host address.
4. On the Ticket service you can add a ticket or delete a ticket. This has an impact for the available_tickets for the event.
5. If the add or delete function worked it should be a change for the available_tickets on details page.


---




**Author:** kingdanxi & pintere6\
**Year:** 2024-2025\
