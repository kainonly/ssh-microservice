# SSH MicroService

Use gRPC to manage remote SSH clients

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kainonly/ssh-microservice?style=flat-square)](https://github.com/kainonly/ssh-microservice)
[![Travis](https://img.shields.io/travis/kainonly/ssh-microservice?style=flat-square)](https://www.travis-ci.org/kainonly/ssh-microservice)
[![Docker Pulls](https://img.shields.io/docker/pulls/kainonly/ssh-microservice.svg?style=flat-square)](https://hub.docker.com/r/kainonly/ssh-microservice)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://raw.githubusercontent.com/kainonly/ssh-microservice/master/LICENSE)

## Configuration

For configuration, please refer to `config/config.example.yml` and create `config/config.yml`

- **debug** `bool` Turn on debugging, that is `net/http/pprof`, and visit the address `http://localhost: 6060`
- **listen** `string` Microservice listening address
- **pool** `uint32` Memory pool in `MB`

## Service

The service is based on gRPC to view `router/router.proto`

```
syntax = "proto3";

service Router {
    rpc Testing (TestingParameter) returns (Response) {
    }

    rpc Put (PutParameter) returns (Response) {
    }

    rpc Exec (ExecParameter) returns (ExecResponse) {
    }

    rpc Delete (DeleteParameter) returns (Response) {
    }

    rpc Get (GetParameter) returns (GetResponse) {
    }

    rpc All (NoParameter) returns (AllResponse) {
    }

    rpc Lists (ListsParameter) returns (ListsResponse) {
    }

    rpc Tunnels (TunnelsParameter) returns (Response) {
    }
}

message NoParameter {
}

message Response {
    uint32 error = 1;
    string msg = 2;
}

message TestingParameter {
    string host = 1;
    uint32 port = 2;
    string username = 3;
    string password = 4;
    string private_key = 5;
    string passphrase = 6;
}

message PutParameter {
    string identity = 1;
    string host = 2;
    uint32 port = 3;
    string username = 4;
    string password = 5;
    string private_key = 6;
    string passphrase = 7;
}

message ExecParameter {
    string identity = 1;
    string bash = 2;
}

message ExecResponse {
    uint32 error = 1;
    string msg = 2;
    string data = 3;
}

message DeleteParameter {
    string identity = 1;
}

message GetParameter {
    string identity = 1;
}

message GetResponse {
    uint32 error = 1;
    string msg = 2;
    Information data = 3;
}

message Information {
    string identity = 1;
    string host = 2;
    uint32 port = 3;
    string username = 4;
    string connected = 5;
    repeated TunnelOption tunnels = 6;
}

message TunnelOption {
    string src_ip = 1;
    uint32 src_port = 2;
    string dst_ip = 3;
    uint32 dst_port = 4;
}

message AllResponse {
    uint32 error = 1;
    string msg = 2;
    repeated string data = 3;
}

message ListsParameter {
    repeated string identity = 1;
}

message ListsResponse {
    uint32 error = 1;
    string msg = 2;
    repeated Information data = 3;
}

message TunnelsParameter {
    string identity = 1;
    repeated TunnelOption tunnels = 2;
}
```