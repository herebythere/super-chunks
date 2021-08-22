import subprocess


def build_and_run_podman():
    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "down"])


if __name__ == "__main__":
    build_and_run_podman()
