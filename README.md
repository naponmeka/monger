<div align="center">
  <a href="https://github.com/naponmeka/monger">
    <img width="200" height="200" src="https://github.com/naponmeka/monger/raw/master/res/app-icon/icon_256x256.png">
  </a>
  <br>
</div>

# Monger
Monger is a simple cross-platform MongoDB management tool. Its main purpose is to provide a better user experience for mongodb users.

## Features
- Switching between MongoDB extend json syntax and MongoDB shell syntax
- Export as CSV
- Connect via mongodb uri

## Disclaimer
Still in alpha stage maybe unstable. Do not use on your production database.
Currently under heavy development/testing.

## Download
- Mac OS: [Monger1.0-MacOS-alpha.dmg](https://github.com/naponmeka/monger/raw/master/releases/Monger1.0.0-MacOS-alpha.dmg)
- Windows: coming soon
- Ubuntu: coming soon

## Todo
- other non crud operations : .explain(), .currentOp() , etc.
- More export/import options (Integrate mongodump mongoexport)
- Connect with SSH and other authentication methods
- Example code generator

## Development
```
git clone https://github.com/naponmeka/monger.git
cd monger

go mod download && go get -u -v github.com/therecipe/qt/cmd/qtsetup && go get -u -v github.com/therecipe/qt/cmd/...

go mod vendor && cd ./vendor/github.com/therecipe && git clone https://github.com/therecipe/env_$(go env GOOS)_amd64_512.git && cd ../../..

go mod tidy

qtsetup // test fail is ok

qtdeploy test desktop .

// next time just:
qtdeploy -fast test desktop .

```

## Acknowledgements
- [Qt](https://www.qt.io/)
- [therecipe/qt](https://github.com/therecipe/qt)
- Logo by [Chanisara Suksomjit](https://github.com/csuksomjit)