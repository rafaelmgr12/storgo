<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a id="readme-top"></a>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/rafaelmgr12/storgo">
    <img src="docs/images/logo.png" alt="Logo">
  </a>

  <h3 align="center">Storgo</h3>

  <p align="center">
    A fully distributed content-addressable file storage system built in Go.
    <br />
    <a href="https://github.com/rafaelmgr12/storgo/docs"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    ·
    <a href="https://github.com/rafaelmgr12/storgo/issues">Report Bug</a>
    ·
    <a href="https://github.com/rafaelmgr12/storgo/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

Storgo is a fully distributed content-addressable file storage system built with Go (Golang). It enables decentralized file storage and retrieval through peer-to-peer (P2P) communication, utilizing content-based addressing to ensure data integrity, redundancy, and efficient storage management. This system is designed to store files across distributed nodes, ensuring high availability and secure data transfers between peers.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

[![Go][Go-shield]][Go-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may set up the project locally.

### Prerequisites
* Go (Golang)
```sh
go install
```

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/rafaelmgr12/storgo.git
   cd storgo
   ```
2. Install dependencies
   ```sh
   go mod download
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE -->
## Usage

### Running the Project
To run the project, execute the main file:
```bash
go run main.go
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->
## Roadmap

### 1. Configuration Management
- [ ] **Implement dynamic configuration management**
  - Use `viper` or `envconfig`.
  - Load configurations from YAML, TOML, environment variables, or CLI flags.
  - Support different configurations for dev, staging, and production environments.
  - 
### 2. API and Communication
- [ ] **Add REST API or gRPC endpoints**
  - Use `net/http` for REST or gRPC for binary protocol.
  - Define API routes for external services to interact with the system.

### 3. CLI Improvements
- [ ] **Enhance the CLI interface**
  - Extend CLI functionality using `cobra` to support more commands and flags.
  - Add commands for configuration management, starting/stopping services, etc.

### 4. Docker and Deployment
- [ ] **Dockerize the application**
  - Create a `Dockerfile` for building the application container.
  - Set up Docker Compose for running services in development.
  - Ensure compatibility with deployment environments.


### 5. Service Discovery using Distributed Hash Table (DHT)
- [ ] **Implement Service Discovery with DHT**
  - Use a Distributed Hash Table (DHT) for peer-to-peer service discovery.
  - Research and implement an existing DHT library (e.g., `Kademlia` or `libp2p`).
  - Set up a mechanism where nodes in the system can discover and communicate with each other via the DHT without centralized coordination.


<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
[Go-shield]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/
