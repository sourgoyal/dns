# DNS (Drone Navigation Service)
DNS helps drones to locate databank to upload data gathered from a sector of a galaxy.

# Install
### Run in a docker container 
```bash
docker build -t dns-ep .
docker run -it --publish 8080:8080 --name dnsApp --rm dns-ep
```
### Run in local
```bash 
. dep-ensure.sh
. install.sh
```

# Usage
```bash
URL: http://localhost:8080/getLoc

POST Method 
Input JSON
Example: {"x":"123.12","y":"456.56","z":"789.89","vel":"20.0"}
```

# North Star 
- Add monitoring using Prometheus and Grafana
- DNS serving multiple sectors at once
- Identification of HTTP client and serve as per their requirements
- Provide backword compatibility when upgrading logic
- Deploy using k8s and Helm

# License
MIT License



