# Fabric中使用国密动态库和静态库

## 使用动态库

### 修改Makefile相关

由于Fabric在编译时默认添加了静态编译选项，所以直接连接动态库会报连接错误，首先将静态编译选项关闭，编辑fabric默认路径下的`docker-env.mk`文件

```
将
DOCKER_GO_LDFLAGS += -linkmode external -extldflags '-static -lpthread'
修改为
DOCKER_GO_LDFLAGS += -linkmode external -extldflags '-lpthread'
```

### 丰富baseos

由于baseos镜像是peer,orderer等二进制程序最后运行的环境，而baseos由于镜像本身过于精简而缺少了一些运行时库，所以要将二进制程序运行时依赖的动态库打包进去，这些库从baseimage中获取

运行一个baseimage容器

```
docker run -it hyperledger/fabric-baseimage:x86_64-0.3.2 bash
```

进到容器里面后，拷贝lib库到用户目录，否则无法直接拷贝出来

```
cp -r /usr/lib/x86_64-linux-gnu ~
```

在宿主机的终端将容器中的lib库拷贝出来

```
docker cp ${CONTAINER_ID}:/root/x86_64-linux-gnu .
```

备份原版baseos

```
docker tag hyperledger/fabric-baseos:x86_64-0.3.2 hyperledger/fabric-baseos:x86_64-0.3.2-ori
docker rmi hyperledger/fabric-baseos:x86_64-0.3.2
```

准备build镜像环境

```
mkdir newbaseos # 新baseos构造目录
cp -r x86_64-linux-gnu newbaseos # baseimage中的动态库
cp -r lib newbaseos # 国密lib库
cp LICENSE newbaseos # 国密运行时证书
cd newbaseos
```

编写Dockerfile

```
FROM hyperledger/fabric-baseos:x86_64-0.3.2-ori
ENV LD_LIBRARY_PATH=/usr/local/lib
ENV OPENSSL_ENGINES=/usr/local/lib/engines-1.1
COPY lib /usr/local/lib/
COPY x86_64-linux-gnu /usr/lib/x86_64-linux-gnu
COPY LICENSE /usr/local/bin
```

在newbaseos目录总编译新镜像

```
# ls
Dockerfile  lib  LICENSE  x86_64-linux-gnu
# docker build -t hyperledger/fabric-baseos:x86_64-0.3.2 .
```
编译本地二进制,注意gm-crypto包中cgo编译选项中的路径为绝对路径

```
// #cgo LDFLAGS: -L /opt/gopath/src/github.com/hyperledger/fabric/vendor/github.com/peersafe/gm-crypto/usr/lib -lciphersuite_crypto -lciphersuite_smengine
// #cgo CFLAGS: -I /opt/gopath/src/github.com/hyperledger/fabric/vendor/github.com/peersafe/gm-crypto/usr/include


make native && make release && make docker
```

运行前，修改docker-compose.yaml，由于国密库需要LICENSE才能正常运行，而且LICENSE需要可执行文件目录下

在本地运行时还需要加上环境变量

```
export GM_CRYPTO=/opt/gopath/src/github.com/hyperledger/fabric/vendor/github.com/peersafe/gm-crypto
export LD_LIBRARY_PATH=$GM_CRYPTO/usr/lib
export OPENSSL_ENGINES=$GM_CRYPTO/usr/lib/engines-1.1
```

### 修改ccenv镜像

由于shim包中引用到了bccsp包，进而关联了gm-crypto包，所以后面在编译chaincode镜像的时候需要有国密的编译环境，所以需要重新打包制作ccenv镜像

#### 修改chaincode编译参数

core/chaincode/platforms/golang/platform.go  GenerateDockerBuild()

```
const ldflags = "-linkmode external -extldflags '-static'"
修改为
const ldflags = "-linkmode external"
```
#### 重新打包ccenv镜像

之所以需要重新打包是因为chaincode中会用到

编写Dockerfile

```
FROM hyperledger/fabric-ccenv:v1.0.4-ori
COPY gm-crypto /opt/gopath/src/github.com/hyperledger/fabric/vendor/github.com/peersafe/gm-crypto
```

执行

```
# docker build -t hyperledger/fabric-ccenv:latest .
```

## 使用静态库

使用静态库只需要修改一下baseos和baseimage镜像就可以

### 修改baseos

只需要把LICENSE文件拷贝到原来baseos镜像中的`/usr/local/bin`目录下

* 备份原版baseos

```
docker tag hyperledger/fabric-baseos:x86_64-0.3.2 hyperledger/fabric-baseos:x86_64-0.3.2-ori
docker rmi hyperledger/fabric-baseos:x86_64-0.3.2
```
* 编写Dockerfile,`LICENSE`文件为国密库依赖的许可证

```
FROM hyperledger/fabric-baseos:x86_64-0.3.2-ori
COPY LICENSE /usr/local/bin
```

* 构建新的baseos

```
# ls
Dockerfile  LICENSE 
# docker build -t hyperledger/fabric-baseos:x86_64-0.3.2 .
```

### 修改baseimage

因为编译chaincode时是在ccenv,ccenv是基于baseimage编译，所以需要将国密的库打包到baseimage镜像中

* 备份原版baseimage

```
docker tag hyperledger/fabric-baseimage:x86_64-0.3.2 hyperledger/fabric-baseimage:x86_64-0.3.2-ori
docker rmi hyperledger/fabric-baseimage:x86_64-0.3.2
```

* 编写Dockerfile,gm-crypto为国密静态库的目录,需要从fabric的vendor目录中拷贝到当前构建镜像的目录

```
FROM hyperledger/fabric-baseimage:x86_64-0.3.2-ori
COPY gm-crypto /opt/gopath/src/github.com/hyperledger/fabric/vendor/github.com/peersafe/gm-crypto
```

* 构建新的baseimage

```
# ls
Dockerfile  gm-crypto 
# docker build -t hyperledger/fabric-baseimage:x86_64-0.3.2 .
```

* 编译完这两个镜像之后，直接执行

```
# make native && make docker && make release
```