README




```bash
docker build -t dns-ep .
docker run -it --publish 6080:8080 --name dnsApp--rm dns-ep
```