![thumbnail](https://lh5.googleusercontent.com/proxy/r5D7LX7gbvXfuJU1SFAfCM1SerPt0KcBvR_R0qpXO_fsa39nwCKhyGE0UQbFP99XpSMRuPWrckLRnkoU747FW6EHY1_Gqf1xzhXYhJnIqIHizuhbBX3fh0sgdxbpIwJrDtC9g-uELzM-xYNfiw=s0-d)



# **INSTAGRAM CRAWLER**

[![license](https://img.shields.io/badge/license-MIT-green)](https://github.com/joshua-dev/instacrawler/blob/master/LICENSE)
[![go version](https://img.shields.io/badge/go-1.14-00ADD8)](https://go.dev)



크롤링 솔루션에 대한 문서는 [**여기**](https://github.com/joshua-dev/instacrawler/blob/master/src/crawler/README.md)에 있습니다.



## :white_check_mark: TODO

* ### *Model*

  * Searcher

    - [x] New: Searcher 생성자
      ```go
      // New returns a new Searcher.
      func New() *Searcher
      ```
  
    - [x] TopSearch: 상단 검색창에서 검색 시 나오는 해시 태그 목록과 게시물 갯수 크롤링
      ```go
      // TopSearch performs a top search in Instagram with a given query.
      func (s *Searcher) TopSearch(query string) error
      ```

  * Crawler

    - [x] New: Crawler 생성자
      ```go
      // New returns a new Crawler.
      func New() *Crawler
      ```

    - [x] init: 해시 태그 검색 시 나오는 첫 페이지 소스 파싱
      ```go
      // init crawls page source from first Instagram hastag explore page with a given query.
      func (c *Crawler) init(query string) error
      ```

    - [x] next: end_cursor (GraphQL end point) 값을 이용해 다음 pagination에서 json 구조체 파싱
      ```go
      // next parses json struct from next pagination.
      func (c *Crawler) next(query string) error
      ```
    
    - [x] Crawl: init과 반복적인 next를 통해 해시 태그 크롤링 완성
      ```go
      // Crawl completes crawling from Instagram through init and repeated next.
      func (c *Crawler) Crawl(query string)
      ```

    - [ ] filter: 2, 3차 카테고리 검색어로 크롤링한 결과에 서로 AND 연산 수행

* ### *View*
  
  - [ ] API 서버 구축

* ### *Controller*
  
  - [x] 게시물 구조체에 대한 String, MarshalJSON 구현


- [ ] 최적화: 병렬 처리

- [ ] API 명세서 작성

---

### [LICENSE](https://github.com/joshua-dev/instacrawler/blob/master/LICENSE)
