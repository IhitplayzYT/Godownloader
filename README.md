# Godownloader

A simple and efficient command-line downloader utility written in Go that supports downloading files from URLs and data URIs.

## Why Godownloader?

Godownloader was created to provide a lightweight, no-dependency solution for downloading files from the command line. It handles:

- **HTTP/HTTPS downloads** - Download files from any HTTP or HTTPS URL
- **Data URI support** - Decode and save base64-encoded data URIs (useful for inline images)
- **Automatic filename detection** - Extracts filenames from Content-Disposition headers or URL paths
- **Batch downloads** - Download multiple files in a single command
- **Directory creation** - Automatically creates target directories if they don't exist

## Dependencies

Godownloader has **zero external dependencies**. It uses only the Go standard library:

- `crypto/sha256` - For hashing data URIs
- `encoding/base64` - For decoding base64 data
- `encoding/hex` - For hex encoding
- `net/http` - For HTTP requests
- `net/url` - For URL parsing
- `os` & `io` - For file operations
- `path/filepath` - For path manipulation
- `mime` - For parsing Content-Disposition headers

## Installation

### From Source

```bash
# Clone or download the source
cd Godownloader

# Build the binary
go build -o godownloader

# (Optional) Install to system path
sudo mv godownloader /usr/local/bin/
```

### Using Go Install

```bash
go install github.com/yourusername/Godownloader@latest
```

## How to Use

### Basic Syntax

```bash
./godownloader [URL DESTINATION]...
```

- **URL**: The source URL to download from
- **DESTINATION**: The directory where the file should be saved
- You can specify multiple URL-DESTINATION pairs

### Examples

#### Download a single file to a specific directory

```bash
./godownloader https://example.com/image.png ./downloads
```

#### Download multiple files

```bash
./godownloader \
  https://example.com/file1.pdf ./documents \
  https://example.com/file2.zip ./archives \
  https://example.com/image.jpg ./images
```

#### Download a data URI (base64-encoded image)

```bash
./godownloader "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg==" ./images
```

The file will be saved with a SHA-256 hash as the filename (e.g., `a1b2c3d4...png`).

#### Download without specifying destination (uses current directory)

```bash
./godownloader https://example.com/file.txt .
```

## Features

### Automatic Filename Detection

Godownloader intelligently determines the filename:

1. **Content-Disposition header** - If the server provides a filename via the `Content-Disposition` header (including RFC 5987 encoding), it will be used
2. **URL path** - Falls back to extracting the filename from the URL path
3. **Error handling** - Returns an error if no filename can be determined

### Data URI Types

Supported data URI MIME types:

- `image/png` → `.png` extension
- `image/jpeg` → `.jpg` extension
- `image/svg+xml` → `.svg` extension
- Other types → `.bin` extension

### Error Handling

The tool provides clear error messages for:

- Invalid URLs
- HTTP errors (non-200 status codes)
- File system permission issues
- Network failures
- Unparseable filenames

## Building

```bash
# Build for current platform
go build -o godownloader

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o godownloader-linux

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o godownloader-macos

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o godownloader.exe
```

## Project Structure

```
Godownloader/
├── main.go       # Entry point and argument parsing
├── helper.go     # Download logic and utility functions
└── README.md     # This file
```

## License

This project is provided as-is for educational and practical use.

## Contributing

Feel free to submit issues or pull requests for improvements.
