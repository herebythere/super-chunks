# SuperChunks

SuperChunks are configurable microservice components.


## Abstract

SuperChunks expose common microservice components like caches and databases as vanilla webapis.

SuperChunks are intended for:

- academic and educational resources
- foundational chunks for small teams and projects (just add auth!)

## How to use

### Prerequisites

SuperChunks assume you have `python3` and `podman` installed.

Run the following (or similar) commands:

```
dnf install podman
dnf install python3 python-pip3 podman
python3-pip3 install docker-compose
```

Then clone SuperChunks

```
git clone https://github.com/taylor-vann/superchunks/
```

### Build a Chunk

Each SuperChunk relies on JSON files for configuration.

Usually this coordination found in the following directory pattern:

```
<chunk_name>/
  /config
    config.json.example
  build_and_run_</chunk_name>.py
  README.md
```

Create a `config.json` file in the `config/` directory based on an example config in the same directory.

```
<chunk_name>/
  /config
    config.json.example
    config.json
  build_and_run_</chunk_name>.py
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

The `<chunk>.py` scripts generate all required files to build and deploy a particular chunk. However, files generated from this script could potentially expose sensitive data.

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

Apache License, Version 2.0