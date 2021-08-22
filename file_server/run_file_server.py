import subprocess


def run_file_server_with_podman():
    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "up", "--detach"])


if __name__ == "__main__":
    run_file_server_with_podman()
