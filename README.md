# robone

## Development
```
git clone git@github.com:naponmeka/robone.git
cd robone
go mod init

go mod download && go get -u -v github.com/therecipe/qt/cmd/qtsetup && go get -u -v github.com/therecipe/qt/cmd/...

go mod vendor && cd ./vendor/github.com/therecipe && git clone https://github.com/therecipe/env_$(go env GOOS)_amd64_512.git && cd ../../..

go mod tidy

qtsetup // test fail is ok

qtdeploy test desktop .

// next time just:
qtdeploy -fast test desktop .

```