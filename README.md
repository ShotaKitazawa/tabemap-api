

# DB

* 備考
    * lat,lngからaddressを引くのは初回のみ、redis等にキャッシュ
        * {pos_id: address}

```
shop{id, name, url, description, type, lat, lng}
```

