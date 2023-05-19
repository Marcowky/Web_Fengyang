# Web_Fengyang

# Quick Start

*在进行部署前，请务必保证安装好以下环境*

> `Go` / `Node.js` / `Vue.js` / `MySQL`

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

然后修改后端中的数据库配置`Web_Fengyang\Web_Fengyang_Server\config\config.json`

将数据库名称修改为你创建的数据库

## 3. 安装vite前端工具

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
- `MySQL`数据库
- `jwt`包：用于生成和验证`token`
- `uuid`包：用于生成唯一`id`

## Client前端

- `vue3`：`Vue3`是一个流行的前端 JavaScript 框架
- `node`：用于构建可扩展的网络应用程序
- `vite`：用于快速地开发`Vue.js`单页应用`SPA`和库的工具
- `axios`：由于在前端方便地进行HTTP请求
- `pinia`：一个简单、轻量级的状态管理库
- `sass`：一种CSS预处理器
- `vue-router`：`Vue.js`官方的路由管理器
- `naive-ui`：`Vue3`的桌面端UI组件库
- `wangeditor`：一款基于`web`的富文本编辑器

