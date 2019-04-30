<div align="center">
  <a href="https://github.com/naponmeka/monger">
    <img width="200" height="200" src="https://github.com/naponmeka/monger/raw/master/res/app-icon/icon_256x256.png">
  </a>
  <br>

  <h1>Monger</h1>
  <p>
    Monger is ...
  </p>
</div>

# Monger

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