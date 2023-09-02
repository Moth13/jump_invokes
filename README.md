# Invokes

A rest api to handle users and invokes

## Introduction
This micro service has been developped to show my capacties in the goland development.

It is based on a boilerplate I've done since I've coded some RestAPI in the past.

I choose to use gin as HTTP web framework, for its perfomance, easiness and capacities.
For the database part, I choose gorm as it handles easily all databases, very pratical to use.

### Though on the coding
Router class is here as a wrapper, gin can be easily swith to mux if wanted.

Code has been cut in logical directory for convenience.

During its development, I've encountered no specific difficulties. I've started by filling the kanban with the task I've thought, then follow the tasks' order I've set.

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
see [config/invokes.yml.sample](config/invokes.yml.sample) for more info


### Launch
You can launch the db given by jump:
```zsh
git clone https://github.com/Freelance-launchpad/backend-interviews.git
cd backend-interviews/database
docker build . -t jump-database
docker run -p 5432:5432 jump-database                   
```
Then launch the service using: 
```zsh
./invokes -conf configs/invokes.yml
```


## Docker usage


### Build

```zsh
docker build -f deployment/Dockerfile -t invokes:0.0.0 .
```

### Run 

### Docker image env variables
The docker image has some env variables to make it configurable
```zsh
DB_ENGINE: postgresql // can be mysql, postgresql
DB_PROTO: postgres:// //postgres://, mysql+pymysql://
DB_USER: jump
DB_PASSWD: password
DB_PORT: 5432
DB_OPTS:
DB_NAME: 
DB_HOST: dbserver
LOG_LEVEL: error // loglevel from debug, info, warning, error (warning and error as realease mode)
CAT_CONF: "true" // true to display the config at boot
```

### Docker-Compose
A docker-compose file is available in the deployment directory
```zsh
docker-compose -f deployment/docker-compose.yml
```
It will launch the db and the service


## Unit testing

### Local testing
You can launch some unit test by:
```zsh
go test -v invokes/internal/...
```