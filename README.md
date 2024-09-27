# Generated With [Spawn](https://github.com/rollchains/spawn)

## Module Scaffolding

- `spawn module new <name>` *Generates a Cosmos module template*

## Content Generation

- `make proto-gen` *Generates golang code from proto files, stubs interfaces*

## Testnet

- `make testnet` *IBC testnet from chain <-> local cosmos-hub*
- `make sh-testnet` *Single node, no IBC. quick iteration*

## Local Images

- `make install`      *Builds the chain's binary*
- `make local-image`  *Builds the chain's docker image*

## Testing

- `go test ./... -v` *Unit test*
- `make ictest-*`  *E2E testing*



Spawn - Example
spawn new rollchain --bech32=roll --denom=uroll --bin=rolld --org=rollchains

Select default options


cd rollchain/



spawn module new nameservice



query.proto


  // ResolveName allows a user to resolve the name of an account.
  rpc GetEvent(QueryGetEventRequest) returns (QueryGetEventResponse) {
    option (google.api.http).get = "/ems/v1/name/{organizer}";
  }
  
  
  message QueryGetEventRequest {
  string organizer = 1;
}

message QueryGetEventResponse {
  string name = 1;
}



tx.proto


  rpc CreateEvent(MsgCreateEvent) returns (MsgCreateEventResponse);
  
  
  
  // MsgSetServiceName defines the structure for setting a name.
message MsgCreateEvent {
  option (cosmos.msg.v1.signer) = "organizer";

  string organizer = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];

  string name = 2;
}

// MsgSetServiceNameResponse is an empty reply.
message MsgCreateEventResponse {}



make proto-gen


keeper.go


update keeper struct

EventMapping collections.Map[string, string]

EventMapping: collections.NewMap(sb, collections.NewPrefix(1), "event_mapping", collections.StringKey, collections.StringValue),



msg_server.go



// CreateEvent implements types.MsgServer.
func (ms msgServer) CreateEvent(ctx context.Context, msg *types.MsgCreateEvent) (*types.MsgCreateEventResponse, error) {
    if err := ms.k.EventMapping.Set(ctx, msg.Organizer, msg.Name); err != nil {
        return nil, err
    }
    return &types.MsgCreateEventResponse{}, nil
}



query_server.go


// GetEvent implements types.QueryServer.
func (k Querier) GetEvent(goCtx context.Context, req *types.QueryGetEventRequest) (*types.QueryGetEventResponse, error) {
    v, err := k.Keeper.EventMapping.Get(goCtx, req.Organizer)
    if err != nil {
        return nil, err
    }

    return &types.QueryGetEventResponse{
        Name: v,
    }, nil
}



autocli.go


func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
    return &autocliv1.ModuleOptions{
        Query: &autocliv1.ServiceCommandDescriptor{
            Service: modulev1.Query_ServiceDesc.ServiceName,
            RpcCommandOptions: []*autocliv1.RpcCommandOptions{
                {
                    RpcMethod: "GetEvent",
                    Use:       "get [organizer]",
                    Short:     "Get event by organizer",
                    PositionalArgs: []*autocliv1.PositionalArgDescriptor{
                        {ProtoField: "organizer"},
                    },
                },
                {
                    RpcMethod: "Params",
                    Use:       "params",
                    Short:     "Query the current consensus parameters",
                },
            },
        },
        Tx: &autocliv1.ServiceCommandDescriptor{
            Service: modulev1.Msg_ServiceDesc.ServiceName,
            RpcCommandOptions: []*autocliv1.RpcCommandOptions{
                {
                    RpcMethod: "CreateEvent",
                    Use:       "create [name]",
                    Short:     "Create an event",
                    PositionalArgs: []*autocliv1.PositionalArgDescriptor{
                        {ProtoField: "name"},
                    },
                },
                {
                    RpcMethod: "UpdateParams",
                    Skip:      false, // set to true if authority gated
                },
            },
        },
    }
}



make sh-testnet






vitwit@vitwit-Vostro-16-5630:~/prathyusha/eventchain$ eventd tx -h
Transactions subcommands

