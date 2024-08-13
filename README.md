# Dumpon 🗑️ Go HTTP Request Dump Server

**Dumpon** is a simple HTTP server designed to dump information about incoming HTTP requests written in Go. It prints details such as request URL, method, headers, parameters, body and file uploads to the console.

## Notes

- I made this to have a portable dump server with me, usable anywhere at anytime, for debugging requests.
- This is my first project in Go lang so suggestion and pull requests are welcome from experts.

## Installation

You can download the binary from [release](../../releases) page or build on your own from source. Ensure that Go^1.22.6 is installed and properly configured on your system to build from source.

## Usage

```bash
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

## Print Format

Each request is separated by start and end boundary.

```bash
-------------Start-------------
Request URL: localhost:8001/hello-world?name=rahul&email=rahul@email.com
Request Method: POST

Request Headers:
Content-Type: multipart/form-data; boundary=--------------------------622393842521055747236185
Authorization: Bearer 12345678
X-Format: custom
Connection: keep-alive
Accept: application/json, text/plain, */*
User-Agent: axios/1.7.2
Content-Length: 20460
Accept-Encoding: gzip, compress, deflate, br
Request-Start-Time: 1723181706839

URL Parameters:
email: rahul@email.com
name: rahul

Form Fields:
name: rahul
email: rahul@email.com
password: password

Files:
avatar: profile.jpg
--------------End--------------
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
