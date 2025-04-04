# dockstat

Simple docker container management cli tool

## Example output

```bash
|         ID        |      Name       |       Status        |
| :---------------- | :-------------: | ------------------: |
| 2e8d111569        |   container-1   | Up 5 days (healthy) |
| 2e8d111567        |   container-2   | Up 5 days (healthy) |
| 2e8d111562        |   container-3   | Up 5 days (healthy) |
| 2e8d111561        |   container-4   | Up 5 days (healthy) |
| 2e8d111560        |   container-5   | Up 5 days (healthy) |
| 2e8d111565        |   container-6   | Up 5 days (healthy) |
| 2e8d111563        |   container-7   | Up 5 days (healthy) |
| 2e8d111568        |   container-8   | Up 5 days (healthy) |
```

## Usage

Clone this repository and run:

```bash
# build
$ make

# install
$ sudo make install
```

Run the command `dck` from anywhere in the terminal.

```bash
Usage:
  dck show     : list running containers
  dck kill <container_id> : kill a running container
  dck log <container_id>  : show logs of a container
  dck run <container_id>  : start a stopped container
```

## Development

Run the test suite:

```bash
# using go
$ go test -v ./...

# using make
$ make test
```
