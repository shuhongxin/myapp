# Kubernetes kubectl 命令用例

一，get命令

get命令用于获取集群的一个或一些resource信息。使用--help查看详细信息。kubectl的帮助信息、示例相当详细，而且简单易懂。建议大家习惯使用帮助信息。kubectl可以列出集群所有resource的详细。resource包括集群节点、运行的pod，ReplicationController，service等。

二，describe命令

describe类似于get，同样用于获取resource的相关信息。不同的是，describe获得的是更详细的resource个性的详细信息，get获得的是resource集群相关的信息。describe命令同get类似，但是describe不支持-o选项，对于同一类型resource，describe输出的信息格式，内容域相同。
如果发现是查询某个resource的信息，使用get命令能够获取更加详尽的信息。但是如果想要查询某个resource的状态，如某个pod并不是在running状态，这时需要获取更详尽的状态信息时，就应该使用describe命令。

三，create命令

kubectl命令用于根据文件或输入创建集群resource。如果已经定义了相应resource的yaml或son文件，直接kubectl create -f filename即可创建文件内定义的resource。也可以直接只用子命令[namespace/secret/configmap/serviceaccount]等直接创建相应的resource。从追踪和维护的角度出发，建议使用json或yaml的方式定义资源。

四，delete命令

根据resource名或label删除resource。

五，logs命令

logs命令用于显示pod运行中，容器内程序输出到标准输出的内容。跟docker的logs命令类似。如果要获得tail -f 的方式，也可以使用-f选项。

六，exec命令

 exec命令同样类似于docker的exec命令，为在一个已经运行的容器中执行一条shell命令，如果一个pod容器中，有多个容器，需要使用-c选项指定容器。

七，apply命令

apply命令提供了比patch，edit等更严格的更新resource的方式。通过apply，用户可以将resource的configuration使用source control的方式维护在版本库中。每次有更新时，将配置文件push到server，然后使用kubectl apply将更新应用到resource。kubernetes会在引用更新前将当前配置文件中的配置同已经应用的配置做比较，并只更新更改的部分，而不会主动更改任何用户未指定的部分。apply命令的使用方式同replace相同，不同的是，apply不会删除原有resource，然后创建新的。apply直接在原有resource的基础上进行更新。同时kubectl apply还会resource中添加一条注释，标记当前的apply。类似于git操作。