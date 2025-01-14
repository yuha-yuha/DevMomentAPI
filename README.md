# DevMomentAPI
## スライドURL
https://docs.google.com/presentation/d/1hIrj-sFH03UwQkq52BSIIs5jl8OGs82c/edit?usp=drive_link&ouid=108848457319239244395&rtpof=true&sd=true
## overview
このプロダクトはモックAPIを簡易的に立てることのできるCLIツールです。

## USAGE
任意のディレクトリに、モックしたいAPIの内容をjson形式でファイルに記述してください。


|key|content| 
----|---------
|apis|モックするAPIのまとまり(Array)|
|apis[n].path|APIのpath|
|apis[n].method|期待するリクエスト種類|
|apis[n].response|APIが返すresponseデータ(一般的に使われるJSONのvalueであればなんでも良い)|

``` sample.json
{
    "apis":[
        {
            "path":"/hoge",
            "method": "GET",
            "response": {
                "message":"json"
            }
        }
    ]
}
```
作成したファイルを元にmockAPIを作成するには以下のようなコマンドを叩いてください
ファイルは実行ディレクトリからの相対パスです
`go run main.go ./sample.json`


実際にlocalhostにリクエストを送ると設定した通りのmockAPIのデータが返ってきます。
```
curl localhost:8080/hoge
{"message":"json"}
```

### model expantion
頻繁に使うデータをモデリングして展開をすることができます。
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


### 工夫したコード
modelの展開を行う際にapis.responseのjson階層を把握する必要があった
responseの値は実際には`interface{}`型として扱っているため、interface{}が実際には何の型か判断しながら階層を下まで深ぼっていく回帰関数を定義した
https://github.com/yuha-yuha/DevMomentAPI/blob/main/services/model_unpack_for_response_json.go#L23
