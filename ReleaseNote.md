Release Note
=============

这个文件将会记录:

- 下次release运维需要更新的配置
- 重大的bug fix
- 重大的feature


##Pending release changes

##更新版本 v2.4.160720

### 功能改进
  * 使用 qdisc queue 取代 iptables SYN DROP，实现无中断 Reload Haproxy，用户体验更好
  * 增加local-link 169.254.255.254，利用 Yelp 提供的 Yocalhost 功能，解决了容器内应用直接访问外部服务端口，不需要指定容器外的 host IP
  * 采用 s6 init 取代 supervisord
  * HAServer 改名为 omega-haproxyctl

### 依赖支持
  * 需要agent 给每台容器启动时docker run加载一个 host ip 到/etc/hosts
--add-host="" : Add a line to /etc/hosts (host:IP)。
```
docker run -it --add-host hostlocal.io:169.254.255.254 omega-slave cat /etc/hosts
```
对于离线应用，加上 hosts，保障解析local ip没问题



