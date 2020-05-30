#### Instagram Crawler API specification

[![GitHub](https://img.shields.io/badge/-GitHub-181717?logo=GitHub)](https://github.com/joshua-dev/instacrawler)

**1. API 명세서 개요**

본 문서는 API 정의와 명세를 포함하며, 개발 목적으로 활용함.

<br />

**2. API 목록**

- 1. /api/v1/topsearch
- 2. /api/v1/crawl

<br />

**3. API 명세서**

0. Server domain:Port number

   Public DNS(IPv4): `http://ec2-13-125-96-172.ap-northeast-2.compute.amazonaws.com:3000`

   <br />

1. /api/v1/topsearch

     <br />

  <img width="250" src="https://user-images.githubusercontent.com/62831866/78666393-44de1a00-7912-11ea-948b-fb77e5833ca8.jpeg">

​	위와 같이 인스타그램 상단 검색 시 출력되는 연관 해시태그 목록을 json 형태로 반환합니다.

​	<br />

| method |       path        |     Request     |            Response            |
| :----: | :---------------: | :-------------: | :----------------------------: |
| `GET`  | /api/v1/topsearch | (string) 검색어 | (json) 상단 해시태그 검색 결과 |

​	<br />

- Reponse Body 예시

  <br />

  - name: 관련 해시태그
  - search_result_subtitle: 대략적인 게시물 갯수
  - media_count: 정확한 게시물 갯수

  ```json
  {
    [
      {
        "name": "café",
        "search_result_subtitle": "7m posts",
        "media_count": "7080757"
      },
      {
        "name": "cafe",
        "search_result_subtitle": "60.5m posts",
        "media_count": "60537500"
      },
      {
        "name": "cafè",
        "search_result_subtitle": "205k posts",
        "media_count": "205626"
      },
      {
        "name": "cafehopper",
        "search_result_subtitle": "184k posts",
        "media_count": "184403"
      }
    ]
  }
  ```

<hr />

2. /api/v1/crawl

메타 검색을 수행하고, 결과를 json 형태로 돌려줍니다.

| method |     path      |                 Request                 |       Response        |
| :----: | :-----------: | :-------------------------------------: | :-------------------: |
| `POST` | /api/v1/crawl | (json) 동의어를 포함한 메타 검색어 정보 | (json) 메타 검색 결과 |

  <br />

- Request Body 예시
  <br />

  higher_layer: (array) 상위 카테고리의 검색어와 동의어를 모아놓은 container

  lower_layer: (array) 하위 카테고리의 검색어와 동의어를 모아놓은 container

  higher_layer_cache: (array) higher_layer의 각 검색어들의 다음 pagination 주소들을 모아놓은 container. higher_layer와 길이가 반드시 같아야하며, 처음 요청 시엔 higher_layer의 길이만큼 빈 문자열이 있도록 초기화합니다.

  lower_layer_cache: (array) lower_layer의 각 검색어들의 다음 pagination 주소들을 모아놓은 container. lower_layer와 길이가 반드시 같아야하며, 처음 요청 시엔 lower_layer의 길이만큼 빈 문자열이 있도록 초기화합니다.

  **higher_layer와 higher_layer_cache의 길이가 서로 다르거나 lower_layer와 lower_layer_cache의 길이가 서로 다르면 null이 반환되니 유의하세요. (Status code: 400 BadRequest)**
  
  ```json
  {
    "higher_layer": ["cafe", "cafestagram", "cafetour"],
    "lower_layer": ["starbucks", "스타벅스"],
    "higher_layer_cache": [
      "QVFBQUNUTTlFa3lxYkp1TnFrcXc1Q01EcUR5RzFMMlBleUo2aGd6VUlkcTNjTnU5d1dzUmphTFFldXdrNV9vNGRrdWJMeVdiNGNxTDR2MGRwOXZFUXF0Zg==",
      "QVFBRDFFcmdrcVRUYlltNHFGTFVnXzFfS2JHeFJfdXhvMm5hMGdiY29zS25zVlNzOXl6aGpqWmt6Sjc1RkY5RzhfclctSlN5ckJEaWdTbTlkME5NNVdZMQ==",
      "QVFERFprbF9ZUS0yaGEzaEEyLTBHYTJhTV9heUd0QVZGci12NFg4andGeXdISnh0N0hsSXI2ajZnTllOcUxiTkpOa24tX3RMZl9tWndvaDZBcVRCZEpoaA=="
    ],
    "lower_layer_cache": [
      "QVFESC1jVnl3RFNwOEoycmtTWU5kbzh2UGFlT1RMSm1iSlFobmNmVkZ1cHloWjdxRllKcFQwaExtZ1FZeHh4YUE0T3hlQmpUVVNSeU1FbWVwWGt2S0loSw==",
      "QVFCdUEwRkF4aWtudC1nd2ljVEtzV3JWY1dOUF9rYUdlak5meDBjbzhmQ3R6Umtpa3cybmh2TGlSTlk1Vk5veEJMRmpkTzdSMzlCOExBWXZ1TndFcXg5Mw=="
    ]
}
  ```
  
    <br />

* Response Body 예시
  <br />

  higher_layer: (array) 상위 카테고리에만 속하는 검색 결과들의 container

  lower_layer: (array) 상위 카테고리와 하위 카테고리 모두에 속하는 검색 결과들의 container

  higher_layer_cache: (array) higher_layer의 각 검색어들의 다음 pagination 주소들을 모아놓은 container

  lower_layer_cache: (array) lower_layer의 각 검색어들의 다음 pagination 주소들을 모아놓은 container

  <br />

  text: (string) 게시물의 텍스트

  url: (string) 게시물의 주소. 실제 주소는 앞에 `https://www.instagram.com/p/` 가 붙습니다.

  src: (string) 게시물 썸네일 이미지의 URL

  like: (int) 게시물이 받은 좋아요 갯수

  ```json
  {
    "higher_layer": [
      {
        "text": "Доброе утро, лол... #доброеутро #лол #рязань #россия#провинция #художник#фото #город#городскойпейзаж #позитив #film #fun #artist #ussr#doll #art#amsterdam#model #kunst #clee #cat#fashion #focus #cofe #cosplay#instagood #smile#aesthetic",
        "url": "B_mF4EjqNl7",
        "src": "https://scontent-gmp1-1.cdninstagram.com/v/t51.2885-15/sh0.08/e35/c0.129.1034.1034a/s640x640/95215546_725892411549121_3193384201677613315_n.jpg?_nc_ht=scontent-gmp1-1.cdninstagram.com&_nc_cat=101&_nc_ohc=Omw1LX-TfuUAX8f8SqF&oh=9db818251361c6546dce1caff936e763&oe=5ED3816C",
        "like": 16
      },
      {
        "text": "Wednesday night prayer from the #northumbriacommunity with added wine #community #love #faith #church #churchofengland #cofe #sacredspace #Manchester #flixton #trafford #easter",
        "url": "B_lCgYvhah2",
        "src": "https://scontent-gmp1-1.cdninstagram.com/v/t51.2885-15/e35/c0.150.540.540a/95262123_142863047281654_4603480536606033666_n.jpg?_nc_ht=scontent-gmp1-1.cdninstagram.com&_nc_cat=110&_nc_ohc=j2UHP62HNkwAX-jebJz&oh=ca30f0635f45ebf1cbc7b1fcfe30709b&oe=5EAC9015",
        "like": 2
      }
    ],
    "lower_layer": [
      {
        "text": "#momofgirls #girldad #bows #hairbows #girlygirl #prettygirls #babygirls #mommyandme #mommytobe #itsagirl #mickeymouse #disney #starbucks #coffee #minniemouse #shopsmall #supportsmallbusiness #sarapebow #mexican #colorful #cowgirl #coyboyboots #boots #bowsandboots #virgencita #easter #bunny #sfgiants #lol #loldolls",
        "url": "B_mKt1KFYU3",
        "src": "https://scontent-gmp1-1.cdninstagram.com/v/t51.2885-15/sh0.08/e35/s640x640/94890594_128293228823301_6353733521411040583_n.jpg?_nc_ht=scontent-gmp1-1.cdninstagram.com&_nc_cat=109&_nc_ohc=3Yy02sOzbcQAX_Vswjk&oh=efa97956384df18e9c1d90687e9f81bc&oe=5ED5BE2B",
        "like": 1
      },
      {
        "text": "@nohutlupil4v \n#aktiftakipci #takipedenitakipederim #takipet #aktiftakipçi #instadaily #takipplus #seguidores #takipetkinliği #chuvadelikes #takipetanindakazan #likeforlikes #instagram #sdv❤️ #instagood #begeniyebegeni #begeni #geritakipyapiyorum #geritakipvar #instalike #gtvardir #sigaosbaloes #sdv #rtb #sdvtodos #starbucks #streetphotography #sdvgeral #instalike #instafashion",
        "url": "B_mKQmQAD5a",
        "src": "https://scontent-gmp1-1.cdninstagram.com/v/t51.2885-15/sh0.08/e35/c10.0.730.730a/s640x640/95542093_935289073610807_693225298869737381_n.jpg?_nc_ht=scontent-gmp1-1.cdninstagram.com&_nc_cat=111&_nc_ohc=fWI676klOH8AX_wmaL0&oh=cc0705b754af875bcc6c417d72273955&oe=5ED36DA7",
        "like": 19
      }
    ],
    "higher_layer_cache": [
      "QVFBQUNUTTlFa3lxYkp1TnFrcXc1Q01EcUR5RzFMMlBleUo2aGd6VUlkcTNjTnU5d1dzUmphTFFldXdrNV9vNGRrdWJMeVdiNGNxTDR2MGRwOXZFUXF0Zg==",
      "QVFBRDFFcmdrcVRUYlltNHFGTFVnXzFfS2JHeFJfdXhvMm5hMGdiY29zS25zVlNzOXl6aGpqWmt6Sjc1RkY5RzhfclctSlN5ckJEaWdTbTlkME5NNVdZMQ==",
      "QVFERFprbF9ZUS0yaGEzaEEyLTBHYTJhTV9heUd0QVZGci12NFg4andGeXdISnh0N0hsSXI2ajZnTllOcUxiTkpOa24tX3RMZl9tWndvaDZBcVRCZEpoaA=="
    ],
    "lower_layer_cache": [
      "QVFESC1jVnl3RFNwOEoycmtTWU5kbzh2UGFlT1RMSm1iSlFobmNmVkZ1cHloWjdxRllKcFQwaExtZ1FZeHh4YUE0T3hlQmpUVVNSeU1FbWVwWGt2S0loSw==",
      "QVFCdUEwRkF4aWtudC1nd2ljVEtzV3JWY1dOUF9rYUdlak5meDBjbzhmQ3R6Umtpa3cybmh2TGlSTlk1Vk5veEJMRmpkTzdSMzlCOExBWXZ1TndFcXg5Mw=="
    ]
  }
  ```

