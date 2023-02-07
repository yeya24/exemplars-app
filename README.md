# exemplars-app

## How to use

App will start at localhost:8080.
```bash
go build -o main main.go
./main
```

## Prometheus Setup

Example Prometheus configuration file.

```yaml
# my global config
global:
  scrape_interval: 15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.

# A scrape configuration containing exactly one endpoint to scrape:
scrape_configs:
  - job_name: "app"
    static_configs:
    - targets: ["localhost:8080"]
```

Start local Prometheus.

```bash
./prometheus --enable-feature=exemplar-storage
```
