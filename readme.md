# COVID19 大数据可视化系统

>作者: **Ethan.Wang**
>
>版本号: **1.0.0**
>
>更新时间: **2022-12-03 21:33:10**



## 在线案例

https://welcome.covid19.ethan9.cn/



## 运行环境

Golang: `1.19.3`

NodeJs: `16.17.0`

- Vue: `3.2.41`
- Pinia: `2.0.23`
- Vite: `3.2.3`
- Element Plus: `2.2.20`
- Vue-router: `4.1.6`

MySQL: `5.7`

Redis: `6.2.6`

Docker: `20.10.21`

Docker-Compose: `2.12.2`



## 部署

> 提醒：
>
> 1. 使用前请编辑配置文件，不然系统无法运行
>
> 2. 由于开发微信登录接口时使用微信公众号接口开发
>
>    所以后端微信配置中微信公众号必须是：`微信服务号` 或`有登录权限的接口`
>    详情请看：[>>微信公众号登录接口文档<<](https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html)

测试系统功能可以使用`微信公众号测试平台`

[>>微信公众平台接口测试地址<<](https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login)



### 前端

> 请修改前端配置文件
>
> `VITE_SERVER_ADDR` 请修改为线上地址，不然微信登录不生效，提示参数错误
>
> `VITE_SIGN_KEY` 签名算法KEY，修改后后端签名KEY同样需要修改，不然前端请求无法获得后端数据

目录： `web/.env.production`

```ini
ENV = 'production'

# web 端口号 | 这个配置不要动
VITE_WEB_PORT = 8080

# server 地址 | 请修改为线上地址
VITE_SERVER_ADDR = https://covid19.ethan9.cn
# server 端口号 | 这个配置不要动
VITE_SERVER_PORT = 8088
# server 接口路径 | 这个配置不要动
VITE_API_PATH = /api

# 签名KEY | 修改后，后端签名KEY也需要修改
VITE_SIGN_KEY = E211655B25D54FC58144E9BDA583A96D
```



### 后端

>请前往后端，复制`config.template.ini`修改文件名为`config.ini`
>
>需要做的事情：
>
>1. 修改公众号配置，不配置系统无法登录
>2. 修改安全配置
>   1. 修改`签名密钥`后，前端配置项中签名密钥同样需要修改
>   2. 修改`JWT密钥`，该配置项只在后端校验，前端只携带该TOKEN返回给后端做身份认证
>
>提示：
>
>1. 服务器配置 | 数据库 | Redis，修改后请前往根目录下的`deploy/docker-compose.yaml`配置中同步修改
>2. `日志配置`，请根据自己的需求调整即可

目录：`server/config/config.template.ini`

```ini
;=====
; 服务器配置
;=====
[server]
; 运行环境配置
; 选项: debug->开发环境 | prod->生产环境
env = prod
; 端口号
port = 8088

;=====
; 日志
;=====
[log]
; 日志目录
dir = log
; 显示行
; 选项: true | false
show_line = true
; 输出编码格式
; 选项: json->以json格式输出 | console->以控制台格式输出
format = console
; 栈名
stacktrace_key = stacktrace
; 编码级
; 选项:
; LowercaseLevelEncoder-小写编码器-默认
; LowercaseColorLevelEncoder-小写编码器带颜色
; CapitalLevelEncoder-大写编码器
; CapitalColorLevelEncoder-大写编码器带颜色
encode_level = LowercaseLevelEncoder
; 日志是否输出到控制台
; 选项: true | false
log_in_console = false
; 前缀
; 例如：[abc] xxxxx  abc就是前缀
prefix = [C19]
; 日志级别
; 选项: debug | info | warn | error | dpanic | panic | fatal
level = info
; 在进行切割之前，日志文件的最大大小（以MB为单位）
max_size = 10
; 保留旧文件的最大个数
; 选项: 0就是全部保留
max_backups = 0
; 保留旧文件的最大天数
max_age = 180

;=====
; 数据库
;=====
[database]
; 地址
host = 172.19.19.13
; 端口号
port = 3306
; 用户名
username = root
; 密码
password = pRd6KL7RnA29JA8
; 数据库名
dbname = c19
; 设置空闲连接池中连接的最大数量
max_idle_conns = 10
; 设置打开数据库连接的最大数量
max_open_conns = 100

;=====
; 公众号配置
;=====
[mp]
; 微信公众号AppID
appid = 
; 微信公众号Secret
secret = 
; 微信公众号TOKEN
token = 

;=====
; Redis
;=====
[redis]
; 地址
host = 172.19.19.14
; 端口号
port = 6379
; 密码
password = 7DCsx0EWNFTRI60
; 数据库位置
db = 0

;=====
; 安全配置
;=====
[safe]
; 签名密钥
sign = E211655B25D54FC58144E9BDA583A96D
; JWT密钥
jwt = 4F38FC20FCDA408F8355FFA7219C14F9
```



### 数据库

本系统中使用的2种数据库，分别是：`Redis`、`MySQL`

如需修改数据库配置请前往`deploy/docker-compose.yaml`

提示：

- 修改数据库配置后请前往`后端`配置文件中修改相关配置参数



### 启动命令

```bash
# 使用docker-compose启动四个容器
docker-compose -f deploy/docker-compose.yaml up
# 如果您修改了某些配置选项,可以使用此命令重新打包镜像
docker-compose -f deploy/docker-compose.yaml up --build
# 使用docker-compose 后台启动
docker-compose -f deploy/docker-compose.yaml up -d
# 使用docker-compose 重新打包镜像并后台启动
docker-compose -f deploy/docker-compose.yaml up --build -d
# 服务都启动成功后,使用此命令行可清除none镜像
docker system prune
```



## 日志目录

> 如果出现无法解决的问题，请查看日志目录

**前端Nginx日志目录：**

```bash
~/docker/covid19/web/log
```

**后端日志目录：**

```bash
~/docker/covid19/server/log
```

