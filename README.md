# Dumpon 🗑️ Go HTTP Request Dump Server

**Dumpon** is a simple HTTP server designed to dump information about incoming HTTP requests written in Go. It prints details such as request URL, method, headers, parameters, body and file uploads to the console.

## Requirements

- Ensure that Go ^1.22.6 is installed and properly configured on your system to build and run the server.

## Usage

```bash
# Build the executable
go build -o dumpon main.go

# See available options
./dumpon -h

# Run the server
# Default port `-p 80`
# Max memory `-m 10`
./dumpon
```

## Sending Request

```bash
# Simple GET request
curl http://localhost:80

# Simple POST request with body
curl -X POST http://localhost:80 -d "name=dumpon"

# Simple POST request with file upload
curl -X POST http://localhost:80 -F "avatar=@path/to/avatar.png"
```

## Changelog

Please see [CHANGELOG](CHANGELOG.md) for more information what has changed recently.

## Security

If you discover any security related issues, please email rahulhaque07@gmail.com instead of using the issue tracker.

## Credits

-   [Rahul Haque](https://github.com/rahulhaque)
-   [All Contributors](../../contributors)

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
