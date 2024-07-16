# go-passwo

A secure password generator CLI built with Go.

This repository contains a command-line tool for generating secure passwords, built using Go. The tool provides various options to customize the password generation process, including setting the length, omitting special characters, and generating PIN codes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Features

- **Customizable Password Length**: Specify the length of the password to be generated.
- **Omit Special Characters**: Generate passwords without including any special characters.
- **Generate PIN Codes**: Generate numeric PIN codes instead of full passwords.
- **Copy to Clipboard**: Automatically copy the generated password to the clipboard for easy use.

## Installation
```bash
go get -u github.com/qovzeash/go-passwo
```
## Usage

Once installed, you can use the tool from the command line. The following options are available:

```--size```: Specify the length of the password to be generated. The default value is 12.

```--omit-symbols```: Generate a password without including any special characters. The default value is false.

```--pin-code```: Generate a numeric pin code instead of a full password. The default value is false.

## Example

Generate a 16-character password with symbols:
```bash
password-generator --size 16
```

Generate a 12-character password without symbols:
```bash
password-generator --omit-symbols
```

Generate a 6-digit pin code:
```bash
password-generator --size 6 --pin-code
```