# key-value-store app

In memory key-value store olarak çalışan bir REST-API servisi.

- key ’i set etmek için bir endpoint  [Done]
- key ’i get etmek için bir endpoint [Done]
- Komple data’yı flush etmek için bir endpoint[Done]
- Belirli bir interval’da (N dakikada bir) dosyaya
  kaydetmeli [Done]
- Uygulama durup tekrar ayağa kalktığında, eğer
  kaydedilmiş dosya varsa, tekrar varolan verileri
  hafızaya yüklemeli ( /tmp/TIMESTAMP-data.json) [Done]

### Localde Çalıştırmak için

`git clone https://github.com/kaantecik/key-value-store`

`go build -o /bin/app cmd/key-value-store/main.go`

`./bin/app`

### Endpoints

```
|----------|--------------------|--------------------| 
|  METHOD  |     ENDPOINT       |    BODY            |
|----------|--------------------|--------------------|
|   POST   |   /api/cache/get   |  { "key": "foo" }  |
|----------|--------------------|--------------------|  
|          |                    |                    |
|          |                    |  {                 |
|          |                    |    "key": "foo",   |
|   POST   |   /api/cache/set   |    "value": "bar"  |
|          |                    |  }                 |
|          |                    |                    |
|----------|--------------------|--------------------|
|  DELETE  |  /api/cache/flush  |         -          |
|----------|--------------------|--------------------|         
```

[Heroku üzerinde çalışan servisin linki](https://limitless-bayou-61923.herokuapp.com/)

[Go Report Sonucu](https://goreportcard.com/report/github.com/kaantecik/key-value-store)