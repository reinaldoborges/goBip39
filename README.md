# My Golang Project

This is a simple Golang project that demonstrates the structure and organization of a Go application.

This program receives a word from [BIP-0039](https://github.com/bitcoin/bips/blob/master/bip-0039/bip-0039-wordlists.md) and return the number (minus 1, to keep it in the range 0-2047).

Or receive an hexadecimal number and return the correspondent word (again, in the range 0-2047).

All the code was created by Copilot plugin inside VSCode.

There is a build script to generate the executable in the folder `dist`

## Project Structure

```
my-golang-project
├── cmd
│   └── main.go        # Entry point of the application
├── pkg
│   └── utils.go       # Utility functions
├── go.mod             # Module dependencies
├── go.sum             # Module checksums
└── README.md          # Project documentation
```

## Getting Started

To get started with this project, follow these steps:

1. Clone the repository:
   ```
   git clone <repository-url>
   ```

2. Navigate to the project directory:
   ```
   cd my-golang-project
   ```

3. Install the dependencies:
   ```
   go mod tidy
   ```

4. Run the application:
   ```
   go run cmd/main.go
   ```

## Utilities

The `pkg/utils.go` file contains utility functions that can be used throughout the project, such as:

- `StringToInt`: Converts a string to an integer.
- `IntToString`: Converts an integer to a string.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.