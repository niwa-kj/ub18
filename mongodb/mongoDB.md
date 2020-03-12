### mongoDB
参考：https://qiita.com/FukuharaYohei/items/d109f5c3f5dce0d972e4


### Install
```
niwa@niwa:~$ wget -qO - https://www.mongodb.org/static/pgp/server-4.2.asc | sudo apt-key add -
OK
niwa@niwa:~$
niwa@niwa:~$ echo "deb [ arch=amd64 ] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.2 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.2.list
deb [ arch=amd64 ] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.2 multiverse
niwa@niwa:~$

niwa@niwa:~$ sudo apt-get update

niwa@niwa:~$ sudo apt-get install -y mongodb-org
```
### 起動確認
```
niwa@niwa:/usr/bin$ sudo service mongod status
[sudo] password for niwa:
● mongod.service - MongoDB Database Server
   Loaded: loaded (/lib/systemd/system/mongod.service; disabled; vendor preset:
   Active: inactive (dead)
     Docs: https://docs.mongodb.org/manual
niwa@niwa:/usr/bin$ sudo service mongod start
niwa@niwa:/usr/bin$ sudo service mongod status
● mongod.service - MongoDB Database Server
   Loaded: loaded (/lib/systemd/system/mongod.service; disabled; vendor preset:
   Active: active (running) since Mon 2020-03-09 16:23:18 JST; 7s ago
     Docs: https://docs.mongodb.org/manual
 Main PID: 3852 (mongod)
   CGroup: /system.slice/mongod.service
           mq3852 /usr/bin/mongod --config /etc/mongod.conf

Mar 09 16:23:18 niwa systemd[1]: Started MongoDB Database Server.
niwa@niwa:/usr/bin$```
```
### 設定ファイル
```
/etc/mongod.conf
  1 # mongod.conf
  2
  3 # for documentation of all options, see:
  4 #   http://docs.mongodb.org/manual/reference/configuration-options/
...中略
 21 # network interfaces
 22 net:
 23   port: 27017
 24   bindIp: 127.0.0.1
```
### mongoDBのデータ構成イメージ
```
Database
  Collection
    Document
    ...
    Document
```
## CLI(コマンド)
### DB切り替え、無ければ作成
```
> use mydb;
switched to db mydb
> db.dropDatabase();
{ "ok" : 1 }
> use mydb;
switched to db mydb
```
### コレクション作成
```
> db.createCollection('CAR');
{ "ok" : 1 }
```

### ドキュメント作成と検索と削除
```
> db.CAR.insert({name:'honda',since:1948});
WriteResult({ "nInserted" : 1 })
> db.CAR.insert({name:'toyota',since:1937});
WriteResult({ "nInserted" : 1 })
> db.CAR.find();
{ "_id" : ObjectId("5e66f6d35e5315bf4f146b39"), "name" : "honda", "since" : 1948 }
{ "_id" : ObjectId("5e66f76d5e5315bf4f146b3a"), "name" : "toyota", "since" : 1937 }
>
> db.CAR.remove({});
WriteResult({ "nRemoved" : 3 })
```
### go言語のドライバ
参考：https://github.com/mongodb/mongo-go-driver

```
niwa@niwa:~/repo/ub18$ env | grep GO
GOPATH=/home/niwa/go
niwa@niwa:~/repo/ub18$
コンパイルが走ってしばらく待つ。
niwa@niwa:~/repo/ub18$ go get go.mongodb.org/mongo-driver/mongo

niwa@niwa:~/repo/ub18$
niwa@niwa:~/go/pkg$ ll linux_amd64/go.mongodb.org/mongo-driver/mongo.a
-rw-rw-r-- 1 niwa niwa 1956452 Mar 10 09:41 linux_amd64/go.mongodb.org/mongo-driver/mongo.a
niwa@niwa:~/go/pkg
```
