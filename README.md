# BVM
Bun Version Manager (BVM) is a tool for managing multiple versions of Bun on a single machine. It is inspired by [nvm](https://github.com/nvm-sh/nvm). 

BVM is written in [Go](https://golang.org/).

## Installation

Head over to the [releases](https://github.com/gaurishhs/bvm/releases) page and download the latest version of BVM for your platform. 

## Usage

### Installing a version of Bun

To install a version of Bun, run the following command:

If no version is specified, the latest version of Bun will be installed.

```bash
bvm install --release <version>
```

### Listing available versions of Bun

To list all available versions of Bun, run the following command:

```bash
bvm list
```

### Get latest version of Bun

To get the latest version of Bun, run the following command:

```bash
bvm latest
```

## License

BVM is licensed under the BSD 3-Clause License. See [LICENSE](https://github.com/gaurishhs/bvm/tree/main/LICENSE) for more information.