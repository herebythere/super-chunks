import os
import json
import subprocess
from string import Template


def create_required_directories():
    if not os.path.exists("store/conf"):
        os.makedirs("store/conf")
    if not os.path.exists("store/data"):
        os.makedirs("store/data")


def get_config(source):
    config_file = open(source, 'r')
    config = json.load(config_file)
    config_file.close()

    return config


def create_template(source, target, keywords):
    source_file = open(source, 'r')
    source_file_template = Template(source_file.read())
    source_file.close()
    updated_source_file_template = source_file_template.substitute(**keywords)

    target_file = open(target, "w+")
    target_file.write(updated_source_file_template)
    target_file.close()


def create_required_templates(config):
    cache_conf = config["cache"]
    config_conf = config["config"]
    server_conf = config["server"]

    compose_conf = {"service_name": config["service_name"],
                    "http_port": server_conf["http_port"],
                    "filepath": config_conf["filepath"],
                    "filepath_test": config_conf["filepath_test"]}

    create_template("templates/webapi.dockerfile.template",
                    "webapi/dockerfile", server_conf)

    create_template("templates/cache.dockerfile.template",
                    "store/dockerfile", cache_conf)

    create_template("templates/redis.conf.template",
                    "store/conf/redis.conf", cache_conf)

    create_template("templates/docker-compose.yml.template",
                    "docker-compose.yml",
                    compose_conf)


def build_cache_with_podman():
    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "build"])


if __name__ == "__main__":
    create_required_directories()
    config = get_config("config/config.json")
    create_required_templates(config)
    build_cache_with_podman()
