# 环境配置

假设代码的根是 ~/qiniu（也就是本库在 qiniu/gobook 目录下），则环境配置步骤如下：

1. 安装 go1
2. 打开 ~/.bashrc，加入： 

    export QINIUROOT=~/qiniu

    source $QINIUROOT/gobook/env.sh

3. 保存 ~/.bashrc ，并 source 之


# 运行代码

对于单文件程序，如sample.go，直接运行go run sample.go即可

对于project，依次install所依赖的包，最后编译主程序文件，具体构建与执行方法见书


# 介绍

《GO语言编程》一书中示例源代码

