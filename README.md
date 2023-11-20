# Grender - Web Page Renderer

❗ **This project is in a very early stage of development.** ❗

Grender is a web rendering service written in Go that leverages a headless Chrome browser to render and cache web pages. The application supports various backends, including AWS S3 and local file systems, for storing cached content.

## Configuration

The server is configured using a `config.yaml` file. Here's an example:

```yaml
version: v1
# There are three modes: "rendering" and "recaching"
modes: 
  rendering: true
  recaching: true
renderingConfig:
  pageWaitTime: 10
  pageWaitCondition: '(function() { return window.prerenderReady === true })()'
  # requestHeaders:
  #   - name: "X-Grender-Request"
  #     value: "1"
server:
  port: "8080"
  responseHeaders:
    - name: "X-Prerender"
      value: "1"
backend:
  fileSystem:
    baseDir: "./tmp"
  # s3:
  #   bucketName: "grender.io"
  #   region: "ap-southeast-2"
invalidate:
  amqp:
    uri: "amqp://user:password@localhost:5672/"
    timeout: 10
    queue: "invalidate" 

```

## Running the Server

To run the server, use the following command:

```bash
go run main.go
```

Make sure to have your config.yaml file in the same directory where you run this command.

## Project Structure

The main components of the project are organized into packages:

- **pilot:** Contains the main Grender application files, including configuration, rendering logic, and server setup.
- **backend:** Provides interfaces and implementations for different backends, such as S3 and file systems.
- **render:** Implements the rendering logic using a headless Chrome browser.

## Dependencies

Grender relies on third-party packages for AWS S3 integration, YAML parsing, and headless Chrome browser automation.
