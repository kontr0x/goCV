# goCV
Blazingly fast minimalsitic resume generator with templating.

Templates are packaged within the binary for now.

## Building from source
```sh
go build main.go
```

## Usage
Show help message
```sh
./goCV --help
```

### With local latex installation
For demonstration using example
```sh
./goCV ./example.content.yaml
```

### With Docker
To run the application using Docker, you can use the provided `docker-compose.yml` file.
```sh
docker compose up
```
