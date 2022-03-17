# 使用：
- 克隆本项目
- 编译 `go build .`
- 首次运行时会在目录下生成 `config.json` 文件，文件内容如下：
``` bash
# ./fund
2022/03/17 11:25:34 create default config file success, path: './config.json'
# cat ./config.json
{
  "key":"",
  "host":"api.day.app",
  "fundcode":[
    "000001"
  ]
}
```
- 获取的信息使用 [bark](https://github.com/Finb/Bark) 进行通知. 需要下载 bark 客户端，将客户端上显示的 key 填入 `config.json` 文件中
- 将要获取的 fund 填入 `config.json` 中的 `fundcode`
