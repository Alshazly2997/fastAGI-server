# Go FastAGI Server

A high-performance FastAGI (Asterisk Gateway Interface) server built with Go. This service allows Asterisk to offload call logic to a standalone Go application, providing better scalability and flexibility than traditional AGI scripts.

##  Features

* **FastAGI Protocol:** Communicates over TCP (Port 4573).
* **Built with `astgo`:** Utilizes the lightweight `wenerme/astgo` library.
* **Media Handling:** Automatically answers calls and plays back audio files.
* **Extension Dialing:** Capable of bridging calls to other internal or external endpoints.
* **Variable Management:** Read and write Asterisk channel variables directly from Go.

## Prerequisites

* **Go:** 1.18 or higher.
* **Asterisk:** 13.x or higher.
* **Network:** Port `4573` must be open on the server running the Go app.

## Installation

1. **Install dependencies:**
   ```Bash
    go get github.com/wenerme/astgo/agi

## Usage

1. **Run the Server:** Start the Go application. It will listen for incoming connections from Asterisk on port 4573.

    ```Bash
    go run main.go

2. **Configure Asterisk Dialplan:** Add an entry to your extensions.conf to route calls to your Go server:

    ```Ini, TOML
    [default]
    exten => 100,1,AGI(agi://localhost:4573)
 
Note: Replace localhost with the IP address of your Go server if it is running on a different machine.
