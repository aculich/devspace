[![DevSpace Logo](docs/website/static/img/github-readme-header.svg)](https://devspace.cloud/)
---

[Website](https://devspace.cloud/) • 
[Documentation](https://devspace.cloud/docs) • 
[Slack](https://devspace.cloud/slack)

[![Build Status](https://travis-ci.org/devspace-cloud/devspace.svg?branch=master)](https://travis-ci.org/devspace-cloud/devspace)
[![Go Report Card](https://goreportcard.com/badge/github.com/devspace-cloud/devspace)](https://goreportcard.com/report/github.com/devspace-cloud/devspace)
[![Slack](https://devspace.cloud/slack/badge.svg)](http://devspace.cloud/slack)
[![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/home?status=Just%20found%20out%20about%20%23DevSpace%20CLI%3A%20https%3A//github.com/devspace-cloud/devspace%0A%0AIt%20lets%20you%20build%20cloud%20native%20software%20directly%20on%20top%20of%20%23Kubernetes%20and%20%23Docker%0A%23CloudNative%20%23k8s)


**How many times do you copy/paste the names of deployments, pods, services etc. per hour when using kubectl?**  
DevSpace CLI eliminates these cumbersome, repetitive tasks through automating and streamlining certain workflows. The goal of this project is to accelerate developing, deploying and debugging applications with Docker and Kubernetes.

## Which workflows can you automate with DevSpace CLI?

<details>
<summary><b>Containerize</b> any project in minutes</summary>

### Containerize your project
```
devspace containerize
```

DevSpace can create a Dockerfile for your project based on the detected programming language in your project.

### Initialize your project
```
devspace init
```

Initializes DevSpace and adds a [highly customizable Helm chart](https://devspace.cloud/docs/charts/devspace-helm-chart) to your project.

Customize your Kubernetes deployment easily:
- [Easily add a database](https://devspace.cloud/docs/chart/customization/predefined-components)
- [Easily add custom kubernetes yamls](https://devspace.cloud/docs/chart/customization/custom-manifests)
- [Configure persistent volumes](https://devspace.cloud/docs/charts/persistent-volumes)
- [Configure environment variables](https://devspace.cloud/docs/charts/environment-variables)

---

</details>

<details>
<summary><b>Deploy</b> containerized applications with ease</summary>

### Deploy your application
```
devspace deploy
```

#### What does `devspace deploy` do?
1. Builds, tags and pushes one or even multiple Docker images
2. Creates pull secrets for your image registries
3. Deploys your project with the newest images (e.g. using Helm)

> DevSpace CLI will use the current kubectl context. If you do not have a Kubernetes cluster, you can use [DevSpace Cloud](https://devspace.cloud) to get a fully managed Kubernetes namespace.

---

</details>

<details>
<summary><b>Develop</b> cloud-native software faster then ever</summary>

### Develop in a production-like environment
```
devspace dev
```
**With DevSpace, you can build and test your application directly inside Kubernetes.** Thanks to our real-time code sync, you can even use hot reloading tools (e.g. nodemon) to refresh your running application without having to waste time on re-building and re-deploying your application every time you change your code. With DevSpace, your containers are updated in real-time without any delay. It works in any container with and without volumes.

Learn more about development with DevSpace:
- [Real-time code synchronization for hot reloading](https://devspace.cloud/docs/cli/development/synchronization)
- [Automatic port forwarding for access via localhost](https://devspace.cloud/docs/cli/development/port-forwarding)
- [Terminal proxy for running commands in your containers](https://devspace.cloud/docs/cli/development/terminal)

---

</details>

<details>
<summary><b>Debug</b> deployments without hassle</summary>

### Speed up finding and solving issues
```
devspace analyze
```
**DevSpace automatically analyzes your deployments**, identifies potential issues and helps you resolve them:
- Identify reasons for image pull failure
- View log snapshots of crashed containers
- Debug networking issues (e.g. misconfigured services)

Learn more about development with DevSpace:
- [Automate issue detection with DevSpace](https://devspace.cloud/docs/cli/debugging/analyze)
- [Stream container logs with DevSpace](https://devspace.cloud/docs/cli/debugging/logs)
- [Use the debugger of your IDE with DevSpace](https://devspace.cloud/docs/cli/debugging/debuggers)
- [Start terminal sessions for debugging](https://devspace.cloud/docs/cli/debugging/enter)

</details>

<br>

## Getting started with DevSpace
### 1. Install DevSpace & Docker

<details>
<summary><b>via Windows Powershell</b></summary>

```
md -Force "$Env:APPDATA\devspace"; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.SecurityProtocolType]'Tls,Tls11,Tls12';
wget -UseBasicParsing ((Invoke-WebRequest -URI "https://github.com/devspace-cloud/devspace/releases/latest" -UseBasicParsing).Content -replace "(?ms).*`"([^`"]*devspace-windows-amd64.exe)`".*","https://github.com/`$1") -o $Env:APPDATA\devspace\devspace.exe; & "$Env:APPDATA\devspace\devspace.exe" "install"; $env:Path = (Get-ItemProperty -Path HKCU:\Environment -Name Path).Path
```

</details>

<details>
<summary><b>via Mac Terminal</b></summary>

```
curl -s -L "https://github.com/devspace-cloud/devspace/releases/latest" | sed -nE 's!.*"([^"]*devspace-darwin-amd64)".*!https://github.com\1!p' | xargs -n 1 curl -L -o devspace && chmod +x devspace;
sudo mv devspace /usr/local/bin;
```

</details>

<details>
<summary><b>via Linux Bash</b></summary>

```
curl -s -L "https://github.com/devspace-cloud/devspace/releases/latest" | sed -nE 's!.*"([^"]*devspace-linux-amd64)".*!https://github.com\1!p' | xargs -n 1 curl -L -o devspace && chmod +x devspace;
sudo mv devspace /usr/local/bin;
```

</details>

Install Docker if you don't have already
<details>
    <summary><b>Install Docker</b></summary>

DevSpace uses Docker to build container images, so you need Docker on your local computer. If you do not have Docker installed yet, you can download the latest stable releases here:
- **Mac**: [Docker Community Edition](https://download.docker.com/mac/stable/Docker.dmg)
- **Windows Pro**: [Docker Community Edition](https://download.docker.com/win/stable/Docker%20for%20Windows%20Installer.exe)
- **Windows 10 Home**: [Docker Toolbox](https://download.docker.com/win/stable/DockerToolbox.exe) (legacy)

</details>

### 2. Containerize your application
If you don't have a Dockerfile run this command within your project:
```
devspace containerize
```

Then initialize DevSpace with this command in your project:
```
devspace init
```
<details>
<summary><b>Don't have a project to test DevSpace with?</b> Check out our example project.</summary>

```
git clone https://github.com/devspace-cloud/devspace-quickstart-nodejs
```

</details>

<br>

**What does `devspace containerize & devspace init` do?**  
DevSpace CLI will automatically detect your programming language and ask for the port your application is listening on. It will then create an Helm chart and a Dockerfile within your project, if you do not already have one.

### 3. Create a space (optional)

> This step is optional, you can also use your custom kubernetes cluster with DevSpace

This command will create and configure a Kubernetes namespace for you in the devspace cloud:
```
devspace create space my-app
```
DevSpace Cloud will provide a fully managed Kubernetes namespace for you. You can create one Space, which is forever free on [DevSpace Cloud](https://devspace.cloud). 

### 4. Deploy your application
Deploy your application to kubernetes:
```
devspace deploy
```

### What's next?
- [Developing applications with DevSpace](https://devspace.cloud/docs/getting-started/development)
- [Debugging deployments with DevSpace](https://devspace.cloud/docs/cli/debugging/overview)
- [Customize the DevSpace Helm Chart](https://devspace.cloud/docs/chart/basics/devspace-helm-chart)

<br>

## Contributing
Help us make DevSpace CLI the best tool for developing, deploying and debugging Kubernetes apps.

### Reporting Issues
If you find a bug while working with the DevSpace CLI, please [open an issue on GitHub](https://github.com/devspace-cloud/devspace/issues/new?labels=kind%2Fbug&template=bug-report.md&title=Bug:) and let us know what went wrong. We will try to fix it as quickly as we can.

### Feedback & Feature Requests
You are more than welcome to open issues in this project to:
- [give feedback](https://github.com/devspace-cloud/devspace/issues/new?labels=kind%2Ffeedback&title=Feedback:)
- [suggest new features](https://github.com/devspace-cloud/devspace/issues/new?labels=kind%2Ffeature&template=feature-request.md&title=Feature%20Request:)
- [ask a question on Slack](https://devspace.cloud/slack)

### Contributing Code
This project is mainly written in Golang. If you want to contribute code:
1. Ensure you are running golang version 1.11.4 or greater for go module support
2. Set the following environment variables:
    ```
    GO111MODULE=on
    GOFLAGS=-mod=vendor
    ```
3. Check-out the project: `git clone https://github.com/devspace-cloud/devspace && cd devspace`
4. Run `go clean -modcache`
5. Run `go mod vendor`
6. Make changes to the code
7. Build the project, e.g. via `go build -o devspace.exe`
8. Evaluate and test your changes `./devspace [SOME_COMMAND]`

See [Contributing Guideslines](CONTRIBUTING.md) for more information.


<br>

## FAQ
<details>
<summary>Do I need a Kubernetes cluster to use DevSpace?</summary>

**No.** You can simply use **the fully managed Spaces** provided by DevSpace Cloud.

</details>

<details>
<summary>Can I use DevSpace with my existing Kubernetes clusters?</summary>

**Yes.** You have two options:
1. [Connect your existing Kubernetes clusters to DevSpace Cloud](https://devspace.cloud/docs/cloud/external-clusters/overview) as external clusters. DevSpace Cloud will then be able to create and manage users and Spaces on top of your Kubernetes clusters.
2. You just use DevSpace CLI without DevSpace Cloud. That means that you manually need to:
    * enforce resource limits
    * configure secure user permissions
    * isolate namespaces of different users
    * connect domains and configure ingresses
    * install and manage basic cluster services (e.g. ingress controller, cert-manager for TLS, monitoring and log aggregation tools)

</details>

<details>
<summary>Do I need to be a Kubernetes expert to use DevSpace?</summary>

**No.** Altough DevSpace provides a lot of advanced tooling for Kubernetes experts, it is optimized for developer experience which makes it especially easy to use for Kubernetes beginners.

</details>

<details>
<summary>What is a Space?</summary>

Spaces are smart Kubernetes namespaces which provide the following features:
- Automatic provisioning via `devspace create space [SPACE_NAME]`
- Automatic allocation of a subdomain for each Space, e.g. `my-app.devspace.host`
- Automatic RBAC configuration for better isolation of users
- Automatic resource limit configuration and enforcement
- Resource auto-scaling within the configured limits
- Smart analysis of issues within your Space via `devspace analyze`

</details>

<details>
<summary>What is DevSpace CLI?</summary>

DevSpace CLI is an open-source command-line tool that provides everything you need to develop, deploy and debug applications with Docker and Kubernetes.

> You can either use DevSpace CLI as standalone solution for your self-managed Kubernetes namespaces or in combination with DevSpace Cloud.

</details>

<details>
<summary>What is DevSpace Cloud?</summary>

DevSpace Cloud is a developer platform for Kubernetes that lets you create and manage Spaces via DevSpace CLI or GUI. 

> The Spaces you create with DevSpace Cloud either run on a Kubernetes cluster within DevSpace Cloud or on your own Kubernetes clusters after connecting them to the platform.

</details>

<details>
<summary>What is a Helm chart?</summary>

[Helm](https://helm.sh/) is the package manager for Kubernetes. Packages in Helm are called Helm charts.

[Learn more about Helm charts](https://helm.sh/docs/)

</details>

<br>

## License
You can use the DevSpace CLI for any private or commercial projects because it is licensed under the Apache 2.0 open source license.
