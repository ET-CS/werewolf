[![Build Status](https://travis-ci.org/ET-CS/werewolf.svg?branch=master)](https://travis-ci.org/ET-CS/werewolf)

Werewolf
========

REAL FAST static html webserver written in Golang.

## Install

```go get github.com/ET-CS/werewolf```

## Usage

Start the server using the `werewolf` command inside any directory contains html files.
the server will run on port 8585 and will serve your html files.

### Logic

* `index.html` will be always root (`http://example.com/`)
* Other files will be on route as filename (for example `about.html` will be `example.com/about`)
* If there are minified files (`index.min.html`) inside the directory - they will be threated too by the same logic. (`about.min.html` will be `example.com/about`)
* If there are two files - one minified and one is not (index.html and index.min.html for example) Werewold will use the minified one.

All html files are cached into memory for performance on server start.

You can use task runner like [Dry](https://github.com/ET-CS/dry) to easily create minified html(s) website from templates based website.

# Start in background (and on boot)

As almost always, There are lot of options.
I personally use `supervisord` to start werewolf server as daemon and `nginx` as proxy.

by ET-CS (Etay Cohen-Solal)