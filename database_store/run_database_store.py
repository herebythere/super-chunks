import subprocess


def run_with_database_store():
    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "up", "--detach"])


if __name__ == "__main__":
    run_with_database_store()
