# Astro Ranking Manager (ARM)

### This ranking is still in early development and doesnt have any kind of security

This is only for study and early tests in Astro development.
Please dont use it in production.

- NOTE: Working with ["JWT"](https://jwt.io/) system

### 🚧 Requirements to start:

- Need Docker installed in your computer.
  - If you dont have Docker installed, you can get it here [Docker](https://www.docker.com/products/docker-desktop/)

### 🚧 Start up this manager:

To start this manager, run the `make build` command. 
This command will compile the API and start a database container with all the dependencies necessary for the successful execution of the manager.

For dev proposes, use `make dev` command. It will only build the database container and export its port.

For more commands run `make help`