# encd

A CLI for file encryption & decryption

## Installation via [Go](https://go.dev/dl/)

Install on any platform using [go install](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies):

```
$ go install github.com/wcygan/encd@latest
```

Disclaimer: the binary will be located at `$GOPATH/bin/encd`. If you want to call the program simply by its name, `encd`, then you need to have $GOPATH/bin on your path:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

## Example

I've provided two files that you can use as an example, [Grumpy Cat](resources/grumpy.jpg)
and [Gary the Snail](resources/gary.png).

We will encode Gary into a temporary file & decode the temporary file back into Gary.

```bash
$ encd enc gary.png -o applesauce -p abcdefghijklmnopqrstuvxyz
```

We can see that `applesauce` is a bunch of nonsense binary:
```bash
$ head applesauce -n 1
==> applesauce <==
vuYu�n�Y�C���Ku_nD��'ހ߼��Q)�Q��S��g7ܽV���<Y����d�!�{���!i��0��ا�Վ[�a�1O�SȆ#��{�<���T$�KS?҅��$��������� �^ŜQ�/d���zs�r�d�Ri}
q�E�&��y��O6�n5\�"p�i�<��9�BH��w�rKq;�ͨ\@���,�L���;�B�����
zv�Uռ����ٲ�S��*N� 9�����XП��O:��&ɸ29�.nV��//����U�'x;���x
```

Unfortunately (or, fortunately?) you will not be able to open this file and see Gary because it is encoded! Let's try to decode it now...

We perform the reverse operation to decode the binary back into the original file

```bash
$ encd dec applesauce -o applesauce.png -p abcdefghijklmnopqrstuvxyz
```

Where is Gary!?!

<img src="resources/gary.png" width="300" height="250">

Oh, there he is...

## Usage

```bash
A tool to encrypt and decrypt files with passwords

Usage:
  encd [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  dec         Decrypt a file that is provided as an argument.
  enc         Encrypt a file that is provided as an argument.
  help        Help about any command

Flags:
  -h, --help              help for encd
  -o, --out string        The file to write to
  -p, --password string   The password used to encode the file
  -t, --toggle            Help message for toggle

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
│   ├── decoder.go
│   ├── encoder.go
│   └── oracle.go
├── go.mod
├── go.sum
├── main.go
├── readme.md
└── resources
    ├── gary.png
    └── grumpy.jpg

```