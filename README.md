# Home Server Hands-On

Docker 上でシンプルな HTTP サーバーを立てるサンプルです．

次の 2 種類の実装があります．

- Python + FastAPI
- Go 標準ライブラリ

どちらも `/` にアクセスすると，シンプルな HTML を返します．

## Python + FastAPI

### ディレクトリ構成

```txt
python-app/
├── Dockerfile
├── requirements.txt
└── app
    └── main.py
````

### ビルド

```bash
cd python-app
docker build -t python-server .
```

### 起動

```bash
docker run --rm -p 8000:8000 python-server
```

### 確認

```bash
curl http://localhost:8000
```

ブラウザで確認する場合は，次の URL を開きます．

```txt
http://localhost:8000
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

```bash
docker run --rm -p 8001:8001 go-server
```

### 確認

```bash
curl http://localhost:8001
```

ブラウザで確認する場合は，次の URL を開きます．

```txt
http://localhost:8001
```

## 停止方法

フォアグラウンドで起動している場合は，`Ctrl-C` で停止できます．

`--rm` を付けて起動しているため，停止後にコンテナは自動で削除されます．

