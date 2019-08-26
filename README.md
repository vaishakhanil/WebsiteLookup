# Website Lookup CLI in Go
A simple Website Lookup CLI project created in Go to understand the concept of extracting command line arguments and using it for performing a simple network lookup#

## Running the netinfo CLI:
```
./netinfo [global options] command [command options] [arguments...]
```
## Help Command:
```
./netinfo
NAME:
   website lookup CLI -  Lets you Query IPs, CNAME, Records and Name servers

USAGE:
   netinfo [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
   ns       Looks up the Name server for the Host
   cname    Looks up the Cononical Name for the host
   ip       Looks up the IP address for the host
   mx       Looks up for Mail-Exchange Records for the given Host
   port     Looks up for the Ports within the given Host
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```
## Build With
1. Go Programming Language
2. Build CLI package - github.com/urfave/cli
