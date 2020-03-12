## Instagram Crawling Solution

### All we have to do is repeating below loop.

* ```1```: go to Instagram page source with a query.
    
  https://www.instagram.com/explore/tags/{query}

* ```2```: parse post data from page source.
  
  post data of index ```idx```:
  
  entry_data.TagPage[0].graphql.hashtag.edge_hashtag_to_media.edges[idx].node

* ```3```: go to next pagination.

  end_cursor: get from ```2```
  next pagination url:
  
  https://www.instagram.com/explore/tags/{query}/?__a=1&max_id={end_cursor}

* go to ```1``` until ```has_next_page``` is true.
