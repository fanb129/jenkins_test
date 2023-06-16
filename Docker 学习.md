# Docker 学习

## 2023-6-16

### Docker简介

<img src="C:\桌面\Docker_imgs\image-20230616091653448.png" alt="image-20230616091653448" style="zoom:67%;" />

### 重要概念

<img src="C:\桌面\Docker_imgs\image-20230616091637499.png" alt="image-20230616091637499" style="zoom:67%;" />

### 应用

<img src="C:\桌面\Docker_imgs\image-20230616091809056.png" alt="image-20230616091809056" style="zoom:67%;" />

> 梳理Docker技术

- 一个命令安装软件
- 打包自己的镜像，Docker hub
- 目录挂载，保存数据
- 多容器通信
- Docker-Compose，更多容器通信
- 备份、迁移数据

### 服务器安装Docker教程

[容器与云|如何在 Ubuntu 22.04 LTS 中安装 Docker 和 Docker Compose (linux.cn)](https://linux.cn/article-14871-1.html)

### Docker底层实现

<img src="C:\桌面\Docker_imgs\image-20230616152034353.png" alt="image-20230616152034353" style="zoom:67%;" />

### 个人理解

容器技术和虚拟机技术的区别在于，容器中不包含操作系统内核，在运行时和宿主机共用一个内核；虚拟机中包含操作系统内核，同样有占用硬件资源多的缺点。