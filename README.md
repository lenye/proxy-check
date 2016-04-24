[![Build Status](https://secure.travis-ci.org/vintikzzz/proxy-check.png?branch=master)](http://travis-ci.org/vintikzzz/proxy-check)

## proxy-check: simple cli and library to check proxy list

### Using cli

Simple usage:
```
proxylist.txt | proxy-check
```
Args:
```
-check string
    Text that expected at target site (default "Aliexpress")
-target string
    Target test url (default "http://aliexpress.com")
-timeout int
    Timeout in seconds (default 5)
-verbose
    Enables verbose mode
-workers int
    Number of workers (default 20)
```

### Credits

hideme is (c) Tatarskiy Pavel, 2016

### License

hideme is distributed under the MIT License, see `LICENSE` file for details.
