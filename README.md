#Omega Haproxy Controller(A.K.A. HAServer) REST API


## Introduction

omega-haproxyctl is haproxy configration controller, it keep sry's service
config can sync update asap.

## Build and Run

Currently project use gvt to manage package dependency.

gvt restore to restore the dependency library.

```bash

# bash build.sh

# bash run.sh
```

## API List
  - [GET http://localhost:5004/api/status](#healthCheck)  :healthCheck  检查服务是否正常运行
  - [PUT http://localhost:5004/api/haproxy](#get master mertrics) 获取ID 为`clusterID`的集群的集群资源消耗信息

#### GET `http://localhost:5004/api/status`
检查服务是否正常运行 (healthCheck)   </br>
***Http Code***
`
500 or 200
`
500表示haproxy模板有问题，200表示haproxy模板正常．


#### PUT `http://localhost:5004/api/haproxy`
检查并更新haproxy的模板   </br>
***Http Code***
`
500 or 200
`
500表示更新haproxy模板失败，200表示更新haproxy模板成功．

## Reference list:
1.
http://engineeringblog.yelp.com/2015/04/true-zero-downtime-haproxy-reloads.html
