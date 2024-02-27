## 部署一个http服务进行测试

1. 首先准备一个简单的http服务，这里使用随便写一个,具体代码在 `web` 目录中
2. 构建镜像，把自己的服务打包成镜像
    通过以下命令构建镜像
    ```shell
    cd http-test/web
    ## 编译linux平台的代码
    GOOS=linux GOARCH=amd64 go build main.go
    ## 构建镜像,注意这里的镜像名字和版本要改成自己的
    docker build -t lstink/http-test:v1.0.0 .
    ```
3. 上传镜像到 `dockerhub` 或者 `阿里云` 之类的镜像仓库(可选)
    
    ```shell
    docker push lstink/http-test:v1.0.0
    ```
   > 注意，这里需要登录到镜像仓库 `docker login`，然后再上传镜像,如果是登录阿里云镜像仓库，则命令如下：
   登录阿里云镜像仓库
   docker login --username=lstink registry.cn-hangzhou.aliyuncs.com
   打包镜像
   docker tag [ImageId] registry.cn-hangzhou.aliyuncs.com/lstink/http-test:v1.0.0
   上传镜像
   docker push registry.cn-hangzhou.aliyuncs.com/lstink/http-test:v1.0.0
4. 部署服务
    ```shell
    kubectl apply -f deployment.yaml
    kubectl apply -f http-test.yaml
    ```
5. 查看服务
    ```shell
   kubectl get svc,pods | grep http-test-service
    ```
    > 这里记录下 `http-test-service` 的 `ClusterIP`，后面会用到
6. 通过ip在浏览器直接访问服务
   例如的到的ip是 `192.168.194.179`，则访问 `http://192.168.194.179`
7. 本地测试，hosts文件中增加域名解析
   `192.168.194.179 http-test.com`,这个域名是 http.yaml 中设置的域名，这时候就可以通过域名访问你的服务了

## 滚动升级服务

1. 修改服务代码
    修改 `web/main.go` 中的代码，例如修改返回的内容
    ```go
    // 响应显示的版本号 把这一行更改为 "var version = "v1.0.1"
	  var version = "v1.0.0" 
    ```
2. 重新构建镜像
    ```shell
    cd http-test/web
    ## 编译linux平台的代码
    GOOS=linux GOARCH=amd64 go build main.go
    ## 构建镜像,注意这里的镜像名字和版本要改成自己的
    docker build -t lstink/http-test:v2.0.0 .
    ```
3. 上传镜像到 `dockerhub` 或者 `阿里云` 之类的镜像仓库(可选)
    
    ```shell
    docker push lstink/http-test:v2.0.0
    ```
   > 如果是登录阿里云镜像仓库，则命令如下：
   打包镜像
   docker tag [ImageId] registry.cn-hangzhou.aliyuncs.com/lstink/http-test:v2.0.0
   上传镜像
   docker push registry.cn-hangzhou.aliyuncs.com/lstink/http-test:v2.0.0
4. 修改 `deployment.yaml` 中的镜像版本
    ```yaml
    spec:
      containers:
      - name: http-test
        image: lstink/http-test:v2.0.0
    ```
5. 执行升级
    ```shell
    kubectl apply -f deployment.yaml
    ```

6. 监视更新进度
   ```shell
   kubectl rollout status deployment http-test
   ```
7. 这将显示更新的当前状态。你还可以使用以下命令查看滚动更新的历史：

   ```shell
   kubectl rollout history deployment http-test
   ```
8. 如果需要回滚到先前的版本，可以使用 kubectl rollout undo 命令。

   ```shell
   kubectl rollout undo deployment http-test
   ```
   这将回滚 Deployment 到上一个成功的版本。请注意，回滚操作也是逐步进行的，以确保服务的可用性。

## 其他常用命令

```shell
# 查看所有的服务
kubectl get svc,pods
# 查看所有的命名空间
kubectl get ns
# 查看某个命名空间的服务
kubectl get svc,pods -n [namespace]
# 查看某个服务的详细信息
kubectl describe svc [service-name]
# 查看某个服务的日志
kubectl logs [pod-name]
# 扩容服务
kubectl scale deployment [deployment-name] --replicas=3 # 数字指的是扩容后 pod 的数量，可以根据实际情况修改

```

## 全部测试完成，删除服务
```shell
kubectl delete -f deployment.yaml
kubectl delete -f http.yaml
```