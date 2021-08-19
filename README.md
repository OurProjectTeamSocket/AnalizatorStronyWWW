## 🚀 MonitorAI 🚀 🚀 🔴 🔵 ✅ ✅

## Running on your local machine

Linux or MacOS

# Requirement 
```
Go 
PostgreSQL
```
## Installation guide
#### 1. install go version 1.14++
```bash
# please read this link installation guide of go
# https://golang.org/doc/install
```


#### 3. Build the application
```bash
# run command :
go mod tidy && go mod download && go mod vendor
go run main.go
```

#### 4. Healthcheck
```bash
curl --request GET \
  --url http://localhost:8080
```