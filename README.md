# Invokes

A rest api to handle users and invokes

## Setup 

### Get the code

Clone code from [https://github.com/Moth13/jump_invokes](https://github.com/Moth13/jump_invokes)

```zsh
git clone https://github.com/Moth13/jump_invokes.git
```

### Kanban

A kanban can be found at [Invokes Kaban](https://github.com/users/Moth13/projects/1/views/1)

### Configuration

```zsh
go mod tidy
```

## Local usage

### Build

```zsh
go build -o invokes -ldflags -s invokes/cmd/invokes
```

### Swag generation
Swagger documention can be generated using swaggo.
See [swaggo](https://github.com/swaggo/swag) for installation

Note that in docker usage, swagger generation is automatised.

```
swag init -o cmd/invokes/docs -d cmd/invokes,internal -g main.go
```

### Configuration file
You can find a configuration file sample and a template into the config directory.
The template file is used by the docker usage during the docker image build.

```zsh
./invokes -conf configs/invokes.yml
```

### Launch

```zsh
./invokes -conf configs/invokes.yml
```


## Docker usage


### Build

```zsh
Todo
```

### Run 

```zsh
Todo
```


## Unit testing

### Local testing

```zsh
Todo
```

### Docker testing

```zsh
Todo
```