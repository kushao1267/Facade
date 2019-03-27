## facade

### Features

facade是一个全网通用的链接预览服务，它的功能如下:

* 支持REST API方式获取链接预览信息
* 支持自定义预览信息的字段
* 有很强的可扩展性，可自定义支持许多网站
* 使用Golang实现，有良好的性能

1.api模块
提供链接预览的API服务，使用gonic-gin搭建

2.config模块
加载配置，使用github.com/BurntSushi/toml库

3.db模块
数据库的一些封装，暂时只使用了redis(go-redis/redis库)

4.consts模块
常量模块

5.extractor模块
用于从页面中提取预览字段，可以加载多个handler模块

6.handler模块
自定义处理模块，可以在此处扩展用户自定义处理

7.requests
请求模块，封装了对链接的请求以及许多网络工具，返回完整的网页内容

### Install

### Usage
阅读README.md

### Example
见example