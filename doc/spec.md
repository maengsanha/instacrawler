# Instagram Crawler API 명세서

[![GitHub](https://img.shields.io/badge/-GitHub-181717?logo=GitHub)](https://github.com/joshua-dev/instacrawler)

## 1. API 명세서 개요

  본 문서는 API 정의와 명세를 포함하며, 개발 목적으로 활용함.

<br />

## 2. API 목록

  - 1. /api/v1/topsearch
  - 2. /api/v1/crawl

<br />

## 3. API 명세서

### 1. /api/v1/topsearch
  
  <br>
  
  <img width="250" src="https://user-images.githubusercontent.com/62831866/78666393-44de1a00-7912-11ea-948b-fb77e5833ca8.jpeg">

  위와 같이 인스타그램 상단 검색 시 출력되는 연관 해시태그 목록을 json 형태로 반환합니다.

  <br>

  | method |       path        |  Request   |       Response       |
  | :----: | :---------------: | :--------: | :------------------: |
  | `GET`  | /api/v1/topsearch | (json) 검색어 | (json) 상단 해시태그 검색 결과 |

  <br>

  - Request Body 예시

    <br>

    - query: 검색어
  
    ```
    {
      "query": "spring"
    }
    ```

  <br>

  - Reponse 예시

    <br>

    - tag: 관련 해시태그
    - approximate: 대략적인 게시물 갯수
    - precise: 정확한 게시물 갯수
  
    ```json
    {
      "data": [
        {
          "tag": "spring",
          "approximate": "117M",
          "precise": "117420799"
        },
        {
          "tag": "springiscoming",
          "approximate": "3.3M",
          "precise": "3360549"
        },
        {
          "tag": "springflowers",
          "approximate": "2.3M",
          "precise": "2309203"
        },
        {
          "tag": "springfashion",
          "approximate": "3.4M",
          "precise": "3420750"
        }
      ]
    }      
    ```

<br>

### 2. /api/v1/crawl

  메타 검색을 수행하고, 결과를 json 형태로 돌려줍니다.

  | method |     path      |          Request          |    Response     |
  | :----: | :-----------: | :-----------------------: | :-------------: |
  | `GET`  | /api/v1/crawl | (json) 동의어를 포함한 메타 검색어 정보 | (json) 메타 검색 결과 |

  <br>

  - Request Body 예시
    <br>

    L2: 2차 카테고리의 검색어와 동의어를 모아놓은 iterable data structure
    <br>
    L3: 3차 카테고리의 검색어와 동의어를 모아놓은 iterable data structure
  
    ```json
    {
      "L2": [
        "가수", "singer", "vocalist"
      ],
      "L3": [
        "bts", "방탄소년단"
      ]
    }
    ```


  <br>
  
  - Response 예시
    <br>

    L2: 2차 카테고리에만 속하는 검색 결과들의 iterable data structure
    <br>
    L3: 2차와 3차 카테고리 모두에 속하는 검색 결과들의 iterable data structure
    
    <br>
    
    text: (string) 게시물의 텍스트
    <br>
    url: (string) 게시물의 주소. 실제 주소는 앞에 `https://instagram.com/p/` 가 붙습니다.
    <br>
    src: (string) 게시물 커버 사진 이미지의 URL
    <br>
    like: (int) 게시물이 받은 좋아요 갯수

    ```json
    {
      "L2": [
        {
          "text": "#좋반 #좋테 #follow #dm #가수지망생 #가수 #소속사 #f4f",
          "url": "Ej7e2uOP-223",
          "src": "https://scontent-gmp1-1.cdinstagram/v/23kksok.d892",
          "like":180
        },
        {
          "text": "#trusty #diskey #트러스티 #디스키",
          "url": "nERT39ui-1",
          "src": "https://scontent-gmp1-1.cdinstagram/v/13141kdiw",
          "like": 218
        }
      ],
      "L3": [
        {
          "text": "#bts #army #btsarmy #태태 #태형 #김태형 #태형이 #fff #sure #f4f",
          "url": "EO790-dsm732",
          "src": "https://scontent-gmp1-1.cdinstagram/v/763je0jd",
          "like": 456
        },
        {
          "text": "거울 탐난다.",
          "url": "56Fti980-112",
          "src": "https://scontent-gmp1-1.cdinstagram/v/241rt50",
          "like": 277
        }
      ]
    }
    ```
