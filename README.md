# Image-Processing-Tool-with-Go-AI
本项目是一个AI图像生成网站，包含后端、前端以及图像生成API三部分，分别使用Golang、Vue.js和Python实现。

教程地址：[https://blog.csdn.net/qq_42638506/category_12833718.html](https://blog.csdn.net/qq_42638506/category_12833718.html)

## 项目结构
```bash
├── ImageGenAPI                        # 图像生成API代码
│   └── instruct-pix2pix.py            # 图像生成API主文件
├── api-performance-testing-with-k6    # 项目测试代码
│   ├── my-test                        # 测试代码文件夹
│   │   ├── my_test.js                 # 测试代码主文件
│   │   ├── utils.js                   # 测试代码需要的工具函数
│   │   └── ...                        # 其他文件
│   └── ...                 
├── backend                            # 后端代码
│   ├── main.go                        # 后端入口文件
│   └── ...                  
├── frontend                           # 前端代码
│   ├── ...                            # Vue.js项目
└── README.md                          # 项目说明文件
```
## 环境依赖

* Python 3.8+ 和 uvicorn
* Golang 1.23.3
* Node.js 16+ 和 npm 8+
* k6 0.55.0

## 启动步骤

按照以下步骤依次启动项目：
### 1. 启动图像生成API

进入 `ImageGenAPI` 目录，并运行以下命令启动图像生成API：
```bash
uvicorn instruct-pix2pix:app --host 0.0.0.0 --port 8001
```
API服务将运行在 [http://0.0.0.0:8001](http://0.0.0.0:8001)。

### 2. 启动后端服务

进入 `backend` 目录，并运行以下命令启动后端服务：
```bash
go run main.go
```

后端服务将运行在 [http://localhost:8000](http://localhost:8000)。

### 3. 启动前端服务

进入 `frontend` 目录，并运行以下命令启动前端服务：
```bash
npm install
npm run serve
```
前端服务将运行在 [http://localhost:8080](http://localhost:8080)。

### 4. 测试代码运行

进入 `api-performance-testing-with-k6/my-test` 文件夹，运行以下命令（`--out cloud` 参数可选）：
```bash
k6 run --out cloud my_test.js
```
加上 `--out cloud` 参数，测试结果将上传到可视化界面，点击测试完成后出现的链接，即可看到可视化结果。

## 功能简介

1. **用户注册与登录：** 支持用户注册、登录和Token刷新。
2. **作品管理：** 支持创建作品分类、添加作品，并按分类查看作品。
3. **图像生成：** 基于上传的图片和描述调用AI生成新图像。

## API说明

### 图像生成API

图像生成API基于 `instruct-pix2pix` 模型和 **FastAPI** 实现，用于处理输入图像并生成AI生成的图像。API默认运行在8001端口。

### 后端API

后端提供以下功能的API：

* 用户认证（注册、登录、Token刷新）
* 分类管理（创建分类、查询分类、删除分类）
* 作品管理（上传作品、查询作品、删除作品）
* 调用图像生成API生成新图像

## 注意事项

1. **前后端及图像生成API的端口需一致：** 确保前端代码中的后端和图像生成API地址配置正确。
2. **依赖安装：** 首次运行需要安装相应依赖。

## 常见问题

1. 图像生成API无法启动：确保安装了uvicorn，并按需检查instruct-pix2pix相关依赖。
2. 前端无法访问后端或API：检查前端配置文件中的API地址是否正确。

## 更新日志
== **2024 年 12 月 17 日** ==
1. 更新图像上传接口调用的API，`UploadImageToSMMS(base64Image string)` --> `UploadImageToOSS(base64Image string)`，API 实现教程：[https://blog.csdn.net/qq_42638506/article/details/144577748](https://blog.csdn.net/qq_42638506/article/details/144577748)
2. 添加 **k6** 测试代码文件夹 `api-performance-testing-with-k6`
3. 更新支持并发版本代码：[https://github.com/Gezelligheid1010/Image-Processing-Tool-with-Go-AI-v2](https://github.com/Gezelligheid1010/Image-Processing-Tool-with-Go-AI-v2)

## 相关网址
1. instruct-pix2pix 模型网址：[https://huggingface.co/timbrooks/instruct-pix2pix](https://huggingface.co/timbrooks/instruct-pix2pix)
