<img src="https://lh5.googleusercontent.com/proxy/r5D7LX7gbvXfuJU1SFAfCM1SerPt0KcBvR_R0qpXO_fsa39nwCKhyGE0UQbFP99XpSMRuPWrckLRnkoU747FW6EHY1_Gqf1xzhXYhJnIqIHizuhbBX3fh0sgdxbpIwJrDtC9g-uELzM-xYNfiw=s0-d">

#### INSTAGRAM CRAWLER

[![build](https://img.shields.io/badge/build-success-green?style=flat&logo=github)](https://github.com/maengsanha/instacrawler/pulse)
[![license](https://img.shields.io/badge/license-MIT-blue)](https://github.com/maengsanha/instacrawler/blob/master/LICENSE)

<br>

**Installation**

```bash
$ go get github.com/maengsanha/instacrawler
```

<br>

**Build (Linux)**

```bash
$ make build_linux
```

<br>

**Build (Raspberry PI)**

```bash
$ make build_raspberry
```

<br>

**Install Docker (Ubuntu)**

```bash
$ curl -fsSL https://get.docker.com/ | sudo sh
$ docker --version

$ sudo systemctl enable docker
$ sudo systemctl status docker
$ sudo systemctl start docker
```

<br>

**Install Kubernetes (Ubuntu)**

```bash
$ sudo apt-get install curl
$ curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add
$ sudo apt-add-repository "deb http://apt.kubernetes.io/ kubernetes-xenial main"
$ sudo apt-get install kubeadm kubelet kubectl
$ sudo apt-mark hold kubeadm kubelet kubectl
$ kubeadm version

# assign unique hostname for each server node
$ sudo swapoff -a
$ sudo hostnamectl set-hostname k8s-master

# initialize Kubernetes on master node
$ sudo kubeadm init --pod-network-cidr=10.244.0.0/16
$ mkdir -p $HOME/.kube/
$ sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
$ sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

<br>

**Start Kubernetes**

```bash
# prepare Dockerfile and .yaml
$ kubectl create -f {pod name}
$ kubectl describe pod
$ kubectl logs {pod name}
```



:white_check_mark: **TODO**

- [x] Change deployment method to **Raspberry PI** & **Docker** using **CI/CD**
- [x] Refactor codes: Follow [Clean architecture in Golang](https://medium.com/@manakuro/clean-architecture-with-go-bce409427d31)
- [x] Make auto-scaled and Add load balancer using **Kubernetes**
