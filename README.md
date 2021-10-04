# 杭电助手破冰认人脸系统-后端
## 一、功能简介
* 登   录：接入杭电助手 `OAuth2` ，登录智慧杭电账号。  
* 照片上传：用户可上传自己的照片信息，并决定是否参与认人脸游戏。  
* 随机人脸：每次向前端随机返回 `n` 个学生信息，与一张照片，该照片属于 `n` 个学生之一。  
* 选项校验：用户在前端从 `n` 个选项中选择一个，前端通过学生与照片的 `uuid` 向后端校验正确与否。

## 二、环境与配置
### 2.1 编译环境
go >= 1.16
### 2.2 配置信息
非敏感配置信息在 `/config/config.toml` 中，请直接阅读该文件注释  
### 2.3 环境变量
敏感配置信息请配置环境变量，如下：
```yaml
# mysql 域名（IP）与密码
MYSQL_HOST
MYSQL_PWD

# 阿里云 OSS 相关
OSS_ENDPOINT
OSS_ACCESS_KEY_ID
OSS_ACCESS_KEY_SECRET

# 杭电助手鉴权相关
HDUHELP_CLIENT_ID
HDUHELP_CLIENT_SECRET
```

## 三、API 文档
### 3.1 总体
~~base URL : http://HOST:8091/~~
建议部署时使用反向代理
base URL ：https://tcualhp.cn/ice/v2.0/

除登录外，所有请求都需要带上鉴权请求头，若 `token` 值为 `abcd12345`，则请求头为:
```yaml
Authorization: token abcd12345
```

总体返回结构

```json
{
  "Error": 20000,// Number, 当且仅当数值为 20000 时无错
  "Msg": "OK",  // string，错误提示
  "Data": ,// Object，当 Error 为 20000 时，为数据；否则为空
  "Redirect": "" // 重定向信息，目前未用到，一直为空
}
```
### 3.2 用户登录与鉴权部分
#### 3.2.1 `GET` /auth/login 登录
无 `Query Param`，后端将页面重定向至助手授权页，授权结束后，将页面重定向回 `/config/config.toml` 中配置的 `FrontEnd.Home` 页。  
并把 `token` 与 学生姓名 `name` 信息，附加到 `Query` 中。前端应及时取出保存。  

#### 3.2.2 `GET` /auth/validate 验证 token 是否有效
用户每次打开主页时，前端都应发送此请求，验证 `token` 是否有效  
无 `Query Param` 参数，`token` 附在请求头中。
若成功，`Error` 为 `20000`。  

### 3.3 图片上传、随机人脸、图片验证
本部分所有请求都需要携带 `Authorization` 请求头，若请求头无效，将被重定向至 `/config/config.toml` 中配置的 `FrontEnd.Home` 页

#### 3.3.1 `POST` /picture/upload 图片上传
使用 `form-data` 上传图片， 格式：`picture`: 图片文件

重复上传可覆盖原有照片。  

Response Body 中的 `Data`：
```json
{
  "Data": {
    "url": "图片的公网访问 url "
  }
}
```
PostMan 示例：

![image-20211002141209549](https://bird-notes.oss-cn-hangzhou.aliyuncs.com/img/image-20211002141209549.png)

#### 3.3.2 `GET` /student/rand 随机人脸与学生姓名选项
Query Param:
* num ：数字。含义为选项的数量，需大于等于2。

Response Body 中的 Data：
```json
{
"Data": {
  "Picture": {
    "Uuid": "814991e5-0e1c-e423-4755-faa68d10791a",
    "Url": "https://***********hangzhou.aliyuncs.com/TBO%405%252Q%24NOU~5707_JGX52.png"
  },
    "Students": [
      {
        "Uuid": "f812535c-6cb9-62ce-32c8-3e5260ae7e88",
        "Name": "姓名1"
      },
      {
        "Uuid": "dbc272b1-dfed-3ee2-f2cb-d3d2af8b304f",
        "Name": "姓名2"
      }
    ]
}}
```
PostMan 示例：

![image-20211002142329876](https://bird-notes.oss-cn-hangzhou.aliyuncs.com/img/image-20211002142329876.png)

#### 3.3.3 `GET` /picture/verify 验证照片是否属于某学生
传入照片和学生的 `uuid`,验证是否匹配，并返回该照片对应的正确学生的名字。  
无论正确与否， `Error` 都是 `20000`，请根据返回体中的 `verify` 字段判断选项是否正确。  
Query Param：
* pictureUuid: 随机时得到的照片的 uuid
* studentUuid：用户选择的学生对应的 uuid

Response Body 中的 Data：
```json
{
  "Data": {
    "studentInfo": {
      "Uuid": "e7abc983-cb50-0d78-844a-8d31c55b7b85",
      "Name": "照片对应的正确的学生的姓名"
    },
    "verify": false
  }
}
```

#### 3.4 参与/退出 认人脸游戏
若参加，则 隐藏人脸 字段 `hide_pic` 为 `false`
~~还没写，目前默认全部参加，想退出？不可能的！不给你机会（接口）~~
### 3.4.1 `GET` /student/status 获取是否参加状态
无 Query Param
Response Body 中的 Data：
```json
{
  "Data": {
    "hide_pic": false
  }
}
```

### 3.4.2 `POST` /student/status 修改是否参加状态
payload:
```json
{
  "hide_pic": false // false 为不隐藏，参加游戏
}
```