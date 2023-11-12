# Grender - Web Page Renderer

❗ **This project is in a very early stage of development.** ❗

This is a Go-based rendering server. It uses different backends for storage and is configured via a YAML file.

## Configuration

The server is configured using a `config.yaml` file. Here's an example:

```yaml
version: v1
renderingConfig:
  pageWaitTime: 1000
  pageWaitCondition: '(function() { return window.prerenderReady === true })()'
server:
  port: "8080"
backend:
  s3:
    bucketName: "grender.io"
    region: "ap-southeast-2"

```

## Running the Server

To run the server, use the following command:

```bash
go run main.go
```

Make sure to have your config.yaml file in the same directory where you run this command.

```bash
go get -u github.com/chromedp/chromedp
go get -u github.com/gin-gonic/gin
```

## Project Structure

The main components of the project are organized into packages:

- **pilot:** Contains the main Grender application files, including configuration, rendering logic, and server setup.
- **backend:** Provides interfaces and implementations for different backends, such as S3 and file systems.
- **render:** Implements the rendering logic using a headless Chrome browser.

## Dependencies

Grender relies on third-party packages for AWS S3 integration, YAML parsing, and headless Chrome browser automation.
