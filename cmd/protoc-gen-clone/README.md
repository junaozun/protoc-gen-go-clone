# 通过proto文件生产接口代码示例

## 安装插件
`go get gitlab.uuzu.com/war/pbtool/cmd/protoc-gen-clone`

## 使用方法
`protoc --proto_path=../../testdata --proto_path=. --clone_out=paths=source_relative:. testdata.proto`

## Example
[Example](example)
