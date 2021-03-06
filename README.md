# KNoT Babeltower

## Installation and usage

### Requirements

*   Install glide dependency manager (<https://github.com/Masterminds/glide>)
*   Be sure the local packages binaries path is in the system's `PATH` environment variable:

```bash
$ export PATH=$PATH:<your_go_workspace>/bin
```

### Configuration

You can set the `ENV` environment variable to `development` and update the `internal/config/development.yaml` when necessary. On the other way, you can use environment variables to configure your installation. In case you are running the published Docker image, you'll need to stick with the environment variables.

The configuration parameters are the following (the environment variable name is in parenthesis):

*   `server`
    *   `port` (`SERVER_PORT`) **Number** Server port number. (Default: 80)

### Setup

```bash
make tools
make deps
```

### Compiling and running

```bash
make run
```

## Local (Development)

### Build and run (Docker)

A container is specified at `docker/Dockerfile`. To use it, execute the following steps:

01. Build the image:

    ```bash
    docker build . -f docker/Dockerfile -t cesarbr/knot-babeltower
    ```

01. Create a file containing the configuration as environment variables.

01. Run the container:

    ```bash
    docker run --env-file knot-babeltower.env -ti cesarbr/knot-babeltower
    ```

### Verify

```bash
curl http://<hostname>:<port>/healthcheck
```
