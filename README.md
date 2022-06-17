# Conference Manager (Server)

论文投稿系统的后端。

## 项目依赖

1. [Gin](https://github.com/gin-gonic/gin): Web 框架
0. [Gin-Cors](https://github.com/gin-contrib/cors): Gin 框架提供的跨域中间件
0. [GORM](https://gorm.io/docs/): 数据库 ORM 框架
0. [OSS](https://github.com/aliyun/aliyun-oss-go-sdk): 阿里云对象存储 SDK
0. [godotenv](https://github.com/joho/godotenv): 环境变量工具

## 环境变量

将 `.env.example` 复制到 `.env` 配置环境变量。

```dotenv
POSTGRES_DSN="host=localhost port=5432 user=name dbname=dbname password=pw sslmode=disable"
AUTH_SECRET=""          # Auth Token 的加密密钥，在生产环境中设置
GIN_MODE="debug"        # debug | test | release
LOG_LEVEL="debug"       # debug | info | warning | error
GORM_LOG_MODE="true"    # true : 打印详细日志 ｜ false : 不打印日志
ALLOW_ORIGIN=""         # 允许的跨域源，例如 "https://example.example
OSS_END_POINT="oss-cn-xxxx.aliyuncs.com"
OSS_ACCESS_KEY_ID="xxxx"
OSS_ACCESS_KEY_SECRET="xxxx"
OSS_BUCKET="xxxx"
```