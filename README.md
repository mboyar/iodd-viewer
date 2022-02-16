[![Go](https://github.com/mboyar/iodd-viewer/actions/workflows/go.yml/badge.svg)](https://github.com/mboyar/iodd-viewer/actions/workflows/go.yml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=mboyar_iodd-viewer&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=mboyar_iodd-viewer)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=mboyar_iodd-viewer&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=mboyar_iodd-viewer)
[![Go Report Card](https://goreportcard.com/badge/github.com/mboyar/iodd-viewer)](https://goreportcard.com/report/github.com/mboyar/iodd-viewer)

## IODD Viewer
IODD files contain the device description details in XML format for the IO-Link-supported industrial sensors. This simple hobby project in golang is just converting the selected parts of this XML to a human-readable HTML table.

### Desktop
#### Build and run

```
$ go build -tags desktop -o bin/iodd-viewer-desktop ./...
$ bin/iodd-viewer-desktop --webview
```

### Cloud
The project's master branch is deployed to Heroku. View deployment: https://iodd-viewer.herokuapp.com/
