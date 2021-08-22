import json
import subprocess
from string import Template


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
    config_conf = config["config"]
    file_conf = config["fileserver"]
    server_conf = config["server"]

    compose_conf = {"service_name": config["service_name"],
                    "http_port": config["server"]["http_port"],
                    "fileserver_filepath": file_conf["filepath"],
                    "not_found_filepath": file_conf["not_found"],
                    "bad_request_filepath": file_conf["bad_request"],
                    "config_filepath": config_conf["filepath"],
                    "config_filepath_test": config_conf["filepath_test"]}

    create_template("templates/webapi.dockerfile.template",
                    "webapi/dockerfile", server_conf)

    create_template("templates/docker-compose.yml.template",
                    "docker-compose.yml",
                    compose_conf)


def build_file_server_with_podman():
    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "build"])


if __name__ == "__main__":
    config = get_config("config/config.json")
    create_required_templates(config)
    build_file_server_with_podman()
