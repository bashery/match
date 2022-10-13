<br><br>

<h1 align="center">Match</h1>

<p align="center">
  <a href="/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg"/></a>
  <a href="https://godoc.org/github.com/bashery/match"><img src="https://godoc.org/github.com/mingrammer/commonregex?status.svg"/></a>
  <a href="https://goreportcard.com/report/github.com/bashery/match"><img src="https://goreportcard.com/badge/github.com/mingrammer/commonregex"/></a>
</p>

<p align="center">
Matching the most common expressions
</p>

<br><br><br>


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
* Phone with exts
* Link
* Email
* IPv4
* IPv6
* IP
* Port without well-known (not known port)
* Pric
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


## Tanks to :heart:
MinJae Kwon
* [@mingrammer](https://github.com/mingrammer)
<br>
# NOTE:
This library is inspired by [CommonRegex](https://github.com/madisonmay/CommonRegex), and a lot of code has been copied and modified from it
<br>
## License

MIT
