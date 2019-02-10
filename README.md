# tap 
tap (short for trap and pass in soccer) is a simple redirect server.

## Getting Started

### Usage with Heroku
1. Set config.json
1. Deply to Heroku
1. Add wildcard domain in settings. DNS Target will be displayed. 
    * ex) `*.hypermkt.jp`
1. Specify the DNS Target to your DNS provider for the destination of CNAME.

### Directory Structure

```sh
.
├── assets
│   ├── css
│   ├── images
│   └── js
├── config.json
├── main.go
└── templates
    ├── 404.html
    └── redirect.html
```

#### config.json

You can set multiple redirect source and destination URL in config.json file.

```json
{
  "redirects": [
    {"from": "localhost", "to": "https://www.bing.com/"},
    {"from": "redirect-example01.hypermkt.jp", "to": "https://www.yahoo.co.jp/"},
    {"from": "redirect-example02.hypermkt.jp", "to": "https://www.google.co.jp/"}
  ]
}
```

#### assets

Static files are servered from assets directory. You can create any directories and specify `/assets/your-new-dir` in html file.

#### templates

file name|use
redirect.html|Displayed before beeing redirected.
404.html|Displayed when specified URL didn't match with config.json

## author
* hypermkt
