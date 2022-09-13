# 运行
```shell
# 修改mysql、redis配置
vi etc/core-api.yaml

git clone https://github.com/xuwuruoshui/cloud-disk.git
cd core
go mod tidy
go run core.go -f etc/core-api.yaml
```

# 数据库
- repository_pool: 文件信息
- share_basic: 分享
- user_basic: 用户信息
- user_repository: 用户仓库信息

# 功能实现
- 认证
    - 登录
    - 注册
        - 邮箱验证码发送
    - jwt双token
- 文件
    - 上传文件(腾讯cos)
    - 用户文件关联存储
    - 获取文件列表
    - 修改文件名
    - 创建文件夹
    - 删除文件/文件夹
    - 移动文件
- 分享
    - 创建分享文件
    - 分享文件
    - 保存分享文件
- 分片上传文件(腾讯cos)
    - 秒传/分片上传初始化
    - 文件分片上传
    - 分片上传完成

# 其他
- api文档
cloud-disk.openapi.json
- sql文件
cloud-disk.sql

