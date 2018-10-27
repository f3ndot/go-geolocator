# go-geolocator

A true micro-service for geolocating IP addresses using Go for speed and MaxMind's GeoIP2 City database.

## Setup

You must first download the geolite databases for local development!

```bash
git clone git@github.com:f3ndot/go-geolocator.git
cd go-geolocator
\curl -SL https://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz | tar xvz --strip-components 1 -C data/ '*/*.mmdb'
\curl -SL https://geolite.maxmind.com/download/geoip/database/GeoLite2-ASN.tar.gz | tar xvz --strip-components 1 -C data/ '*/*.mmdb'
```

## Licenses

### `go-geolocator`

See the [LICENSE](license) file.

### GeoLite2

The GeoLite2 databases are distributed under the [Creative Commons Attribution-ShareAlike 4.0 International License](https://creativecommons.org/licenses/by-sa/4.0/). The attribution requirement may be met by including the following in all advertising and documentation mentioning features of or use of this database:

```html
This product includes GeoLite2 data created by MaxMind, available from
<a href="http://www.maxmind.com">http://www.maxmind.com</a>.
```

### `geoip2-golang`

```
ISC License

Copyright (c) 2015, Gregory J. Oschwald <oschwald@gmail.com>

Permission to use, copy, modify, and/or distribute this software for any
purpose with or without fee is hereby granted, provided that the above
copyright notice and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR
OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
PERFORMANCE OF THIS SOFTWARE.
```