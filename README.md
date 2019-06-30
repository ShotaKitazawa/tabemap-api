

# DB

* 備考
    * lat,lngからaddressを引くのは初回のみ、redis等にキャッシュ
        * {pos_id: address}

```
map{id, shop_id, pos_id, type}
shop{id, name, url, description}
pos_id{id, lat, lng}
```

