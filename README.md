# SSH MicroService

Use gRPC to manage remote SSH clients

[![Github Actions](https://img.shields.io/github/workflow/status/kain-lab/ssh-microservice/release?style=flat-square)](https://github.com/kain-lab/ssh-microservice/actions)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kain-lab/ssh-microservice?style=flat-square)](https://github.com/kain-lab/ssh-microservice)
[![Image Size](https://img.shields.io/docker/image-size/kainonly/ssh-microservice?style=flat-square)](https://hub.docker.com/r/kainonly/ssh-microservice)
[![Docker Pulls](https://img.shields.io/docker/pulls/kainonly/ssh-microservice.svg?style=flat-square)](https://hub.docker.com/r/kainonly/ssh-microservice)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://raw.githubusercontent.com/kain-lab/ssh-microservice/master/LICENSE)

## Setup

Example using docker compose

```yaml
version: "3.8"
services: 
  ssh:
    image: kainonly/ssh-microservice
    restart: always
    volumes:
      - ./ssh:/app/config
    ports:
      - 6000:6000
      - 8080:8080
```

## Configuration

For configuration, please refer to `config/config.example.yml` and create `config/config.yml`

- **debug** `string` Turn on debugging, that is `net/http/pprof`, and visit the address `http://localhost: 6060/debug/pprof`
- **listen** `string` grpc server listening address
- **gateway** `string` API gateway server listening address

## Service

The service is based on gRPC to view `api/api.proto`

```protobuf
syntax = "proto3";
package ssh;
option go_package = "ssh-microservice/gen/go/ssh";
import "google/protobuf/empty.proto";
import "google/api/annotations.proto";

service API {
  rpc Testing (Option) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      post: "/testing",
      body: "*",
    };
  }
  rpc Put (IOption) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      put: "/client",
      body: "*",
    };
  }
  rpc Exec (Bash) returns (Output) {
    option (google.api.http) = {
      post: "/exec",
      body: "*",
    };
  }
  rpc Delete (ID) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/client",
    };
  }
  rpc Get (ID) returns (Data) {
    option (google.api.http) = {
      get: "/client",
    };
  }
  rpc All (google.protobuf.Empty) returns (IDs) {
    option (google.api.http) = {
      get: "/clients",
    };
  }
  rpc Lists (IDs) returns (DataLists) {
    option (google.api.http) = {
      post: "/clients",
      body: "*"
    };
  }
  rpc Tunnels (TunnelsOption) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      put: "/tunnels",
      body: "*",
    };
  }
  rpc FreePort (google.protobuf.Empty) returns (Port) {
    option (google.api.http) = {
      get: "/free_port",
    };
  }
}

message Option {
  string host = 1;
  uint32 port = 2;
  string username = 3;
  string password = 4;
  string private_key = 5;
  string passphrase = 6;
}

message IOption {
  string id = 1;
  Option option = 2;
}

message Bash {
  string id = 1;
  string bash = 2;
}

message Output {
  bytes data = 1;
}

message ID {
  string id = 1;
}

message Data {
  string id = 1;
  string host = 2;
  uint32 port = 3;
  string username = 4;
  string connected = 5;
  repeated Tunnel tunnels = 6;
}

message Tunnel {
  string src_ip = 1;
  uint32 src_port = 2;
  string dst_ip = 3;
  uint32 dst_port = 4;
}

message IDs {
  repeated string ids = 1;
}

message DataLists {
  repeated Data data = 1;
}

message TunnelsOption {
  string id = 1;
  repeated Tunnel tunnels = 2;
}

message Port {
  uint32 data = 1;
}
```

## Testing (Option) returns (google.protobuf.Empty) {} 

test for ssh client connection

### RPC

- **Option**
  - **host** `string`
  - **port** `uint32`
  - **username** `string`
  - **password** `string` password (default empty)
  - **private_key** `string` private key (Base64)
  - **passphrase** `string` key passphrase (Base64)

```golang
client := pb.NewRouterClient(conn)
response, err := client.Testing(
  context.Background(),
  &pb.Option{
    Host:       debug.Host,
    Port:       debug.Port,
    Username:   debug.Username,
    Password:   debug.Password,
    PrivateKey: debug.PrivateKey,
    Passphrase: debug.Passphrase,
  },
)
```

### API Gateway

- **POST** `/testing`

```http
POST /testing HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
    "host":"dell",
    "port":22,
    "username":"root",
    "private_key":"LS0......o="
}
```

## Put (IOption) returns (google.protobuf.Empty) {}

Update the ssh client configuration to the service

### RPC

- **IOption**
  - **id** `string` ID
  - **option** `Option`
    - **host** `string`
    - **port** `uint32`
    - **username** `string`
    - **password** `string` password (default empty)
    - **private_key** `string` private key (Base64)
    - **passphrase** `string` key passphrase (Base64)

```golang
client := pb.NewRouterClient(conn)
response, err := client.Put(
  context.Background(),
  &pb.IOption{
    Id: "debug",
    Option: &pb.Option{
      Host:       debug.Host,
      Port:       debug.Port,
      Username:   debug.Username,
      Password:   debug.Password,
      PrivateKey: debug.PrivateKey,
      Passphrase: debug.Passphrase,
    },
  },
)
```

### API Gateway

- **PUT** `/client`

```http
PUT /client HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
    "id": "debug",
    "option": {
        "host": "dell",
        "port": 22,
        "username": "root",
        "private_key": "LS0......o="
    }
}
```

## Exec (Bash) returns (Output) {}

Send commands to the server via ssh

### RPC

- **Bash**
  - **id** `string` ID
  - **bash** `string` shell command
- **Output**
  - **data** command output result

```golang
client := pb.NewRouterClient(conn)
response, err := client.Exec(
    context.Background(),
    &pb.ExecParameter{
        Identity: "debug",
        Bash:     "uptime",
    },
)
```

### API Gateway

- **POST** `/exec`

```http
POST /exec HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
    "id": "debug",
    "bash": "uptime"
}
```

## Delete (ID) returns (google.protobuf.Empty) {}

Remove an ssh client from the service

### RPC

- **ID**
  - **id** `string` ID

```golang
client := pb.NewRouterClient(conn)
response, err := client.Delete(
  context.Background(),
  &pb.ID{
    Id: "debug",
  },
)
```

### API Gateway

- **DELETE** `/client`

```http
DELETE /client?id=debug HTTP/1.1
Host: localhost:8080
```

## Get (ID) returns (Data) {}

Get the details of an ssh client from the service

### RPC

- **ID**
  - **id** `string` ID
- **Data**
  - **id** `string` ID
  - **host** `string`
  - **port** `uint32`
  - **username** `string`
  - **connected** ssh connected client version
  - **tunnels** `[]Tunnel` ssh tunnels
    - **src_ip** `string` origin ip
    - **src_port** `uint32` origin port
    - **dst_ip** `string` target ip
    - **dst_port** `uint32` target port

```golang
client := pb.NewRouterClient(conn)
response, err := client.Get(
  context.Background(),
  &pb.ID{
    Id: "debug",
  },
)
```

### API Gateway

- **GET** `/client`

```http
GET /client?id=debug HTTP/1.1
Host: localhost:8080
```

## All (google.protobuf.Empty) returns (IDs) {}

Get all ssh client IDs from the service

### RPC

- **IDs**
  - **ids** `[]string` IDs

```golang
client := pb.NewRouterClient(conn)
response, err := client.All(
  context.Background(),
  &empty.Empty{},
)
```

### API Gateway

- **GET** `/clients`

```http
GET /clients HTTP/1.1
Host: localhost:8080
```

## Lists (IDs) returns (DataLists) {}

Get the specified list ssh client details from the service

### RPC

- **IDs**
  - **ids** `[]string` IDs
- **DataLists** 
  - **data** `[]Data`
    - **id** `string` ID
    - **host** `string`
    - **port** `uint32`
    - **username** `string`
    - **connected** ssh connected client version
    - **tunnels** `[]Tunnel` ssh tunnels
      - **src_ip** `string` origin ip
      - **src_port** `uint32` origin port
      - **dst_ip** `string` target ip
      - **dst_port** `uint32` target port

```golang
client := pb.NewRouterClient(conn)
response, err := client.Lists(
  context.Background(),
  &pb.IDs{
    Ids: []string{"debug", "debug-next"},
  },
)
```

### API Gateway

- **POST** `/clients`

```http
POST /clients HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Content-Length: 39

{
    "ids": [
        "debug"
    ]
}
```

## Tunnels (TunnelsOption) returns (google.protobuf.Empty) {}

Set up a tunnel for the ssh client

### RPC

- **TunnelsOption**
  - **id** `string` ID
  - **tunnels** `[]Tunnel` ssh tunnels
    - **src_ip** `string` origin ip
    - **src_port** `uint32` origin port
    - **dst_ip** `string` target ip
    - **dst_port** `uint32` target port

```golang
client := pb.NewRouterClient(conn)
response, err := client.Tunnels(
  context.Background(),
  &pb.TunnelsOption{
    Id: "debug-1",
    Tunnels: []*pb.Tunnel{
      {
        SrcIp:   "127.0.0.1",
        SrcPort: 9200,
        DstIp:   "127.0.0.1",
        DstPort: 9200,
      },
      {
        SrcIp:   "127.0.0.1",
        SrcPort: 5601,
        DstIp:   "127.0.0.1",
        DstPort: 5601,
      },
    },
  },
)
```

### API Gateway

- **PUT** `/tunnels`

```http
PUT /tunnels HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
    "id": "debug",
    "tunnels": [
        {
            "src_ip": "127.0.0.1",
            "src_port": 9200,
            "dst_ip": "127.0.0.1",
            "dst_port": 9200
        },
        {
            "src_ip": "127.0.0.1",
            "src_port": 5601,
            "dst_ip": "127.0.0.1",
            "dst_port": 5601
        }
    ]
}
```

## FreePort (google.protobuf.Empty) returns (Port) {}

Get available ports on the host

### RPC

- **Port**
  - **data** `uint32` port

```golang
client := pb.NewRouterClient(conn)
response, err := client.FreePort(
  context.Background(),
  &empty.Empty{},
)
```

### API Gateway

- **GET** `/free_port`

```http
GET /free_port HTTP/1.1
Host: localhost:8080
```