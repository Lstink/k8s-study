## dashboard 构建指南

### 1. 通过 `helm` 安装 `metrics-server`，并且创建 `metrics-server` 命名空间
```shell

通过 `helm` 安装 `kubernetes-dashboard`，并且创建 `kubernetes-dashboard` 命名空间
```shell
helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/
helm upgrade --install kubernetes-dashboard kubernetes-dashboard/kubernetes-dashboard --create-namespace --namespace kubernetes-dashboard
```
卸载
```shell
helm delete kubernetes-dashboard --namespace kubernetes-dashboard
```

### 2. 通过 `kubectl` 创建 `dashboard-admin` 服务账号，并且绑定 `cluster-admin` 角色
```shell
kubectl apply -f admin-user.yaml
```

### 3. 通过 `kubectl` 获取 `dashboard-admin` 服务账号的 `token`
```shell
kubectl -n kubernetes-dashboard create token admin-user
```