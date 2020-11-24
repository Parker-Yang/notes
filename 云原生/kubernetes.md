# Kubernetes

---

### Pod ：k8s里可以部署的最小单元

- 为啥是pod而不是container
  - 为了更好的管理container，获取container信息；例如重启策略
- pod中的资源共用，Pod中的容器共享IP地址和端口号，它们之间可以通过localhost互相发现