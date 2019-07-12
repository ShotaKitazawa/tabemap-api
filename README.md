

# DB

* 備考
    * lat,lngからaddressを引くのは初回のみ、redis等にキャッシュ
        * {pos_id: address}

```
shop{id, name, url, description, type, lat, lng}
```

# Jenkinsfile

やりたいこと

* golang コンテナの中で
    * go test
* diuid コンテナの中で
    * Docker build
    * Docker run app & Docker run mysql
    * Test with Postman
* テストが通ったら
    * Docker tag
    * Docker push
    * CDへ移行

