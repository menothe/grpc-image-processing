# What it is

This is a lightweight image processing tool that converts a JPEG image to grayscale.
It leverages gRPC for efficient processing.

![[Example]](example.png)

## How to use

After cloning, run `go run main.go` at the root to boot the server, which is then
ready for any client to submit image requests via gRPC calls. You can use a CLI tool
like grpcurl to send the requests.

Sample json with a base64 encoded value (file's called "image_request.json" for example):

```json
{
    "image_data": "O9vOmyyJvEbVFAKkEdgTtQIEgb"
}
```

Sample grpcurl call to the server: ```grpcurl -plaintext -d "`cat image_request.json`" localhost:8080 ImageProcessor/ConvertImage```

The server will return a new base64 encoded value which you can copy to clipboard
or save in a file (i.e. a .txt file) then decode the returned base64 using:

`base64 --decode -i grayscale.txt -o grayscale.jpg`

Opening grayscale.jpg should reveal a grayscale converted image of the original.

