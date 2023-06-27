# Druktezoeker

Dummy application to query the (acceptance) crowdedness API.

## Help

### Druktezoeker

```shell
NAME:
   druktezoeker - Bevraag de Crowdedness API.

USAGE:
   druktezoeker [global options] command [command options] [arguments...]

COMMANDS:
   bikes
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

### Fietsen

```shell
NAME:
   druktezoeker bikes

USAGE:
   druktezoeker bikes [command options] [arguments...]

DESCRIPTION:
   Zoek totaal aantal fietsplaatsen voor een trein

OPTIONS:
   --trains value [ --trains value ]  train numbers
   --api_key value                     [$APIM_SUBSCRIPTION_KEY]
   --host value                        [$HOST]
   --help, -h                         show help

```

## Example

```shell
go run main.go bikes -trains=3612 -trains=36190
```