Usage:
  eventd tx [flags]
  eventd tx [command]

Available Commands:
  auth                   Transactions commands for the auth module
  authz                  Authorization transactions subcommands
  bank                   Bank transaction subcommands
  broadcast              Broadcast transactions generated offline
  circuit                Transactions commands for the circuit module
  consensus              Transactions commands for the consensus module
  crisis                 Transactions commands for the crisis module
  decode                 Decode a binary encoded transaction string
  distribution           Distribution transactions subcommands
  encode                 Encode transactions generated offline
  etm                    Transactions commands for the etm module
  evidence               Evidence transaction subcommands
  feegrant               Feegrant transactions sub-commands
  gov                    Governance transactions subcommands
  group                  Group transaction subcommands
  ibc                    IBC transaction subcommands
  ibc-fee                IBC relayer incentivization transaction subcommands
  ibc-transfer           IBC fungible token transfer transaction subcommands
  interchain-accounts    IBC interchain accounts transaction subcommands
  mint                   Transactions commands for the mint module
  multi-sign             Generate multisig signatures for transactions generated offline
  multisign-batch        Assemble multisig transactions in batch from batch signatures
  nft                    Transactions commands for the nft module
  packetfowardmiddleware Transactions commands for the packetfowardmiddleware module
  poa                    poatransaction subcommands
  sign                   Sign a transaction generated offline
  sign-batch             Sign transaction batch files
  simulate               Simulate the gas usage of a transaction
  slashing               Transactions commands for the slashing module
  staking                Staking transaction subcommands
  tokenfactory           tokenfactory transactions subcommands
  upgrade                Upgrade transaction subcommands
  validate-signatures    validate transactions signatures
  vesting                Vesting transaction subcommands

Flags:
  -h, --help   help for tx

Global Flags:
      --home string         directory for config and data (default "/home/vitwit/.eventchain")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic|disabled or '*:<level>,<key>:<level>') (default "info")
      --log_no_color        Disable colored logs
      --trace               print out full stack trace on errors

Use "eventd tx [command] --help" for more information about a command.
vitwit@vitwit-Vostro-16-5630:~/prathyusha/eventchain$ eventd tx etm -h
Transactions commands for the etm module

Usage:
  eventd tx etm [flags]
  eventd tx etm [command]

Available Commands:
  create        Create an event
  update-params Execute the UpdateParams RPC method

Flags:
  -h, --help   help for etm

Global Flags:
      --home string         directory for config and data (default "/home/vitwit/.eventchain")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic|disabled or '*:<level>,<key>:<level>') (default "info")
      --log_no_color        Disable colored logs
      --trace               print out full stack trace on errors

Use "eventd tx etm [command] --help" for more information about a command.
vitwit@vitwit-Vostro-16-5630:~/prathyusha/eventchain$ eventd tx etm create -h
Create an event

Usage:
  eventd tx etm create [name] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
      --aux                      Generate aux signer data instead of sending a tx
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async) (default "sync")
      --chain-id string          The network chain ID
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it (when enabled, the local Keybase is not accessible)
      --fee-granter string       Fee granter grants fees for the transaction
      --fee-payer string         Fee payer pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically. Note: "auto" option doesn't always report accurate results. Set a valid coin value to adjust the result. Can be used instead of "fees". (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase only accessed when providing a key name)
  -h, --help                     help for create
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to CometBFT rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality)
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json|direct-aux|textual), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
      --tip string               Tip is the amount that is going to be transferred to the fee payer on the target chain. This flag is only valid when used with --aux, and is ignored if the target chain didn't enable the TipDecorator
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --home string         directory for config and data (default "/home/vitwit/.eventchain")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic|disabled or '*:<level>,<key>:<level>') (default "info")
      --log_no_color        Disable colored logs
      --trace               print out full stack trace on errors
