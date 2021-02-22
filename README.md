# Lobster

[![Release](https://img.shields.io/github/release/blackhorseya/lobster)](https://github.com/blackhorseya/lobster/releases/latest)
[![Build Status](https://travis-ci.com/blackhorseya/lobster.svg?branch=main)](https://travis-ci.com/blackhorseya/lobster)
[![codecov](https://codecov.io/gh/blackhorseya/lobster/branch/main/graph/badge.svg?token=DJHL70E6ZT)](https://codecov.io/gh/blackhorseya/lobster)
[![Go Report Card](https://goreportcard.com/badge/github.com/blackhorseya/lobster)](https://goreportcard.com/report/github.com/blackhorseya/lobster)
[![Go Reference](https://pkg.go.dev/badge/github.com/blackhorseya/lobster)](https://pkg.go.dev/github.com/blackhorseya/lobster)
[![GitHub license](https://img.shields.io/github/license/blackhorseya/lobster)](https://github.com/blackhorseya/lobster/blob/main/LICENSE)

[Lobster](https://lobster.seancheng.space) is a tool which integration todo list, OKRs, sprint board, pomodoro and
report etc. functional

## Concept

I very like concept of OKR and Agile, so I want to bring this mind to my life for self growth. The project wish to set
goals for long term using concept of OKR then divide to task And sprint tasks via Sprint board, finally achieve that
goal.

## CLI usage

```shell
$ lobster help

██╗      ██████╗ ██████╗ ███████╗████████╗███████╗██████╗ 
██║     ██╔═══██╗██╔══██╗██╔════╝╚══██╔══╝██╔════╝██╔══██╗
██║     ██║   ██║██████╔╝███████╗   ██║   █████╗  ██████╔╝
██║     ██║   ██║██╔══██╗╚════██║   ██║   ██╔══╝  ██╔══██╗
███████╗╚██████╔╝██████╔╝███████║   ██║   ███████╗██║  ██║
╚══════╝ ╚═════╝ ╚═════╝ ╚══════╝   ╚═╝   ╚══════╝╚═╝  ╚═╝

Lobster is a tool which integration todo list, OKRs, sprint board, pomodoro and report etc. functional.

Usage:
  lobster [command]

Available Commands:
  completion  Generate completion script
  config      View and edit Lobster configuration
  create      Create one resource
  delete      Delete one resource
  get         Display one resource
  help        Help about any command
  update      Update one resource

Flags:
      --config string   config file (default is $HOME/.lobster.yaml)
  -h, --help            help for lobster

Use "lobster [command] --help" for more information about a command.

```

### Auto completion

#### Zsh

```shell
lobster completion zsh > ~/.zsh/completion/_lobster
```

### Tasks

```shell
lobster get tasks
+--------------------------------------+--------------------------------------+-----------+---------+---------------------+
| ID                                   | RESULT ID                            | TITLE     | STATUS  | CREATE AT           |
+--------------------------------------+--------------------------------------+-----------+---------+---------------------+
| 836d2bfb-b002-4c78-8479-177b1dad3e5f | f897a24e-14ac-4293-a783-e16f3b885a1b | interview | BACKLOG | 2021-02-22 21:01:30 |
+--------------------------------------+--------------------------------------+-----------+---------+---------------------+
```

## Tech

### WorkFlow

using [Trunk-based development](https://blog.seancheng.space/posts/what-is-trunk-based-development)
Reference [official documents](https://cloud.google.com/solutions/devops/devops-tech-trunk-based-development)

### Dependencies

- [gin](https://github.com/gin-gonic/gin)
- [cobra](https://github.com/spf13/cobra)
- [swaggo](https://github.com/swaggo/swag)
- [wire](https://github.com/google/wire)
- [logrus](https://github.com/sirupsen/logrus)
- [testify](https://github.com/stretchr/testify)
- [mockery](https://github.com/vektra/mockery)
- [viper](https://github.com/spf13/viper)
- [sqlx](https://github.com/jmoiron/sqlx)

### CI/CD

- [Travis-CI](https://travis-ci.com/blackhorseya/lobster)
- [Helm 3](https://helm.sh/)
- [Terraform](https://www.terraform.io/)

### Infrastructure

- [Google Kubernetes Engine](https://cloud.google.com/kubernetes-engine)
- [Cloudflare](https://www.cloudflare.com/zh-tw/)
