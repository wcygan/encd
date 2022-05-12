# encd

An image encoder & decoder

This should NOT make any assumptions about file extensions

It simply takes a blob of bytes & performs some modifications to them

## TODO:
1. write the encrypted stuff to a file
2. decrypt the stuff into another file
3. figure out if you can do it without 32 byte password
4. rename everything about "image" to "file", this is generic
5. Renamed "Encode & Decode" to "Encrypt & Decrypt"

## Installation via [Go](https://go.dev/dl/)

Install on any platform using `go get`:

```
$ go get github.com/wcygan/encd
```

## Usage

```
An image encoder & decoder

Usage:
  encd [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  dec         Decode an image
  enc         Encode an image
  help        Help about any command

Flags:
  -h, --help     help for encd
  -t, --toggle   Help message for toggle

```

## Resources

Maybe we can use the Go [image](https://pkg.go.dev/image) package?

- https://github.com/averagesecurityguy/crypto/blob/master/crypt.go
- https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/
- https://medium.com/swlh/protect-image-file-using-encryption-written-in-go-7d016c5a4719