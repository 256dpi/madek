# madek

**A Go library and command line tool that simplifies accessing the [Madek API](https://medienarchiv.zhdk.ch/api/browser/index.html).**

## Installation

```
$ go get github.com/IAD-ZHDK/madek/cmd/madek
```
 
## Usage

```
madek.

Usage:
  madek fetch <id> [options]
  madek server [options]

Options:
  -h --help                   Show this screen.
  -a --address=<url>          The address of the madek instance [default: https://medienarchiv.zhdk.ch].
  -u --username=<username>    The username used for authentication.
  -p --password=<password>    The password used for authentication.
```
