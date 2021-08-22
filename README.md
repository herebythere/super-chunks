# SuperChunks

SuperChunks are configurable microservice components made available as vanilla web apis.

## Abstract

SuperChunks isolate common microservice components in a compositional manner.

Current chunks include:

- cache_store (redis)
- database_store (postgres)
- reverse_proxy
- file_server

## How to use

### Prerequisites

SuperChunks assume you have `python3` and `podman` installed.

Run the following (or similar) commands:

```
dnf install podman
dnf install python3 python-pip3 podman
python3-pip3 install docker-compose
```

Clone SuperChunks

```
git clone https://github.com/taylor-vann/superchunks/
```

Review the code! You'd be surprised how many engineers don't peek 
at the code they're using.

SuperChunks does not have any observational support or logging. 
You are completely on your own! And this is a good thing.

SuperChunks are bare metal implmentations of components used to 
compose a microservice exposed as web apis. They help encourage 
_good faith_ standards and decouple databases and caches from
web apis. 

### What is a chunk?

Each SuperChunk relies on JSON files for configuration.

Each chunk could potentially have its own particular set of 
instructions.

But the common directory structure between chunks is as follows:

```
<chunk_name>/
  /config
    config.json.example
  build_</chunk_name>.py
  run_</chunk_name>.py
  down_</chunk_name>.py
  README.md
```

### Build a Chunk

Create a `config.json` file in the `config/` directory based on the corresponding `config.json.example`.

```
<chunk_name>/
  /config
    config.json.example
    config.json
  build_</chunk_name>.py
  run_</chunk_name>.py
  down_</chunk_name>.py
  README.md
```

Build a chunk by running:

```
python3 build_</chunk_name>.py
```

Deploy a chunk by running:

```
python3 run_</chunk_name>.py
```

Down a chunk by running:

```
python3 down_</chunk_name>.py
```

### Sensitive Data

Micro-services have a tendency to expose sensitive information like passwords, certificates, and API keys.

Don't do that.

SuperChunks uses `.gitignore` to hide all configuration files.

The `.py` scripts generate all required files to build and deploy a particular chunk. However, these generated files could potentially expose sensitive data.

All generated and environment artifacts are ignored in SuperChunks. 

This includes to the following file types:

- .env
- .txt
- .crt
- .key
- .pem
- .json
- .conf
- .rdb
- .cf
- dockerfile
- docker-compose.yml

I also recommend that you use a similar approach.

## License

SuperChunks are intended for:

- academic and educational resources
- small teams and projects

SuperChunks is licensed under the _Apache License, Version 2.0_