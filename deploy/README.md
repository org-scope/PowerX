# 本地打包PowerX的Docker镜像


```shell
# 按照你的环境和需求，可替换{xxx}里的值

# 进入你的项目文件根目录下
> cd {your_project_path}

# 编译项目中deploy/docker/Dockerfile
> docker build -t {powerx}:{latest} -f ./deploy/docker/Dockerfile .

# 编译完后的镜像，直接用docker可以跑起来
> docker run -p {8888}:{8888} -it {powerx-dashboard}:{latest}

```

