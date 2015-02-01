Werewolf
========

REAL FAST static html webserver written in Golang.

Running the `werewolf` command inside any folder contains html files will start a new webserver
on port 8585 to serve your files.

index.html will be root ("/")
other files will be on route as filename (about.html = "/about")

if minified files exist (index.min.html) - they will be loaded too with the same logic.

if there are two files - one minified and one not (index.html and index.min.html for example) Werewold will load the minified one.

All html files will be cached into memory on server start for performance.

You can use task runner like [Dry](https://github.com/ET-CS/dry) to easily create minified files from templates based website.

by ET-CS (Etay Cohen-Solal)