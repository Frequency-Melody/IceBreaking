# API_Reference_v1.0

## 一、总体

baseUrl: `https://tcualhp.cn/ice/v1.0/`

返回体结构

```json
{
    "Error": 20000, //数字，只有 20000 才是正确返回，其余均有错
    "Msg": "OK", // 正确为 OK，其余时候为报错信息
    "Data": , // 数据，当无错时返回数据，有错时为空
    "Redirect": , // 网页跳转信息，目前用不到
}
```

## 二、接口详解

### 2.1 添加学生 `POST` /student/add

传入名字、部门、学号、是否隐藏照片，返回 `uuid` 。判重标准为学号。

`payload` 和 `response body` 都是 `json`，看截图吧，应该好懂的。

![image-20210929195530107](https://bird-notes.oss-cn-hangzhou.aliyuncs.com/img/image-20210929195530107.png)

### 2.2 获取所有学生信息 `GET` /student/all

直接看图吧

![image-20210929201502899](https://bird-notes.oss-cn-hangzhou.aliyuncs.com/img/image-20210929201502899.png)

### 2.3 上传图片 `POST` /picture/upload

传入图片和学生的 `uuid`，上传学生图片，返回图片公网地址。同一 `uuid` 重复上传图片可覆盖。

请求参数是 `form-data` 格式

![image-20210929200936776](https://bird-notes.oss-cn-hangzhou.aliyuncs.com/img/image-20210929200936776.png)

### 2.4 获取随机图片与学生 `GET` /student/rand

传入随机学生的总数 `num`，返回一张学生照片，`num` 个学生信息。其中 `num` >= 2。

![image-20210929202232034](https://bird-notes.oss-cn-hangzhou.aliyuncs.com/img/image-20210929202232034.png)

### 2.5 图片验证 `GET` /picture/verify

传入学生与图片的 `uuid`（由 2.4 得，验证是否匹配）。无论是否正确，都将返回照片对应的学生信息。

![image-20210929202417149](https://bird-notes.oss-cn-hangzhou.aliyuncs.com/img/image-20210929202417149.png)