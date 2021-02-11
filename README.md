# libcontainer.go
libcontainerとはGo言語で書かれているlibcontainerというライブラリです。 libcontainerを利用してコンテナを作成します。

## Dependency
使用言語は Golang　で　ver 1.15.2

## Setup
コンテナ内で新しいOSを起動するために必要なファイル一式を用意する必要があるため作成する
ここではサイズが小さいAlpine Linuxのイメージに含まれているファイルを使用する。

```
$docker pull alpine

$ docker run --name alpine alpine 
 
$ docker export alpine > alpine.tar 
 
$ docker rm alpine 
```

ライブラリを以下のコマンドで取得してくる
```
$go get github.com/opencontainers/runc/libcontainer ⏎
$ go get golang.org/x/sys/unix
```
## Usage


## License


## Authors


## References

# main.go

## Dependency
使用言語は GOlang　で　ver 1.15.2

## Setup


## Usage


## License


## Authors


## References
