[TOC]

## Running-Spark-on-Kubernetes

### 1. 安全性 security

默认情况下，Spark 中的安全性是关闭的。这可能意味着默认情况下您很容易受到攻击。在运行 Spark 之前，请参阅[Spark 安全性](https://spark.apache.org/docs/latest/security.html)和下面的具体建议。

### 2. 用户身份 user identity

从项目提供的 Dockerfiles 构建的图像包含一个默认[`USER`](https://docs.docker.com/engine/reference/builder/#user)指令，默认 UID 为`185`. 这意味着生成的图像将在容器内作为此 UID 运行 Spark 进程。具有安全意识的部署应考虑提供带有`USER`指令的自定义映像，指定其所需的非特权 UID 和 GID。生成的 UID 应在其补充组中包含根组，以便能够运行 Spark 可执行文件。`docker-image-tool.sh`使用提供的脚本构建自己的图像的用户可以使用该`-u <uid>`选项指定所需的 UID。

或者，[Pod Template](https://spark.apache.org/docs/latest/running-on-kubernetes.html#pod-template)功能，可用于添加[安全上下文](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#volumes-and-file-systems)用`runAsUser`的豆荚火花的提交。这可用于覆盖`USER`图像本身中的指令。请记住，这需要您的用户合作，因此可能不是共享环境的合适解决方案。如果集群管理员希望限制 Pod 可以作为[哪些](https://kubernetes.io/docs/concepts/policy/pod-security-policy/#users-and-groups)用户运行，他们应该使用[Pod 安全策略](https://kubernetes.io/docs/concepts/policy/pod-security-policy/#users-and-groups)。
