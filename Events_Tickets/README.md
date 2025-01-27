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
First of all, you must create a `.env` file in the Event_Tickets folder. A corresponding `env_exmaple` can be found in the mentioned folder. After providing the `.env` file, start the container stack with the following command:

```bash
docker compose up --build
```

### Test
1. Navigate to the GUI of the `Event service`, the URL should be: http://container-host:8001
2. Click on the `Details` button for the first event (you may also create or update events under the admin panel in the `Event service`)
3. On the `Details` page, you can see additional information for the selected event. There is also a `Buy tickets` which redirects to the `Tickets service`. Keep in mind that the redirect IP address must be changed to your container-host address in `/Events/event_svelte/src/routes/details/<id>/+page.svelte`.
4. On the Ticket service you can add or delete a ticket. This has an impact on the `available_tickets` for the event.
5. If the add or delete function worked, there should be a change of the `available_tickets` shown on the event details page.

---

**Author:** kingdanxi & pintere6\
**Year:** 2024-2025\
