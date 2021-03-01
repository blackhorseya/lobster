# Lobster

[![Build Status](https://travis-ci.com/blackhorseya/lobster.svg?branch=main)](https://travis-ci.com/blackhorseya/lobster)
[![codecov](https://codecov.io/gh/blackhorseya/lobster/branch/main/graph/badge.svg?token=DJHL70E6ZT)](https://codecov.io/gh/blackhorseya/lobster)
[![Go Report Card](https://goreportcard.com/badge/github.com/blackhorseya/lobster)](https://goreportcard.com/report/github.com/blackhorseya/lobster)
[![Release](https://img.shields.io/github/release/blackhorseya/lobster)](https://github.com/blackhorseya/lobster/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/blackhorseya/lobster)](https://pkg.go.dev/github.com/blackhorseya/lobster)
[![GitHub license](https://img.shields.io/github/license/blackhorseya/lobster)](https://github.com/blackhorseya/lobster/blob/main/LICENSE)

[Lobster](https://lobster.seancheng.space/api/docs/index.html) is a tool which integration todo list, OKRs, sprint
board, pomodoro and report etc. functional

## Concept

I very like concept of OKR and Agile, so I want to bring this mind to my life for self growth. The project wish to set
goals for long term using concept of OKR then divide to task And sprint tasks via Sprint board, finally achieve that
goal. I don't like develop front-end so using [Lobster-CLI](https://github.com/blackhorseya/lobster-cli) for control
the API.

## Tech

### WorkFlow

using [Trunk-based development](https://blog.seancheng.space/posts/what-is-trunk-based-development)
Reference [official documents](https://cloud.google.com/solutions/devops/devops-tech-trunk-based-development)

### Dependencies

- [gin](https://github.com/gin-gonic/gin) for web server framework
- [swaggo](https://github.com/swaggo/swag) for swagger spec
- [wire](https://github.com/google/wire) for dependency inject
- [logrus](https://github.com/sirupsen/logrus) for logger
- [testify](https://github.com/stretchr/testify) for unit test
- [mockery](https://github.com/vektra/mockery) for mock
- [viper](https://github.com/spf13/viper) for configuration
- [sqlx](https://github.com/jmoiron/sqlx) for sql driver

### CI/CD

- [Travis-CI](https://travis-ci.com/blackhorseya/lobster) for CI/CD
- [Helm 3](https://helm.sh/) for manage deployment to Kubernetes

### Infrastructure

- [Google Kubernetes Engine](https://cloud.google.com/kubernetes-engine) for Kubernetes
- [Cloudflare](https://www.cloudflare.com/zh-tw/) for DNS
- [Terraform](https://www.terraform.io/) for Infra-as-Code
