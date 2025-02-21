# dockstat

Simple docker container management cli tool

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

## Usage

Clone this repository and run:

```bash
make
sudo make install
```

Run the command `dockstat` from anywhere in the terminal.

```bash
Usage:
  dockstat          : list running containers
  dockstat kill <container_id> : kill a running container
  dockstat log <container_id>  : show logs of a container
```
