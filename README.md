## IODD Viewer
IODD files contains the device description details in XML format for the IO-Link supported industrial sensors. This simple hobby project in golang is just converting the selected parts of this XML to human-readble HTML table.

### Desktop
#### Build and run

```
$ go build -tags desktop -o bin/iodd-viewer-desktop ./...
$ bin/iodd-viewer-desktop --webview
```

### Cloud
The project's master branch is deployed to heroku. View deployment: https://iodd-viewer.herokuapp.com/
