# DevMomentAPI

## overview
 this is mockAPI tool. 
 write API specification to Json file
## USAGE
Write json moment api format `./sample.json` at the now.
 
```
{
    apis:[
        {
            path:"/hoge",
            method: "GET",
            response: {
                "message":"json"
            }
        }
    ]
}
```

next, execution this command
`go run main.go ./sample.json`

Access localhost:8080/hoge GET method
```
curl localhost:8080/hoge
{"message":"json"}
```



### model expantion

```
{
    apis:[
        {
            path:"/hoge",
            response: {
                "message":"${user}"
            }
        }
    ]

    model:{
        "user": {"name": "hoge"}
    }
}
```
