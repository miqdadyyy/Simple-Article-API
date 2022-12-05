# API Documentation
If you're using postman, please download this [collection](./api.postman_collection.json)

## Get All Article
**Query**
```
- limit (default: 20)
- page (default: 1)
- search (can be empty)
```
**Request**
```
GET /article
```
**200 Response**
```
{
    "status": 200,
    "message": "Success",
    "data": {
        "limit": 10,
        "page": 1,
        "total": 2,
        "articles": [
            {
                "id": 2,
                "username": "Some Random Guy",
                "title": "How to Lorem Ipsum Dolor Sit",
                "slug": "how-to-lorem-ipsum-dolor-sit-702e058f",
                "thumbnail": "https://avatars.dicebear.com/api/identicon/how-to-lorem-ipsum-dolor-sit-702e058f.svg?size=256",
                "content": "Vestibulum a tincidunt dolor. Mauris nec turpis a arcu rhoncus ultrices sit amet at felis. Aliquam e...",
                "created_at": "2022-12-06T02:09:11+07:00",
                "updated_at": "2022-12-06T02:09:11+07:00"
            },
            {
                "id": 1,
                "username": "Miqdad",
                "title": "Hello World",
                "slug": "hello-world-afdabb07",
                "thumbnail": "https://avatars.dicebear.com/api/identicon/hello-world-afdabb07.svg?size=256",
                "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis eget tincidunt risus, et auctor ipsum....",
                "created_at": "2022-12-06T02:07:46+07:00",
                "updated_at": "2022-12-06T02:07:46+07:00"
            }
        ]
    }
}
```

## Get Article Detail
**Param**
```
- slug (slug of the article) 
```
**Request**
```
GET /article/:path
```
**200 Response**
```
{
    "status": 200,
    "message": "Success",
    "data": {
        "article": {
            "id": 2,
            "username": "Some Random Guy",
            "title": "How to Lorem Ipsum Dolor Sit",
            "slug": "how-to-lorem-ipsum-dolor-sit-702e058f",
            "thumbnail": "https://avatars.dicebear.com/api/identicon/how-to-lorem-ipsum-dolor-sit-702e058f.svg?size=256",
            "content": "Vestibulum a tincidunt dolor. Mauris nec turpis a arcu rhoncus ultrices sit amet at felis. Aliquam erat volutpat. In eget ligula ut sem volutpat condimentum eu ut arcu. Cras lobortis sem vel est suscipit aliquet. Interdum et malesuada fames ac ante ipsum primis in faucibus. Aliquam a lacinia dolor, ac auctor mi.",
            "created_at": "2022-12-06T02:09:11+07:00",
            "updated_at": "2022-12-06T02:09:11+07:00"
        },
        "comments": [
            {
                "id": 1,
                "sub_comments": [
                    {
                        "id": 2,
                        "sub_comments": null,
                        "username": "Miqdad Child's",
                        "content": "Just try the nested comment feature",
                        "created_at": "2022-12-06T02:13:57+07:00",
                        "updated_at": "2022-12-06T02:13:57+07:00"
                    }
                ],
                "username": "Miqdad",
                "content": "Just try the comment feature",
                "created_at": "2022-12-06T02:13:21+07:00",
                "updated_at": "2022-12-06T02:13:21+07:00"
            }
        ]
    }
}
```

**404 Response**
```
{
    "status": 404,
    "message": "Not Found",
    "data": null
}
```

## Create Article
**Body**
```
- title (required)
- username (required)
- content (required)
```
**Request**
```
POST /article
{
    "title": "How to Lorem Ipsum Dolor Sit",
    "username": "Some Random Guy",
    "content": "Vestibulum a tincidunt dolor. Mauris nec turpis a arcu rhoncus ultrices sit amet at felis. Aliquam erat volutpat. In eget ligula ut sem volutpat condimentum eu ut arcu. Cras lobortis sem vel est suscipit aliquet. Interdum et malesuada fames ac ante ipsum primis in faucibus. Aliquam a lacinia dolor, ac auctor mi."
}
```

**200 Response**
```
{
    "status": 200,
    "message": "Success",
    "data": null
}
```

## Create Comment
**Param**
```
- slug (required) 
```
**Body**
```
- parent_id (optional)
- username (required)
- content (required) 
```
**Request**
```
{
    "username": "Miqdad",
    "content": "Just try the comment feature"
}
```
**200 Response**
```
{
    "status": 200,
    "message": "Success",
    "data": null
}
```