# OG Gofer
Retrieves the embedded OpenGraph metadata for a given url, and returns it in a friendly JSON format

I've deployed the server for testing purposes at [og-gofer.herokuapp.com](https://og-gofer.herokuapp.com/), but you can also deploy it to your own Heroku instance by clicking here:

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

## API
### Request
There is only one endpoint at `/` which requires the `url` GET parameter, set to a valid URL.

Example:

[https://og-gofer.herokuapp.com/?url=https://www.youtube.com/watch?v=oHg5SJYRHA0](https://og-gofer.herokuapp.com/?url=https://www.youtube.com/watch?v=oHg5SJYRHA0)
### Response
The response is a JSON object with a top-level `success` boolean, and a nested `payload` dictionary. In the event of an error, `succcess` will be `false`, and the payload will contain an `error` key with a message explaining the issue:
``` json
> http https://og-gofer.herokuapp.com/?url=http://www.thisurldoesntexist.com

{
    "success": false,
    "payload": {
        "error": "HTTP Error retrieving URL http://www.thisurldoesntexist.com (Get http://www.thisurldoesntexist.com: dial tcp: lookup www.thisurldoesntexist.com on 172.16.0.23:53: no such host)",
        "url": "http://www.thisurldoesntexist.com"
    }
}
```

If everything goes well, `success` will be `true`, and the payload will contain an `ogTags` dictionary, as well as a tag counter,  `tagsFound`, and a unique tag counter, `uniqueTagsFound`.

```json
> http https://og-gofer.herokuapp.com/?url=https://www.nytimes.com/2017/11/28/us/politics/republicans-tax-bill-senate.html

{
    "payload": {
        "ogTags": {
            "og:description": [
                "Republicans took a significant step forward on Tuesday when a key panel passed the $1.5 trillion tax cut, clearing the way for a full Senate vote later in the week."
            ],
            "og:image": [
                "https://static01.nyt.com/images/2017/11/29/us/politics/29dc-tax5/29dc-tax5-facebookJumbo.jpg"
            ],
            "og:title": [
                "Republicans Clear Major Hurdle as Tax Bill Advances"
            ],
            "og:type": [
                "article"
            ],
            "og:url": [
                "https://www.nytimes.com/2017/11/28/us/politics/republicans-tax-bill-senate.html"
            ]
        },
        "tagsFound": 5,
        "uniqueTagsFound": 5,
        "url": "https://www.nytimes.com/2017/11/28/us/politics/republicans-tax-bill-senate.html"
    },
    "success": true
}
```
