## ping matrix
- Supply two tools.
  1. rpc_server: Will listen on <IP:port> and send fping result to databases.
  2. rpc_client: Will report fping result and send fping result to rpc_server by jsonrpc.
- install
  - cd ping-matrix/rpc_server && go mod tidy && go build
  - cd ping-matrix/rpc_client && go mod tidy && go build
- run
  - ./rpc_server --help
  - ./rpc_client --help
