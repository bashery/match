<br><br>

<h1 align="center">Match</h1>

<p align="center">
  <a href="/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg"/></a>
  <a href="https://app.fossa.io/projects/git%2Bgithub.com%2Fmingrammer%2Fcommonregex?ref=badge_shield" alt="FOSSA Status"><img src="https://app.fossa.io/api/projects/git%2Bgithub.com%2Fmingrammer%2Fcommonregex.svg?type=shield"/></a>
  <a href="https://godoc.org/github.com/bashery/match"><img src="https://godoc.org/github.com/mingrammer/commonregex?status.svg"/></a>
  <a href="https://goreportcard.com/report/github.com/bashery/match"><img src="https://goreportcard.com/badge/github.com/mingrammer/commonregex"/></a>
</p>

<p align="center">
Matching the most common expressionskkk
</p>

<br><br><br>

> Inspired by [CommonRegex](https://github.com/madisonmay/CommonRegex)

This is a collection of often used regular expressions. It provides these as simple functions for getting the matched strings corresponding to specific patterns.

## Installation
```shell
go get github.com/bashery/match
```

## Usage

```go
import (
    "github.com/bashery/match"
)

func main() {

    match.IsDate("Jan 9th 2012") // true

    match.IsTime("5:00PM") // true

    match.IsLink("www.linkedin.com") // true

    match.IsEmails("adams@email.com") // true
}
```

## Features

* Date
* Time
* Phone
* Phones with exts
* Link
* Email
* IPv4
* IPv6
* IP
* Ports without well-known (not known ports)
* Price
* Hex color
* Credit card
* VISA credit card
* MC credit card
* ISBN 10/13
* BTC address
* Street address
* Zip code
* Po box
* SSN
* MD5
* SHA1
* SHA256
* GUID
* MAC address
* IBAN
* Git Repository

## Thanks to :heart:

* [@cschoede](https://github.com/cschoede)
* [@schigh](https://github.com/schigh)
* [@emaraschio](https://github.com/emaraschio)
* [@mamal72](https://github.com/mamal72)
* [@ahmdrz](https://github.com/ahmdrz)
* [@fakenine](https://github.com/fakenine)
* [@Bill-Park](https://github.com/Bill-Park)
* [@jakewarren](https://github.com/jakewarren)

## Spicial hanks to :heart:
** MinJae Kwon 
* [@mingrammer](https://github.com/mingrammer)

## License

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fmingrammer%2Fcommonregex.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fmingrammer%2Fcommonregex?ref=badge_large)
