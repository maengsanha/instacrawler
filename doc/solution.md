## **Instagram Crawling Solution**

### All we have to do is repeating below loop.

* `1`: go to Instagram hashtag explore page source with a query.
    
  https://www.instagram.com/explore/tags/{query}/

* `2`: parse post data from page source. (Init)
  
  post data of index `idx`:
  
  entry_data.TagPage[0].graphql.hashtag.edge_hashtag_to_media.edges[idx].node

  next GraphQL endpoint: **end_cursor**

* `3`: go to next pagination, and you will get freaky **json**. (Next)

  next pagination url:
  
  https://www.instagram.com/explore/tags/{query}/?__a=1&max_id={end_cursor}

  then parse post data, next **end_cursor**.

* `4`: go to `3` until **has_next_page** is true.

***

## Result

![result_screen_shot](https://user-images.githubusercontent.com/29545214/77256323-df9eed80-6cb0-11ea-841d-0c6a5b759644.png)
