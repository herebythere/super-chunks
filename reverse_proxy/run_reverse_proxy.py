import subprocess


def run_reverse_proxy_with_podman():
    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "up", "--detach"])


if __name__ == "__main__":
    run_reverse_proxy_with_podman()
