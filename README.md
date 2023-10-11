# vatz-plugin-sei
Vatz plugin for sei node monitoring

## Plugins
- node_block_sync : monitor block sync status
- node_is_alived : monitor process running status
- node_peer_count : monitor the number of peers
- node_active_status : monitor the validator include in active set
- node_governance_alarm : monitor the new governance proposal and whether or not to vote
- pfd_status: monitor the price-feeder oracle status

## Installation and Usage
> Please make sure [Vatz](https://github.com/dsrvlabs/vatz) is running with proper configuration. [Vatz Installation Guide](https://github.com/dsrvlabs/vatz/blob/main/docs/installation.md)

### Install Plugins
- Install with source
```
$ git clone https://github.com/dsrvlabs/vatz-plugin-sei.git
$ cd vatz-plugin-sei
$ make install
```
- Install with Vatz CLI command
```
$ ./vatz plugin install --help
Install new plugin

Usage:
   plugin install [flags]

Examples:
vatz plugin install github.com/dsrvlabs/<somewhere> name

Flags:
  -h, --help   help for install
```
> please make sure install path for the plugins repository URL.
```
$ ./vatz plugin install github.com/dsrvlabs/vatz-plugin-sei/plugins/node_block_sync node_block_sync
$ ./vatz plugin install github.com/dsrvlabs/vatz-plugin-sei/plugins/node_is_alived node_is_alived
$ ./vatz plugin install github.com/dsrvlabs/vatz-plugin-sei/plugins/node_peer_count node_peer_count
$ ./vatz plugin install github.com/dsrvlabs/vatz-plugin-sei/plugins/node_active_status node_active_status
$ ./vatz plugin install github.com/dsrvlabs/vatz-plugin-sei/plugins/node_governance_alarm node_governance_alarm
$ ./vatz plugin install github.com/dsrvlabs/vatz-plugin-sei/plugins/pfd_status pfd_status
```
- Check plugins list with Vatz CLI command
```
$ vatz plugin list
+-----------------------+------------+---------------------+-------------------------------------------------------------------+---------+
| NAME                  | IS ENABLED | INSTALL DATE        | REPOSITORY                                                        | VERSION |
+-----------------------+------------+---------------------+-------------------------------------------------------------------+---------+
| node_block_sync       | true       | 2023-09-27 01:14:53 | github.com/dsrvlabs/vatz-plugin-sei/plugins/node_block_sync       | latest  |
| node_is_alived        | true       | 2023-09-27 01:15:41 | github.com/dsrvlabs/vatz-plugin-sei/plugins/node_is_alived        | latest  |
| node_peer_count       | true       | 2023-09-27 01:15:46 | github.com/dsrvlabs/vatz-plugin-sei/plugins/node_peer_count       | latest  |
| node_active_status    | true       | 2023-09-27 01:15:51 | github.com/dsrvlabs/vatz-plugin-sei/plugins/node_active_status    | latest  |
| node_governance_alarm | true       | 2023-09-27 01:15:59 | github.com/dsrvlabs/vatz-plugin-sei/plugins/node_governance_alarm | latest  |
| pfd_status            | true       | 2023-09-27 01:16:16 | github.com/dsrvlabs/vatz-plugin-sei/plugins/pfd_status            | latest  |
+-----------------------+------------+---------------------+-------------------------------------------------------------------+---------+
```

### Run
> Run as default config or option flags
```
$ node_block_sync
2023-05-31T07:07:36Z INF Register module=grpc
2023-05-31T07:07:36Z INF Start 127.0.0.1 10001 module=sdk
2023-05-31T07:07:36Z INF Start module=grpc
2023-05-31T07:08:10Z INF Execute module=grpc
2023-05-31T07:08:10Z INF previous block height: 0, latest block height: 5969512 module=plugin
2023-05-31T07:08:10Z DBG block height increasing module=plugin
```
```
$ node_is_alived
2023-05-31T07:07:36Z INF Register module=grpc
2023-05-31T07:07:36Z INF Start 127.0.0.1 10002 module=sdk
2023-05-31T07:07:36Z INF Start module=grpc
2023-05-31T07:08:10Z INF Execute module=grpc
2023-05-31T07:08:10Z INF HEALTHY process=up
2023-05-31T07:08:40Z INF Execute module=grpc
2023-05-31T07:08:40Z INF HEALTHY process=up
```
```
$ node_peer_count
2023-05-31T07:07:36Z INF Register module=grpc
2023-05-31T07:07:36Z INF Start 127.0.0.1 10003 module=sdk
2023-05-31T07:07:36Z INF Start module=grpc
2023-05-31T07:08:10Z INF Execute module=grpc
2023-05-31T07:08:10Z INF Good: peer_count is 50 moudle=plugin
2023-05-31T07:08:40Z INF Execute module=grpc
2023-05-31T07:08:40Z INF Good: peer_count is 50 moudle=plugin
```
```
$ node_active_status -valoperAddr <VALIDATOR_OPERATOR_ADDRESS>
2023-05-31T07:07:36Z INF Register module=grpc
2023-05-31T07:07:36Z INF Start 127.0.0.1 10004 module=sdk
2023-05-31T07:07:36Z INF Start module=grpc
2023-05-31T07:08:10Z INF Execute module=grpc
2023-05-31T07:08:10Z DBG Validator bonded. included active set module=plugin
2023-05-31T07:08:40Z INF Execute module=grpc
2023-05-31T07:08:40Z DBG Validator bonded. included active set module=plugin
```
```
# Your node have to enable API configuration ({HOME_DIR}/config/app.toml)
$ node_governance_alarm -apiPort <API server port{default is 1317}> -voterAddr <Account Address>
2023-05-31T07:07:36Z INF Register module=grpc
2023-05-31T07:07:36Z INF Start 127.0.0.1 10005 module=sdk
2023-05-31T07:07:36Z INF Start module=grpc
2023-05-31T07:08:10Z INF Execute module=grpc
2023-05-31T07:08:10Z DBG DEBUG : tmp == proposalId module=plugin
2023-05-31T07:08:10Z INF Lastest proposal is #51
 module=plugin
```

```
$ pfd_status -port <API server port> -valoperAddr <Valoper Address> -seiHome <Home PATH>
2023-09-26T02:04:52Z INF Register module=grpc
2023-09-26T02:04:52Z INF Start 127.0.0.1 10006 module=sdk
2023-09-26T02:04:52Z INF Start module=grpc
2023-09-26T02:05:22Z INF Execute module=grpc
2023-09-26T02:05:22Z DBG Price-Feeder oracle missing rate: 0.90%
 module=plugin
```
## Command line arguments
- node_block_sync
```
Usage of node_block_sync:
  -addr string
	Listening address (default "127.0.0.1")
  -critical int
	block height stucked count to raise critical level of alert (default 3)
  -port int
	Listening port (default 10001)
  -rpcURI string
	Tendermint RPC URI Address (default "http://localhost:26657")
```
- node_is_alived
```
Usage of node_is_alived:
  -addr string
    	IP Address(e.g. 0.0.0.0, 127.0.0.1) (default "127.0.0.1")
  -port int
    	Port number (default 10002)
  -rpcAddr string
    	RPC addrest:port (e.g. http://127.0.0.1:26667) (default "http://localhost:26657")
```
- node_peer_count
```
Usage of node_peer_count:
  -addr string
        IP Address(e.g. 0.0.0.0, 127.0.0.1) (default "127.0.0.1")
  -minPeer int
        minimum peer count, default 5 (default 5)
  -port int
        Port number (default 10003)
  -rpcAddr string
    	RPC Address, default http://localhost:26657 (default "https://localhost:26657")
```
- node_active_status
```
Usage of node_active_status:
  -addr string
    	Listening address (default "127.0.0.1")
  -port int
    	Listening port (default 10004)
  -rpcURI string
    	CosmosHub RPC URI Address (default "http://localhost:1317")
  -valoperAddr string
    	CosmosHub validator operator address
```
- node_governance_alarm
```
Usage of node_governance_alarm:
  -addr string
    	IP Address(e.g. 0.0.0.0, 127.0.0.1) (default "127.0.0.1")
  -apiPort uint
    	Need to know proposal id (default 1317)
  -port int
    	Port number (default 10005)
  -proposalId uint
    	Need to know last proposal id
  -voterAddr string
    	Need to voter address (default "address")
```

- pfd_status
```
Usage of pfd_status:
  -addr string
    	IP Address(e.g. 0.0.0.0, 127.0.0.1) (default "127.0.0.1")
  -port int
    	Port number (default 10006)
  -seiHome
	HOME PATH
  -valoperAddr string
    	Need to valoperAddress address (default "address")
```
## TroubleShooting
1. Encountered issue related with `Device or Resource Busy` or `Too many open files` error.
 - Check your open file limit and recommended to increase it.
 ```
 $ ulimit -n
 1000000
 ```

## License

`vatz-plugin-sei` is licensed under the [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also included in our repository in the `LICENSE` file.
