# Web_Fengyang_项目开发

# 1_项目需求

## 1.1_项目内容

1. 使用Go语言进行后端代码编写、程序调试和测试、版本迭代
   - 框架：Gin、net/http
2. 使用HTML/JS/TS进行前端程序设计
   - 框架：vue
3. 掌握单元测试知识，使用go test进行单元测试
4. 掌握版本管理知识，用Git进行管理

## 1.2_成绩判定

|          | 比例 |            说明             |
| :------: | :--: | :-------------------------: |
| 个人得分 | 60%  | 个人贡献(80%)+助教评价(20%) |
| 项目得分 | 40%  |        项目报告+开源        |

- 个人贡献：组内打分
- 助教评价：助教打分
- 项目报告
  - 项目背景
  - 核心代码
  - 测试报告
- 原码上传Gitee

# 2_开发整体流程

1. 需求分析：4.25-4.30
2. 功能设计：5.1-5.7
3. 前后端分离编码：5.8-6.4
4. 功能迭代完善：6.4-7.5
5. 测试：7.6-7.15
6. 报告撰写：7.6-7.15

# 3_项目开发意义

*这一部分可以结合现实，谈谈对凤阳，对三下乡*

# 4_项目需求与功能

## 4.1_功能分析

*我们的网站具体包含哪些功能？功能是否能实现我们的项目意义？*

## 4.2_功能设计

*功能需要怎样实现？功能是否能细分成更小的功能*

## 4.3_数据库设计

*功能的实现需要哪些数据？需要哪些类？数据库需要怎样设计？*

## 4.4_接口设计

*功能需要哪些接口？前后端需要定义哪些接口？*

# 5_人员分工

# 6_项目具体编码

# 7_项目特色功能

# 8_项目测试

# 0_更新日志

*4.26*

- 上传了前端和后端的基本框架
- 定义了数据库的User类
- 实现了网站注册与登录功能
- 编写了网站注册与登录界面

# TODO

- 4.29 完成文章功能
  - 文章发布 4.29 ok
  - 文章查看 4.30 ok
  - 文章删除 4.30 ok
  - 文章分类 5.3 ok
- 4.29 完成视频功能
  - 视频发布
  - 视频查看
  - 视频删除

# Quick Start

## 1. 下载代码


```bash
git clone [url]
```

## 2. 创建数据库
   
```bash
# 进入mysql数据库
mysql -u root -p
# 输入密码
...
# 创建数据库
CREATE DATABASE <your database name>
```

然后修改后端中的数据库配置`D:\codefile\Project\Web\Web_Fengyang\Web_Fengyang_Server\common\database.go`

将数据库名称修改为你创建的数据库

## 3. 安装vite

```bash
cd Web_Fengyang_Client # 这里为client的目录
npm install vite
```

## 4. 启动项目

启动后端

```bash
cd Web_Fengyang_SServer
go run main.go
```

启动前端

```bash
cd Web_Fengyang_Client
npm run dev
```

# 技术栈

## Server后端

- `Golang`语言
- `Gin`框架：一个基于`golang`的简单`Web`微框架
- `Gorm`库：`Go`语言编写的`ORM`(`Object-Relational Mapping`)库，用于简化数据库操作
- `mysql`数据库
- `jwt`包：用于生成和验证`token`
- `uuid`包：用于生成唯一`id`

## Client前端

- `vue3`：`Vue3`是一个流行的前端 JavaScript 框架
- `node`：
- `vite`：用于快速地开发`Vue.js`单页应用`SPA`和库
- `axios`：
- `pinia`：
- `sass`：
- `vue-router`：
- `naive-ui`：
- `wangeditor`：

