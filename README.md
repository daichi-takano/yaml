# libcontainer.go
libcontainerとはGo言語で書かれているlibcontainerというライブラリである。 libcontainerを利用してコンテナを作成する。

### Dependency
使用言語は Golang　で　ver 1.15.2

### Setup
コンテナ内で新しいOSを起動するために必要なファイル一式を用意する必要があるため作成する。
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
### Usage
以下のコマンドで実行が可能

```
$ go run libcontainer.go 
 
$ sudo ./libcontainer 
[sudo] ○○○○のパスワード: [sudoパスワードを入力] 
/bin/sh: can't access tty; job control turned off
/ # /bin/hostname 
testing

```

### Authors
libcontainerの作者   Open Container Initiative 

https://github.com/opencontainers/runc/tree/master/libcontainer

### References
実装などの参考文献

Go言語とコンテナ　渋川よしき

https://ascii.jp/elem/000/001/502/1502967/

# main.go
mian.goはYAMLファイルを利用したDockerfileの改善プログラムである。
またbuildkitを使用しDockerよりも素早いコンテナ作成を行う。

### Dependency
使用言語は GOlang　で　ver 1.15.2

### Setup

YAMLファイルを使用するのでパッケージを取得してくる
```
go get go-yaml/yaml
go get gopkg.in/yaml.v2
```

次にbuildkitを有効にするために以下のコマンドを実行する
```
$ export DOCKER_BUILDKIT=1
$ docker build .
```



### Usage
sample1.yamlの名前でDockerfileを作成する。
YAMLファイルの書き換え方法は以下の通りである。

|  Dockerfile  |  YAMLファイル  |
| ---- | ---- |
|  FROM  |  from  |
|  EXPOSE  |  exposes  |
|  RUN  |  runs  |
|  CMD  |  commands  |
|  LABEL  |  labels  |
|  ENV  |  env  |
|  ADD  |  add  |
|  COPY  |  copy  |
|  MAINTAINER  |  maintainer  |
|  ENTRYPOINT  |  entrypoint  |
|  VOLUME  |  volume  |
|  user  |  user  |
|  WORKDIR  |  workdir  |


以下のコマンドで実行が可能である。
```
$ go run main.go 

```

### Authors
作者　東京工科大学　コンピュータサイエンス学部　高野大地

### References
golangでYAMLファイルを読み込む

https://ota42y.com/blog/2014/11/13/go-yaml/
