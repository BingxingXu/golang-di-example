## Feature
1. 使用[queryset](https://github.com/jirfag/go-queryset)自动生成type safe的gorm代码. (gorm v2新feature 也会支持)
2. 使用wire依赖注入code gen避免reflection的性能损耗
3. module -> controller -> dto -> service -> repository -> entity -> database 的项目结构划分
4. 使用 go mod

## How to Start
1. go mod tidy 安装项目依赖
2. 安装全局依赖 
    wire: 
    ```bash
    go get github.com/google/wire/cmd/wire
    ```
    queryset: 
    ```bash
    go get -u github.com/jirfag/go-queryset/cmd/goqueryset
    ```
2. go generate ./... 生成代码
    autogenerated_XX.go -> typesafe orm methods
    wire_gen.go -> DI Code

## Create New
1. 安装code gen工具
```bash
go get -u github.com/hexdigest/gowrap/cmd/gowrap
```
2. 代码模板在templates文件夹中
```bash
```
- TODO: gen Code with gowrap

## Cache
1. Router level Cache

## Limitation

### Other
inspired by [Blog Post](https://hellokoding.com/crud-restful-apis-with-go-modules-wire-gin-gorm-and-mysql/)
