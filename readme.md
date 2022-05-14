# encd

A CLI for file encryption & decryption.

## Installation via [Go](https://go.dev/dl/)

Install on any platform using [go install](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies):

```
$ go install github.com/wcygan/encd@latest
```

Hint: make sure
you've [set up your Go workspace](https://www.digitalocean.com/community/tutorials/understanding-the-gopath#anatomy-of-the-go-workspace)
to access `$GOPATH/bin`.

## Example

### Encryption

Encrypt a list of files or directories:

```bash
$ encd enc [ARGS...] -p YourSecretPhrase
```

### Decryption

Decrypt a list of files or directories:

```bash
$ encd dec [ARGS...] -p YourSecretPhrase
```

## Usage

```bash
$ encd
A tool to encrypt and decrypt files and directories with passwords

Usage:
  encd [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  dec         Decrypt a list of files or directories
  enc         Encrypt a list of files or directories
  help        Help about any command

Flags:
  -h, --help              help for encd
  -p, --password string   The secret phrase used for encryption and decryption
  -t, --toggle            Help message for toggle

Use "encd [command] --help" for more information about a command.
```

## Cross Compilation

Want to compile this program for supported platforms? (You can find supported platforms via `go tool dist list`). 

To cross-compile this project, use the following command:

```bash
$ chmod +x cross-compile.bash && ./cross-compile.bash
```

## File Tree

```bash
$ tree
.
├── LICENSE
├── cmd
│   ├── dec.go
│   ├── enc.go
│   └── root.go
├── crypto
│   ├── decrypt.go
│   ├── encrypt.go
│   └── oracle.go
├── go.mod
├── go.sum
├── main.go
├── readme.md
└── examples
    ├── gary.png
    └── grumpy.jpeg
```

## Disclaimer

This tool is powerful, use it at your own risk. If you aren't clever about how you encrypt, decrypt, or backup your data you may end up losing data.