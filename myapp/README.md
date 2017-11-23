# MyApp使用说明
## 1.myapp介绍
	    myapp是一个基于beego框架的简单的提供登录注册的web小程序，同时myapp也通过grpc实现了echo服务以及自动报时服务
        
## 2.相关页面功能介绍
       192.168.34.93:8080                    显示myapp初始界面
       192.168.34.93:8080/user/signup        显示注册界面，
       192.168.34.93:8080/user/signup_result 显示注册是成功还是失败界面
       192.168.34.93:8080/user/login         显示登陆界面
       192.168.34.93:8080/user/profile       登录成功后会显示个人信息，登录失败会显示‘登陆失败’
       
     
## 3.接口说明
	|  接口名称  |  接口地址  |  请求方式  |  输入  |  说明  |
	|  --------  |  --------  |  :--------:|  :----:  |  ----  |
	|  主页预览  |192.168.34.93:8080|GET|无| 提示myapp处于运行状态 |
	|  注册界面  |192.168.34.93:8080/user/signup|GET|无|显示注册界面|
	|  登陆界面  |192.168.34.93:8080/user/login|GET|无|显示登陆界面|
	|注册界面提交信息|192.168.34.52:8080/user/signup_result|POST|用户名(字符串)<br>密码(字符串)<br>个人简介(字符串，可选)|会提示注册是否成功|
	|登陆界面提交信息|192.168.34.52:8080/user/profile|POST|用户名(字符串)<br>密码(字符串)|登录成功后才会显示个人信息，失败会提示登陆失败|


## 4.部署方法
- 拉取CentOS镜像
- 编译myapp项目，生成可执行文件--myapp在项目目录下
- 编写Dockerfile
   * 复制myapp项目到CentOS容器中
   * 添加OEM和VER环境变量
- 执行Dockerfile
- 上传镜像到Ekos网站
- 创建应用
	* 使用上传的镜像
	* 映射所需端口(8080)
	* 启动命令为“./myapp -logtostderr=true”
- 创建负载均衡，将映射出来的端口进行关联

## 5.遇到的问题和解决方法
	1.部署myapp容器后，浏览器访问程序的所有节目都出错，提示在/目录找不到view文件夹
	  解决方案:在编写Dockerfile文件的时候，最后使用WORKDIR命令指定容器的初始运行没有了
    
## 6.通过kubelet命令将部署的应用副本数扩展为4
    先查看资源分配情况
	[root@node1 ~]# kubectl get po
    NAME                                    READY     STATUS    RESTARTS   AGE
    default-http-backend-54c4db8898-zqmhd   1/1       Running   0          19h
    lb-myapp-7f466f565d-tkn9w               1/1       Running   0          16h
    lb-mysql-567489767-jbslt                1/1       Running   0          19h
    myapp-1-5fb94b8796-tnk9v                1/1       Running   0          16h
    mysql-0                                 1/1       Running   0          19h

    
    使用scale指令
    [root@node1 ~]# kubectl scale --replicas=4 deployment/myapp
	deployment "myapp" scaled
    
    再次查看资源分配情况
    [root@node1 ~]# kubectl get po
    NAME                                    READY     STATUS    RESTARTS   AGE
    default-http-backend-54c4db8898-zqmhd   1/1       Running   0          23h
    lb-myapp-7f466f565d-tkn9w               1/1       Running   0          20h
    lb-mysql-567489767-jbslt                1/1       Running   0          23h
    myapp-1-5fb94b8796-62rvh                1/1       Running   0          9m
    myapp-1-5fb94b8796-cck92                1/1       Running   0          5m
    myapp-1-5fb94b8796-jswkr                1/1       Running   0          9m
    myapp-1-5fb94b8796-tnk9v                1/1       Running   0          20h
    mysql-0                                 1/1       Running   0          23h
    结果显示myapp-1产生了4个副本


## 7.查看pod信息
	[root@node1 ~]# kubectl get po -o wide
    NAME                                    READY     STATUS    RESTARTS   AGE       IP              NODE
    default-http-backend-54c4db8898-zqmhd   1/1       Running   1          1d        10.233.71.33    node3
    lb-myapp-7f466f565d-tkn9w               1/1       Running   1          1d        192.168.34.93   node3
    lb-mysql-567489767-jbslt                1/1       Running   1          1d        192.168.34.92   node2
    myapp-1-5fb94b8796-62rvh                1/1       Running   1          4h        10.233.71.38    node3
    myapp-1-5fb94b8796-cck92                1/1       Running   1          4h        10.233.71.30    node3
    myapp-1-5fb94b8796-jswkr                1/1       Running   1          4h        10.233.75.40    node2
    myapp-1-5fb94b8796-tnk9v                1/1       Running   1          1d        10.233.75.38    node2
    mysql-0                                 1/1       Running   1          1d    

## 8.查看pod输出的日志
    #命令格式:kubectl logs -f <pod name>
    [root@node1 ~]# kubectl logs -f myapp-1-5fb94b8796-62rvh
    2017/11/22 06:15:35 [I] [asm_amd64.s:2197] http server Running on http://:8080
    2017/11/22 06:19:30 [D] [server.go:2568] |  192.168.34.93| 200 |    941.246µs|   match| GET      /user/profile   r:/user/profile
    这是访问web首页所产生的日志