vitwit@vitwit-Vostro-16-5630:~/prathyusha/eventchain$ eventd tx etm create vitwit-event-1 --from=acc0
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /etm.v1.MsgCreateEvent
    name: vitwit-event-1
    organizer: event1hj5fveer5cjtn4wd6wstzugjfdxzl0xp0a2w8n
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: ""
timestamp: ""
tx: null
txhash: B233F3E33F53D5B9993F61217D3D83EA1B2BC4F1BD0CFA5EBDFB21F9DA35FE2D
vitwit@vitwit-Vostro-16-5630:~/prathyusha/eventchain$ eventd q 
Querying subcommands

Usage:
  eventd query [flags]
  eventd query [command]

Aliases:
  query, q

Available Commands:
  auth                Querying commands for the auth module
  authz               Querying commands for the authz module
  bank                Querying commands for the bank module
  block               Query for a committed block by height, hash, or event(s)
  block-results       Query for a committed block's results by height
  blocks              Query for paginated blocks that match a set of events
  circuit             Querying commands for the circuit module
  consensus           Querying commands for the consensus module
  distribution        Querying commands for the distribution module
  etm                 Querying commands for the etm module
  evidence            Querying commands for the evidence module
  feegrant            Querying commands for the feegrant module
  gov                 Querying commands for the gov module
  group               Querying commands for the group module
  ibc                 Querying commands for the IBC module
  ibc-fee             IBC relayer incentivization query subcommands
  ibc-transfer        IBC fungible token transfer query subcommands
  interchain-accounts IBC interchain accounts query subcommands
  mint                Querying commands for the mint module
  nft                 Querying commands for the nft module
  packetforward       Querying commands for the packetforward module
  params              Querying commands for the params module
  poa                 Querying commands for the poa module
  slashing            Querying commands for the slashing module
  staking             Querying commands for the staking module
  tokenfactory        Querying commands for the tokenfactory module
  tx                  Query for a transaction by hash, "<addr>/<seq>" combination or comma-separated signatures in a committed block
  txs                 Query for paginated transactions that match a set of events
  upgrade             Querying commands for the upgrade module
  wait-tx             Wait for a transaction to be included in a block

Flags:
  -h, --help   help for query

Global Flags:
      --home string         directory for config and data (default "/home/vitwit/.eventchain")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic|disabled or '*:<level>,<key>:<level>') (default "info")
      --log_no_color        Disable colored logs
      --trace               print out full stack trace on errors

Use "eventd query [command] --help" for more information about a command.
vitwit@vitwit-Vostro-16-5630:~/prathyusha/eventchain$ eventd q etm -h
Querying commands for the etm module

Usage:
  eventd query etm [flags]
  eventd query etm [command]

Available Commands:
  get         Get event by organizer
  params      Query the current consensus parameters

Flags:
  -h, --help   help for etm

Global Flags:
      --home string         directory for config and data (default "/home/vitwit/.eventchain")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic|disabled or '*:<level>,<key>:<level>') (default "info")
      --log_no_color        Disable colored logs
      --trace               print out full stack trace on errors

Use "eventd query etm [command] --help" for more information about a command.
vitwit@vitwit-Vostro-16-5630:~/prathyusha/eventchain$ eventd q etm get -h
Get event by organizer

Usage:
  eventd query etm get [organizer] [flags]

Flags:
      --grpc-addr string   the gRPC endpoint to use for this chain
      --grpc-insecure      allow gRPC over insecure channels, if not the server must use TLS
      --height int         Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help               help for get
      --no-indent          Do not indent JSON output
      --node string        <host>:<port> to CometBFT RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string      Output format (text|json) (default "text")

Global Flags:
      --home string         directory for config and data (default "/home/vitwit/.eventchain")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic|disabled or '*:<level>,<key>:<level>') (default "info")
      --log_no_color        Disable colored logs
      --trace               print out full stack trace on errors
vitwit@vitwit-Vostro-16-5630:~/prathyusha/eventchain$ eventd q etm get event1hj5fveer5cjtn4wd6wstzugjfdxzl0xp0a2w8n
name: vitwit-event-1

vitwit@vitwit-Vostro-16-5630:~/prathyusha/eventchain$ ^C

