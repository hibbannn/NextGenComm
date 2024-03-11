
# NextGenComm: gRPC & HTTP Communication Boilerplate

Welcome to **NextGenComm**, the advanced boilerplate for building efficient and flexible communication applications using gRPC and HTTP protocols. Designed with a modular and scalable architecture, this boilerplate provides a solid foundation for developing microservices that can be easily integrated into various environments and use cases.

## Features

- **Adaptable Architecture**: Includes adapters for caching, clients, mail services, and repositories, enabling seamless adaptation to your specific application needs.
- **Well-Defined Contracts**: Ensures consistent interfaces for service communication, facilitating easy development and testing.
- **Support for gRPC and HTTP Protocols**: Offers flexibility in communication protocol choice, ensuring efficient communication in diverse environments.
- **Organized Structure**: Enhances code maintainability and navigability with a clear separation of functions and roles.
- **Automation and Documentation Ready**: Comes with automation scripts, API documentation, and customizable configuration settings to accelerate the development process.

## Getting Started

### Prerequisites

- Go 1.16 or later
- Protobuf compiler (protoc)
- Docker (for containerization and database services)

### Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/yourusername/NextGenComm.git
   ```

2. **Install Dependencies**

   Navigate to the project directory and install the required Go dependencies.

   ```bash
   cd NextGenComm
   go mod tidy
   ```

3. **Generate Protobuf Stubs**

   Utilize the provided script to generate gRPC and Protoc Gateway stubs.

   ```bash
   ./script/generate-stub.sh
   ```

4. **Configure Environment**

   Adjust the configuration files in `doc/config` according to your environment (development, staging, production).

### Running the Application

- **Development Mode**

  ```bash
  go run cmd/main.go
  ```

- **Production Mode**

  Ensure you have configured your production environment correctly, then build and run the application.

  ```bash
  go build -o NextGenComm cmd/main.go
  ./NextGenComm
  ```

## Documentation

- **API Documentation**: Located in `doc/api/index.html` for detailed endpoint descriptions.
- **Code Structure**: Refer to the project's `tree` structure for an overview of the organization.

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Your Name - [@hibbann](https://instagram.com/hibban) - primary.hibban.nurcholis@gmail.com

Project Link: [https://github.com/hibbannn/NextGenComm](https://github.com/hibbannn/NextGenComm)

## Acknowledgements

- [Go](https://golang.org/)
- [gRPC](https://grpc.io/)
- [Protobuf](https://developers.google.com/protocol-buffers)
