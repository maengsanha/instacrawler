<img src="https://lh5.googleusercontent.com/proxy/r5D7LX7gbvXfuJU1SFAfCM1SerPt0KcBvR_R0qpXO_fsa39nwCKhyGE0UQbFP99XpSMRuPWrckLRnkoU747FW6EHY1_Gqf1xzhXYhJnIqIHizuhbBX3fh0sgdxbpIwJrDtC9g-uELzM-xYNfiw=s0-d">

<br/>

# **INSTAGRAM CRAWLER**

[![license](https://img.shields.io/badge/license-MIT-blue)](https://github.com/joshua-dev/instacrawler/blob/master/LICENSE)
[![go version](https://img.shields.io/badge/go-1.14-00ADD8)](https://go.dev)

<br/>

API 명세서는 [**여기**](https://github.com/joshua-dev/instacrawler/blob/master/doc/spec.md)에 있습니다.

크롤링 솔루션에 대한 문서는 [**여기**](https://github.com/joshua-dev/instacrawler/blob/master/doc/solution.md)에 있습니다.

<br/>

## :white_check_mark: TODO

- ### _Core_

  - [x] InstaPost: 인스타그램 게시물

    ```go
    // InstaPost is an Instagram post type.
    type InstaPost struct
    ```

  <br/>

  - [x] SpriteHashtag: 인스타그램 상단 검색 시 제공되는 관련 해시 태그 목록

    ```go
    // SpriteHashtag is a type of associated search term
    // recommended when searching on top of Instagram.
    type SpriteHashtag struct
    ```

<br/>

- ### _Controllers_

  - top

    - [x] Search: 상단 검색창에서 검색 시 나오는 해시 태그 목록과 게시물 갯수 크롤링

      ```go
      // Search implements top search on Instagram with a given query.
      func Search(query string) (hashtags []core.Hashtag, err error)
      ```

  <br/>

  - crawler

    - [x] Generator: 인스타그램 게시물 크롤링 함수 생성
  
      ```go
      // Generator generates an Instagram crawler.
      func Generator(query string) func() ([]core.InstaPost, string, error)
      ```

      

  <br/>

  - Meta
  
    - [x] _AND: 2계층과 3계층 검색 결과에 AND 연산 수행
  
      ```go
      // _AND implements AND operation.
    func _AND(secondLayer, thirdLayer core.PostSet) (secondLayerResult *core.InstaPosts, thirdLayerResult *core.InstaPosts)
      ```

      
  
    <br/>
  
    - [x] _OR: 같은 계층의 검색 결과들에 중복 없는 OR 연산 수행
    
      ```go
      // _OR implements OR operation.
  func _OR(layer []*crawler.Crawler) (set core.PostSet)
      ```
  
      
  
    <br/>
    
    - [x] categorize: 2계층과 3계층 검색 결과로 OR/AND 연산 수행
    
  ```go
      // categorize implements categorization of meta-search.
    func categorize(secondLayer, thirdLayer []*crawler.Crawler) ([]core.InstaPost, []core.InstaPost)
      ```
  
      
    
    <br/>
    
    - [x] Search: 2계층과 3계층 검색어로 메타 검색 구현
    
      ```go
      // Search implements meta-search with the given search terms of second layer and third layer.
    func Search(secondLayer, thirdLayer []string) ([]byte, error)
      ```

<br/>

- ### _Routers_

  - [x] /api/v1/topsearch
  - [x] /api/v1/crawl

<br/>
<br/>

- ### _Optimization_

  - [x] Concurrency
  - [ ] Pipelining
  - [ ] Heuristic Search

<br/>

- [x] API 명세서 작성

<br/>
<br/>

## [LICENSE](https://github.com/joshua-dev/instacrawler/blob/master/LICENSE)

