# docker-images-transfer

## 描述 

一个用于把 docker images 进行打包、解包的小工具

## 文件说明

* packager.go
  * 用来并发请求进行 docker image 存储为 tar 文件的工具

* images.cfg
  * 用于配置需要存储为 tar 文件的 docker image 的文件, 需要增加 tag

* build.sh
  * 用于构建 transfer 工具的脚本(一次编译, 各处执行)
  
* copy/load.sh
  * 对当前目录下所有 tar 文件尝试进行 docker load 导入 docker image
  
* clean.sh
  * 一次使用后清除工具(删除所有 copy 中的 tar 文件)
  

## 使用规范

1. 执行 build.sh 进行 transfer 编译
2. 修改 images.cfg 明确转移的 image
3. 执行 packager 文件, 将需要导出的 docker image 存储为 tar 文件(存储目录 copy)
4. 复制 copy 文件内所有内容到需要导入 docker image 的机器上
5. 执行 load.sh 对 tar 文件导入到机器上
6. 本地环境执行 clean.sh 对导出的 tar 文件进行删除

## 协议

MIT
