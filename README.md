# GoJack - Chiasenhac.vn API for Go

```console
go get github.com/ReeganExE/gojack
gojack
Listening on [::]:8989
```

## Docker

```sh
docker run --rm -p 8989:8989 reeganexe/gojack
```

### Get track detail
```http
GET http://localhost:8989/detail?link=http://chiasenhac.vn/nghe-album/carry-you-home~tisto-stargate-aloe-blacc~tsvd3w6cqmw9kv.html
```

```json
{
  "album": "http://chiasenhac.vn/nghe-album/carry-you-home~tisto-stargate-aloe-blacc~tsvd3w6cqmw9kv.html",
  "artist": "Tiësto; StarGate; Aloe Blacc",
  "cover": "http://125.212.211.135/~csn/data/cover/78/77340.jpg",
  "title": "Carry You Home",
  "url": "http://chiasenhac.vn/nghe-album/carry-you-home~tisto-stargate-aloe-blacc~tsvd3w6cqmw9kv.html",
  "media": "http://data00.chiasenhac.com/downloads/1831/2/1830207-50f6a2cd/320/Carry You Home - Ti__sto_ StarGate_ Aloe.mp3",
  "palette": {
    "color": {
      "background": "#f5f90a",
      "text": "#000000"
    },
    "dark": {
      "background": "#04452a",
      "text": "#ffffff"
    },
    "darkMuted": {
      "background": "#3d563f",
      "text": "#ffffff"
    },
    "light": {
      "background": "#e77670",
      "text": "#000000"
    },
    "lightMuted": {
      "background": "#c5d1a5",
      "text": "#000000"
    },
    "muted": {
      "background": "#6da161",
      "text": "#000000"
    }
  }
}
```

### Get track list
```http
GET http://localhost:8989/tracks?link=http://chiasenhac.vn/nghe-album/tam-su-nang-xuan~dam-vinh-hung~tsvcr0dqqvanme.html
```

```json
[
  {
    "artist": "Đàm Vĩnh Hưng",
    "title": "Tâm Sự Nàng Xuân",
    "url": "http://chiasenhac.vn/nghe-album/tam-su-nang-xuan~dam-vinh-hung~tsvcr0dqqvanme.html"
  },
  {
    "artist": "Đàm Vĩnh Hưng",
    "title": "Xuân Trên Đất Việt (Rumba Version)",
    "url": "http://chiasenhac.vn/nghe-album/xuan-tren-dat-viet-rumba-version~dam-vinh-hung~tsvcr0dcqvanmv.html"
  },
  {
    "artist": "Đàm Vĩnh Hưng",
    "title": "Xuân Trên Đất Việt (Techno Version)",
    "url": "http://chiasenhac.vn/nghe-album/xuan-tren-dat-viet-techno-version~dam-vinh-hung~tsvcr0dbqvanmq.html"
  },
  {
    "artist": "Đàm Vĩnh Hưng",
    "title": "Xuân Trên Đất Việt (Reggae Version)",
    "url": "http://chiasenhac.vn/nghe-album/xuan-tren-dat-viet-reggae-version~dam-vinh-hung~tsvcr0ddqvanmm.html"
  },
  {
    "artist": "Đàm Vĩnh Hưng",
    "title": "Mẹ Việt Nam",
    "url": "http://chiasenhac.vn/nghe-album/me-viet-nam~dam-vinh-hung~tsvcr35wqvawh9.html"
  }
]
```


### Get recently shared tracks
```http
GET http://localhost:8989/preset?link=http://chiasenhac.vn/mp3/vietnam/
```

```json
[
  {
    "artist": "Bích Phương",
    "title": "Drama Queen",
    "url": "http://chiasenhac.vn/nghe-album/drama-queen~bich-phuong~tsvbv5qtqq2hef.html"
  },
  {
    "artist": "Andiez",
    "title": "Mãi Mãi Sẽ Hết Vào Ngày Mai",
    "url": "http://chiasenhac.vn/nghe-album/mai-mai-se-het-vao-ngay-mai~andiez~tsvbv6cvqq2kv2.html"
  },
  {
    "artist": "Ly Mít; Huy Vạc",
    "title": "Thằng Điên (Cover)",
    "url": "http://chiasenhac.vn/nghe-album/thang-dien-cover~ly-mit-huy-vac~tsvbv7dzqq2tm1.html"
  },
  {
    "artist": "Việt Athen",
    "title": "Ma Ma",
    "url": "http://chiasenhac.vn/nghe-album/ma-ma~viet-athen~tsvbvm70qq28tn.html"
  },
  {
    "artist": "Bùi Lan Hương",
    "title": "Ngày Chưa Giông Bão",
    "url": "http://chiasenhac.vn/nghe-album/ngay-chua-giong-bao~bui-lan-huong~tsvbvwwcqq299v.html"
  },
  {
    "artist": "Vũ Cát Tường",
    "title": "Leader",
    "url": "http://chiasenhac.vn/nghe-album/leader~vu-cat-tuong~tsvbvr6wqq2ak9.html"
  },
  {
    "artist": "Vũ Cát Tường",
    "title": "The Party Song",
    "url": "http://chiasenhac.vn/nghe-album/the-party-song~vu-cat-tuong~tsvbvr60qq2akn.html"
  }
]
```

# LICENSE
MIT @ Ninh Pham

# DISCLAIMER
Use for study & research purpose.
