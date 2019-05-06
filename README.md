# development environment
これはGolang + Vue.js + MariaDBの開発用コンテナイメージのサンプルリポジトリです。

## 起動方法
```
$ docker-compose up -d --build   
```

## WEBサーバー
### projectの作成
```
$ docker-compose exec web vue create [project name]
```
### サーバーの起動
```
$ docker-compose exec web sh

#app> cd [project name] && yarn serve   
```

## APIサーバー
api/main.goにmain関数があります。
PINGを返すだけのサンプルメソッドのみ実装。

