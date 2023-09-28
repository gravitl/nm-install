### nm-install
# a tool to setup a netmaker server
````
installs netmaker (ce or pro)
checks and installs required dependencies (docker docker-compose wireguard)
obtains ssl certificates and starts netmaker server

creates network, installs netclient, joins network, sets default host, and creates egress gateway

Usage:
  nm-install [flags]
  nm-install [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     version information

Flags:
  -d, --domain string   custom domain to use
  -e, --email string    email to use for certificate registration
  -h, --help            help for nm-install
  -p, --pro             install pro version

Use "nm-install [command] --help" for more information about a command.
````
