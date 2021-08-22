import subprocess

def create_self_signed_certificates():
    subprocess.run(["openssl", "req", "-x509", "-nodes", "-newkey",
                    "rsa:4096", "-keyout", "./config/privkey.pem", "-out",
                    "./config/fullchain.pem", "-days", "365", "-subj",
                    "/C=US/ST=California/L=San Francisco/O=Company/OU=Org/CN=super-reverse-proxy.com"])


if __name__ == "__main__":
    create_self_signed_certificates()
