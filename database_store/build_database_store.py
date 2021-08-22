import os
import json
import subprocess
from string import Template


def create_required_directories():
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
    db_conf = config["db"]
    server_conf = config["server"]
    config_conf = config["config"]

    create_template("templates/webapi.dockerfile.template",
                    "webapi/dockerfile", server_conf)

    create_template("templates/db.dockerfile.template",
                    "store/dockerfile", db_conf)

    create_template("templates/docker-compose.yml.template",
                    "docker-compose.yml",
                    {"service_name": config["service_name"],
                        "http_port": server_conf["http_port"],
                        "db_port": db_conf["port"],
                        "username": db_conf["username"],
                        "password": db_conf["password"],
                        "database_name": db_conf["database_name"],
                        "filepath": config_conf["filepath"],
                        "filepath_test": config_conf["filepath_test"]})


def build_database_store_with_podman():
    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "build"])


if __name__ == "__main__":
    create_required_directories()
    config = get_config("config/config.json")
    create_required_templates(config)
    build_database_store_with_podman()
