## Facade

![GitHub](https://img.shields.io/github/license/kushao1267/facade.svg)
![GitHub repo size](https://img.shields.io/github/repo-size/kushao1267/facade.svg)
![AppVeyor](https://img.shields.io/appveyor/ci/kushao1267/facade.svg)
![Codecov](https://img.shields.io/codecov/c/github/kushao1267/facade.svg)


## Features

Facade是一个全网通用的链接预览服务，它的功能如下:

* 支持REST API方式获取链接预览信息
* 支持自定义预览信息的字段（已能够支持视频，feed，音频，标题，描述等字段）
* 支持提取图片url以及图片解析
* 有很强的可扩展性，能够自定义支持许多网站，可将technique自由搭配使用
* 使用Golang实现，有良好的性能
* 使用docker-compose，一键启动API和缓存服务


## Installation

`go get -u github.com/kushao1267/facade`


## Module

* api模块
提供链接预览的API服务，gonic-gin框架，[gin文档](https://gin-gonic.com/docs/)

* config模块
加载toml文件配置，使用github.com/BurntSushi/toml库, [详见](https://github.com/BurntSushi/toml)

* db模块
数据库封装模块，目前只使用redis，用的是go-redis/redis库, [详见](https://github.com/go-redis/redis)

* extractors模块
使用者能够自定义extractor，来组合使用已有的technique，也可以调节使用technique的优先级，从而保证输出预览信息的完善和精确。
例如: WeiboTechnique -> HeadTagsTechnique -> SemanticTagsTechnique 的调用优先级，能够在抓取weibo feed预览信息失败时，调用相应
的通用technique来兜底。

* techniques模块
techniques中每个technique都提供了针对特定网站的多字段提取方法；
此外，还有通用的common technique，在其他特定technique提取信息失败时，可以用来兜底。
使用者能够加入更多网站的technique，欢迎提PR :)

* utils
工具模块，包含加密相关工具，http网络请求工具，图片解析工具，时间相关工具等方法


## Example
见example文件夹


## LICENSE
[MIT License](https://github.com/kushao1267/facade/blob/master/LICENSE)