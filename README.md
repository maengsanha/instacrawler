![Thumbnail](https://lh5.googleusercontent.com/proxy/r5D7LX7gbvXfuJU1SFAfCM1SerPt0KcBvR_R0qpXO_fsa39nwCKhyGE0UQbFP99XpSMRuPWrckLRnkoU747FW6EHY1_Gqf1xzhXYhJnIqIHizuhbBX3fh0sgdxbpIwJrDtC9g-uELzM-xYNfiw=s0-d)



# INSTAGRAM CRAWLER

[![LICENSE](https://img.shields.io/badge/license-Apache%202-blue)](https://github.com/joshua-dev/instacrawler/blob/master/LICENSE)
[![go version](https://img.shields.io/badge/go-1.14-00ADD8)](https://go.dev)



크롤링 솔루션에 대한 문서는 [여기](https://github.com/joshua-dev/instacrawler/blob/master/src/crawler/README.md)에 있습니다.



## :white_check_mark: TODO

* ### Model

  * Searcher
  
    - [x] TopSearch: 상단 검색창에서 검색 시 나오는 해시 태그 목록과 게시물 갯수를 크롤링
      ```go
      func (s *Searcher) TopSearch(query string) error
      ```

  * Parser
    
    - [ ] parseInstaPageSource: 인스타그램 페이지 소스 파싱 수행
      ```go
      func (p *Parser) parseInstaPageSource(url string) error
      ```

  * Crawler

    - [ ] Crawl: 해시 태그 검색 결과 나타나는 게시물들에 대한 크롤링 수행
      ```go
      func (c *Crawler) Crawl(query string) []Posting
      ```

    - [ ] Next: 다음 pagination을 반환
      ```go
      func (c *Crawler) Next(parsedInstaPageSource string) (string, error)
      ```

* ### View
  
  - [ ] API 서버 구축

* ### Controller
  
  - [ ] 게시물 구조체에 대한 String, MarshalJSON 구현


- [ ] 최적화: 병렬 처리

- [ ] API 명세서 작성

---

### [LICENSE](https://github.com/joshua-dev/instacrawler/blob/master/LICENSE)
