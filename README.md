

# DB

* 備考
    * lat,lngからaddressを引くのは初回のみ、redis等にキャッシュ
        * {pos_id: address}

```
map{id, shop_id, pos_id}
shop{id, name, url, description, type}
pos_id{id, lat, lng}
```

