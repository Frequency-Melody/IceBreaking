# 助手破冰认人脸系统后端

## 一、开发计划
1. 数据库设计
2. 增删改查接口，除图片、鉴权。
3. 实现图片上传，利用阿里云 OSS
4. 实现鉴权 oauth

## 二、数据库设计
### 2.1 设计思想
* 学生表，只存储个人信息，以及是否隐藏图片  
* 图片表，存储图片在 `OSS` 中的外链地址
* 学生与图片关联表，存储图片 id 与 学生 id  
* 学生表 mini 版：只存放学生的 id 以及是否隐藏图片，为了随机照片的时候查表快。（其实「隐藏图片」这个属性出现两次了，可以把学生表里的那个删了）  
### 2.2 数据库定义
```go
type Student struct {
    gorm.Model
    Name       string `gorm:"unique" binding:"required"`
    Department string `gorm:"comment:部门" binding:"required"`
    HidePic    bool   `gorm:"comment:是否隐藏照片"`
}

type Picture struct {
    gorm.Model
    Url string `gorm:"comment:图片在阿里云 OSS 中的地址"`
    //Student   Student
    //StudentId int
}

// 学生与照片的关联表
// 其实一对一关系，把 StuId 放在 Picture 字段里也行
// 但是为了防止传给前端的时候，被前端知道 StuId 不太好
type RelationStudentPic struct {
    gorm.Model
    StudentId int
    PictureId int
}

// 一张表只存学生 id，这样能快速检索到有哪些学生
// 同时必须复制 HidePic 字段，否则可能查出的内容无效
type StudentId struct {
    gorm.Model
    StudentId int
    HidePic   bool
}
```
## 三、增删改查接口
### 3.1 随机抽取功能的实现
这里最难的应该是，如果实现，返回一个人的照片，但同时返回多个学生信息。并且，返回的人中必须有一个是对的，剩下的是混淆，且返回的人都不能重复？

如果分段取，先取一个人，再随机其他的，会导致重复问题。

所以我采取的办法是，直接取 n 个人（包含混淆和正确的），然后 n 个人随机选取一个，提取图片，再一起返回。

所以问题就分解成，如何随机从数据库中抽取 n 个学生。

为了速度，我设计了个 mini 表（StudentId），只存放学生 id 以及是否「隐藏图片」。每次查询，得到这张表中不隐藏图片的所有的学生 id。

然后对这个 id 数组随机抽取 n 个数。过程简单描述一下就是，从 [0,id总数)，中生成不重复的 n 个数字，即 id 的下标。 

然后通过下标取得 id，再通过 图片学生关联表 ，利用学生 id 找到学生对应的图片。

最后把 n 个学生信息和一张图片信息返回。
### 3.2 图片上传的实现