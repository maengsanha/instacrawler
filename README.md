<img src="https://lh5.googleusercontent.com/proxy/r5D7LX7gbvXfuJU1SFAfCM1SerPt0KcBvR_R0qpXO_fsa39nwCKhyGE0UQbFP99XpSMRuPWrckLRnkoU747FW6EHY1_Gqf1xzhXYhJnIqIHizuhbBX3fh0sgdxbpIwJrDtC9g-uELzM-xYNfiw=s0-d">

#### INSTAGRAM CRAWLER

[![build](https://img.shields.io/badge/build-failed-red?style=flat&logo=github)](https://github.com/maengsanha/instacrawler/pulse)
[![license](https://img.shields.io/badge/license-MIT-blue)](https://github.com/maengsanha/instacrawler/blob/master/LICENSE)

<br>

**Installation**

```bash
$ go get -u github.com/maengsanha/instacrawler
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
```

<br>

**Install kubectl (Ubuntu)**

```bash
$ snap install kubecetl --classic
```



:white_check_mark: **TODO**

- [x] Change deployment method to **Raspberry PI** & **Docker** using **CI/CD**
- [x] Refactor codes: Follow [Clean architecture in Golang](https://medium.com/@manakuro/clean-architecture-with-go-bce409427d31)
- [ ] Make auto-scaled and Add load balancer using **Kubernetes**
