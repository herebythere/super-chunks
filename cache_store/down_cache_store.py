import subprocess


def down_cache_with_podman():
    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "down"])


if __name__ == "__main__":
    down_cache_with_podman()
