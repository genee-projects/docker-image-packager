# docker-images-transfer

## 描述 

一个用于把 docker images 进行转移的小工具

## 文件说明

* transfer.go
  * 用来并发请求进行 docker image 存储为 tar 文件的工具

* images.cfg
  * 用于配置需要存储为 tar 文件的 docker image 的文件, 需要增加 tag

* build.sh
  * 用于构建 transfer 工具的脚本(一次编译, 各处执行)
  
* copy/load.sh
  * 对当前目录下所有 tar 文件尝试进行 docker load 导入 docker image
  
* clean.sh
  * 一次使用后清除工具(删除所有 copy 中的 tar 文件)
  
