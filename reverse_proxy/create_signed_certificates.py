import subprocess

def create_signed_certificates_with_lets_encrypt():
    subprocess.run(["openssl", "req", "-x509", "-nodes", "-newkey", "rsa:4096", "-keyout", "key.pem", "-out",
                   "fullchain.pem", "-days", "365", "-subj", "/C=US/ST=California/L=San Francisco/O=Company/OU=Org/CN=unsigned.reverse.proxy.com"])

# copy lets encrypt over to config folder
  

if __name__ == "__main__":
    create_signed_certificates_with_lets_encrypt()
