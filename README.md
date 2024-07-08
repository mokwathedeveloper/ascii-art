

# ASCII Art Generator

This project is an ASCII Art Generator that converts input text into ASCII art using predefined templates. The generator supports validation of template files through SHA-256 checksums and downloading templates if they are missing.

## Table of Contents

1. [Features](#features)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Project Structure](#project-structure)
5. [Contributing](#contributing)

## Features

- Generate ASCII art from text input.
- Validate template files using SHA-256 checksums.
- Automatically download missing template files.
- Error handling for unprintable ASCII characters.

## Installation

To use this ASCII Art Generator, ensure you have Go installed on your system. Then, follow these steps:

1. Clone the repository:
	```sh
	git clone https://learn.zone01kisumu.ke/git/mmoffat/ascii-art.git
	cd ascii-art
	```

## Usage

Run the ASCII Art Generator with the following command:

```sh
./ascii-art <hello>
```

Replace `<input_text>` with the text you want to convert to ASCII art.

## Project Structure

```plaintext
.
├── component
│   └── ascii.go       	# Functions to handle ASCII art generation
├── main
│   ├── checksum.go    	# Functions for file checksum validation
│   ├── download.go    	# Functions for downloading missing template files
│   └── main.go        	# Main entry point of the application
├── templates          	# Directory containing ASCII art templates
├── go.mod             	# Go module file
└── go.sum             	# Go dependencies file
```

### `component/ascii.go`

This file contains functions to handle ASCII art generation. It includes:
- `MakeAsciiGraphicsMap()`: Creates a map from the ASCII art template.
- `GetTheCorrectFile()`: Retrieves the correct template file.
- `Output()`: Generates the ASCII art output from the input string.

### `main/checksum.go`

This file contains functions for file checksum validation. It includes:
- `validateFileChecksum()`: Validates the checksum of a given file.
- `calculateChecksum()`: Calculates the SHA-256 checksum of a file.

### `main/download.go`

This file contains functions for downloading missing template files. It includes:
- `downloadFile()`: Downloads a file from a specified URL.

### `main/main.go`

This is the main entry point of the application. It reads the input text, validates and reads the template file, and generates the ASCII art.

## Contributing

Contributions are welcome! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch: `git checkout -b my-feature-branch`
3. Make your changes and commit them: `git commit -m 'Add new feature'`
4. Push to the branch: `git push origin my-feature-branch`
5. Create a pull request.

