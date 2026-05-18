# Home Server Hands-On

Docker 上でシンプルな HTTP サーバーを立てるサンプルです．

次の 2 種類の実装があります．

* Python + FastAPI
* Go 標準ライブラリ

どちらも `/` にアクセスすると，シンプルな HTML を返します．

## Python + FastAPI

### ディレクトリ構成

```txt
python-app/
├── Dockerfile
├── requirements.txt
└── app
    └── main.py
```

### ビルド

```bash
cd python-app
docker build -t python-server .
```

### 起動

バックグラウンドで起動します．

```bash
docker run -d --name python-server -p 8000:8000 python-server
```

### 確認

```bash
curl http://localhost:8000
```

ブラウザで確認する場合は，次の URL を開きます．

```txt
http://localhost:8000
```

### ログの確認

```bash
docker logs python-server
```

ログを追い続ける場合は，次のようにします．

```bash
docker logs -f python-server
```

### 停止

```bash
docker stop python-server
```

### 再起動

停止したコンテナを再起動する場合は，次のようにします．

```bash
docker start python-server
```

### 削除

停止したコンテナを削除します．

```bash
docker rm python-server
```

## Go

### ディレクトリ構成

```txt
go-app/
├── Dockerfile
├── go.mod
└── main.go
```

### ビルド

```bash
cd go-app
docker build -t go-server .
```

### 起動

バックグラウンドで起動します．

```bash
docker run -d --name go-server -p 8001:8001 go-server
```

### 確認

```bash
curl http://localhost:8001
```

ブラウザで確認する場合は，次の URL を開きます．

```txt
http://localhost:8001
```

### ログの確認

```bash
docker logs go-server
```

ログを追い続ける場合は，次のようにします．

```bash
docker logs -f go-server
```

### 停止

```bash
docker stop go-server
```

### 再起動

停止したコンテナを再起動する場合は，次のようにします．

```bash
docker start go-server
```

### 削除

停止したコンテナを削除します．

```bash
docker rm go-server
```

## 起動中のコンテナ確認

起動中のコンテナは，次のコマンドで確認できます．

```bash
docker ps
```

停止済みのコンテナも含めて確認する場合は，次のようにします．

```bash
docker ps -a
```

## 停止方法

バックグラウンドで起動しているため，`Ctrl-C` では停止しません．

停止する場合は，次のコマンドを使います．

```bash
docker stop python-server
docker stop go-server
```

停止後にコンテナを削除する場合は，次のようにします．

```bash
docker rm python-server
docker rm go-server
```

## 補足

バックグラウンド起動では，`docker run` に `-d` を付けます．

また，コンテナを操作しやすくするために，`--name` で名前を付けています．

```bash
docker run -d --name python-server -p 8000:8000 python-server
```

以前のように `--rm` を付けると，停止時にコンテナが自動削除されます．

```bash
docker run --rm -p 8000:8000 python-server
```

一時的に試すだけなら `--rm` 付きでもよいですが，バックグラウンドで動かしてログ確認や再起動をしたい場合は，`--rm` を付けない方が扱いやすいです．

