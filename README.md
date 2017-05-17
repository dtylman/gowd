# gowd

Build cross platform GUI apps with GO and HTML/JS/CSS (powered by [nwjs](https://nwjs.io/))

Build [![CircleCI](https://circleci.com/gh/dtylman/gowd.svg?style=svg)](https://circleci.com/gh/dtylman/gowd)

### How to use this library:

1. Download and install [nwjs](https://nwjs.io/)
1. Clone this repo.
1. Place `package.json`, `index.html`, `main.go` and `main.js` from [template](cmd/template/) in a new folder. 
1. `go build`
1. Edit `main.js` and set `goBinary` to your your executable name:
    ```javascript
    var goBinary = "./template"; //or template.exe
    ```
1. Run `nw .`, the hello world template should appear:
![hello-world](cmd/template/hello-world.png)
