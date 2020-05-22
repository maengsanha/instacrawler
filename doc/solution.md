## **Instagram Crawling Solution**

### All we have to do is repeating below loop.

* `1`  Go to Instagram hashtag explore page with a given `query`.
  
  https://www.instagram.com/explore/tags/`query`/?__a=1

* `2`  Parse **json** data.
  
  post data of index `idx`:
  
  entry_data.TagPage[0].graphql.hashtag.edge_hashtag_to_media.edges[idx].node

  next GraphQL endpoint: **end_cursor**

* `3`  Go to next pagination, and you will get freaky **json**.

  next pagination url:
  
  https://www.instagram.com/explore/tags/`query`/?__a=1&max_id={end_cursor}

  then parse post data, next **end_cursor**.

* `4`  Go to `3` until **has_next_page** is true.



:heavy_exclamation_mark: Warning

You will often get **404 ERROR** while performing above `1` and `3`.

Do not panic, retry request, and you will get normal response.

