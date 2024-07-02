# DevMomentAPI
this makes json API server from json format. 
## USAGE
Write json moment api format `./sample.json` at the now.
 
```
{
    apis:[
        {
            path:"/hoge",
            response: {
                "message":"json"
            }
        }
    ]
}
```
Access localhost:8080/hoge
```
curl localhost:8080/hoge
{"message":"json"}
```
