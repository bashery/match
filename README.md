<br><br>

<h1 align="center">Match</h1>

<p align="center">
  <a href="/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg"/></a>
</p>

<p align="center">Matching the most common expressions</p>

<br><br>

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

## Todo

- [x] Email
- [] Date
- [] Time
- [] Phone
- [] Phone with exts
- [] Link
- [] IPv4
- [] IPv6
- [] IP
- [] Port without well-known (not known port)
- [] Pric
- [] Hex color
- [] Credit card
- [] VISA credit card
- [] MC credit card
- [] ISBN 10/13
- [] BTC address
- [] Street address
- [] Zip code
- [] Po box
- [] SSN
- [] MD5
- [] SHA1
- [] SHA256
- [] GUID
- [] MAC address
- [] IBAN
- [] Git Repository


## License

MIT
