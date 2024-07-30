# Redis Clone

This is a simple Redis clone implemented in Go. It supports basic Redis commands and uses an in-memory data store.

## Features

- In-memory key-value store
- Support for string commands (GET, SET)
- RESP (Redis Serialization Protocol) implementation
- TCP server

## Usage

To run the server:

`go run cmd/server/main.go`

The server will start listening on port 6379 by default.

To connect to the server, use redis-cli:

`redis-cli -p 6379`

```
SET mykey "Hello, Redis!"
GET mykey
```


## Supported Commands

- GET
- SET

More commands will be added in future updates.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.