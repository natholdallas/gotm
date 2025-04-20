# README

项目还在测试完工

## 参考介绍 📄

本项目默认功能拥有以下:

1. 定时任务
2. 配置文件 & 热重载 `pkg/conf`
3. 一些常用的 sql 语句封装 `pkg/db/repos.go`
4. 默认用户模板 `pkg/db/schema.go`
5. 媒体文件模板 `pkg/db/schema.go`
6. 配置好的自动清理媒体文件的定时任务 `pkg/task`
7. 微型图床实现 `pkg/handler/media.go`
8. 一些常用的信息 `pkg/enum/enum.go`
9. 数据库在 dev 模式下的 `create-drop` 策略 **考虑到有些用户更喜欢使用 migration, 你可以自己调整** `pkg/db/db.go`
10. 数据校验实现 `pkg/fibers/validator.go`
11. fiber 一些常用的封装工具函数 `pkg/fibers/response.go` `pkg/fibers/request.go`
12. 全局的错误通用处理实现 `pkg/handler/error.go`
13. JWT 技术实现与处理及中间件 `pkg/mid/auth.go`
14. 一些可能会用到的数学函数和在你遇到 bug 时需要借助工具输出一个 struct 信息的工具
15. 谷歌三方验证登录支持 `pkg/mid/role.go` `pkg/client/google.go` `pkg/db/schema.go`
16. 路由注册 `pkg/router/router.go`
17. 分页实现 `pkg/handler/struct.go` `pkg/db/repos.go`
18. 打包命令相关 `scripts`

## 主要使用到的包 📦

1. `github.com/Pallinder/go-randomdata v1.2.0` 随机数据
2. `github.com/go-playground/validator/v10` v10.26.0 数据校验
3. `github.com/gofiber/contrib/jwt v1.0.10` jwt fiber 支持
4. `github.com/gofiber/fiber/v2 v2.52.6`
5. `github.com/google/uuid v1.6.0` google uuid
6. github.com/jinzhu/copier v0.4.0
7. github.com/robfig/cron/v3 v3.0.0
8. github.com/spf13/viper v1.20.1
9. gorm.io/driver/mysql v1.5.7
10. gorm.io/gorm v1.25.12
11. resty.dev/v3 v3.0.0-beta.2

## 部署文档 🚀

1. `go mod tidy` 先整理并安装所需依赖，项目根目录下
2. 了解 scripts 目录下的执行脚本
   - `dev.sh` 启动 dev 模式，使用 gowatch 热重载，确保你拥有 gowatch 全局安装并添加 `$GOBIN` 到 `$PATH` 中, `go install github.com/silenceper/gowatch@latest`
   - `build.sh` 编译项目
   - `run.sh` 启动项目

## 规范 📄

### 接口方法中变量词汇

| 变量名 | 描述                                                          |
| ------ | ------------------------------------------------------------- |
| d      | 前端传入的 json data                                          |
| mo     | 代表你从数据库查出来的数据库对象                              |
| result | gorm.DB 对象                                                  |
| xxx    | 是你最终返回的对象，可以是任何名称，比如 user, product, paper |
