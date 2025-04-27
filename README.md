# kube-metrics-collector

A simple Go application to collect and expose basic metrics from a Kubernetes cluster.

## Overview

This project provides a basic HTTP server that interacts with the Kubernetes API to retrieve information about nodes and pods. It can be configured to use either a real Kubernetes client (requiring access to a cluster's `kubeconfig`) or a fake client for testing purposes.

## Features

* Lists nodes in the Kubernetes cluster.
* Lists pods in a specified namespace (defaults to `default`).
* Exposes basic cluster metrics (number of nodes, pods, and placeholder metrics) via an HTTP endpoint.
* Supports a fake Kubernetes client for local development and testing without a real cluster.
* Configurable via environment variables.
* Containerized with Docker for easy deployment.

## Getting Started

### Prerequisites

* Go (version 1.24 or later)
* Docker (if running in a container)
* Minikube (optional, for local Kubernetes testing)
* kubectl (optional, for interacting with Kubernetes clusters)

### Local Setup and Running Against Minikube

1.  **Clone the Repository:**
    ```bash
    git clone <repository_url>
    cd kube-metrics-collector
    ```
    (Replace `<repository_url>` with your repository URL)

2.  **Start Minikube (if needed):**
    ```bash
    minikube start
    ```

3.  **Set `kubectl` Context to Minikube:**
    ```bash
    kubectl config use-context minikube
    ```

4.  **Run the Go Application:**
    ```bash
    go run main.go
    ```

5.  **Access Endpoints:**
    * `http://localhost:8080/nodes`: List nodes.
    * `http://localhost:8080/pods`: List pods in the `default` namespace.
    * `http://localhost:8080/pods?namespace=<namespace>`: List pods in a specific namespace.
    * `http://localhost:8080/metrics`: Get basic cluster metrics.
    * `http://localhost:8080/print-metrics`: Print metrics to server logs.

### Running with Docker

1.  **Build the Docker Image:**
    ```bash
    docker build -t kube-metrics-collector .
    ```

2.  **Run with Fake Client:**
    ```bash
    docker run -p 8080:8080 -e USE_FAKE=true kube-metrics-collector
    ```

3.  **Run with Real Client (Mount `kubeconfig`):**
    ```bash
    # Linux/macOS
    docker run -p 8080:8080 -v "$HOME/.kube/config:/root/.kube/config" kube-metrics-collector

    # Windows (PowerShell)
    docker run -p 8080:8080 -v "$env:USERPROFILE\.kube\config:/root/.kube/config" kube-metrics-collector

    # Windows (Command Prompt)
    docker run -p 8080:8080 -v "%USERPROFILE%\.kube\config:/root/.kube/config" kube-metrics-collector
    ```

4.  **Access Endpoints (Docker):**
    Access the same endpoints as in the local setup via `http://localhost:8080` on your host.

### Configuration

Environment variables can be used to configure the application:

* `USE_FAKE`: Set to `true` to use the fake Kubernetes client. Defaults to `false`.
* `CLUSTER_URL`: (Currently not actively used in the code).
* `PORT`: Sets the HTTP server listening port. Defaults to `8080`.
* `KUBECONFIG`: Specifies the path to the `kubeconfig` file inside the container (e.g., `/root/.kube/config`).