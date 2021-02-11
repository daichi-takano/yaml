# libcontainer.go
libcontainerとはGo言語で書かれているlibcontainerというライブラリです。 libcontainerを利用してコンテナを作成します。

## Dependency
使用言語は Golang　で　ver 1.15.2

## Setup
コンテナ内で新しいOSを起動するために必要なファイル一式を用意する必要があるため作成する
ここではサイズが小さいAlpine Linuxのイメージに含まれているファイルを使用する。

```
$ docker pull alpine

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
以下のコマンド実行が可能

```
$ go run libcontainer.go 
 
$ sudo ./libcontainer 
[sudo] ○○○○のパスワード: [sudoパスワードを入力] 
/bin/sh: can't access tty; job control turned off
/ # /bin/hostname 
testing

```

## Authors
libcontainerの作者   Open Container Initiative 
https://github.com/opencontainers/runc/tree/master/libcontainer

## References
実装などの参考文献

Go言語とコンテナ　渋川よしき
https://ascii.jp/elem/000/001/502/1502967/

# main.go

## Dependency
使用言語は GOlang　で　ver 1.15.2

## Setup


## Usage


## License


## Authors


## References
