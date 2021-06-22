# image-resizer

[![RELEASE](https://img.shields.io/github/v/release/blyndusk/image-resizer)](https://github.com/blyndusk/image-resizer/releases)
[![RELEASE DATE](https://img.shields.io/github/release-date/blyndusk/image-resizer)](https://github.com/blyndusk/image-resizer/commits/main)
[![ISSUES](https://img.shields.io/github/issues/blyndusk/image-resizer)](https://github.com/blyndusk/image-resizer/issues)
[![PULL REQUESTS](https://img.shields.io/github/issues-pr/blyndusk/image-resizer)](https://github.com/blyndusk/image-resizer/pulls)
[![LAST COMMIT](https://img.shields.io/github/last-commit/blyndusk/image-resizer)](https://github.com/blyndusk/image-resizer/commits/main)
[![LICENSE](https://img.shields.io/github/license/blyndusk/image-resizer)](https://github.com/blyndusk/image-resizer/blob/main/LICENSE)
[![REPO SIZE](https://img.shields.io/github/repo-size/blyndusk/image-resizer)](https://github.com/blyndusk/image-resizer)





- [image-resizer](#image-resizer)
  - [I - Intro](#i---intro)
    - [1 - Goal](#1---goal)
    - [2 - Stack](#2---stack)
  - [III - Continuous integration](#iii---continuous-integration)
    - [1 - On each push on every branch](#1---on-each-push-on-every-branch)
    - [2 - On each push on `main` branch](#2---on-each-push-on-main-branch)
    - [3 - Manually triggered](#3---manually-triggered)
  - [III - Schemas](#iii---schemas)
    - [1 - Local infrastructure](#1---local-infrastructure)
    - [2 - Cloud infrastructure](#2---cloud-infrastructure)
    - [3 - Database schema](#3---database-schema)
  - [IV - Conventions, templates and guidelines](#iv---conventions-templates-and-guidelines)
    - [A - Commit conventions](#a---commit-conventions)
    - [B - Issue template](#b---issue-template)
    - [C - Branch naming convention](#c---branch-naming-convention)
    - [D - Pull request template](#d---pull-request-template)
  - [V - Project use](#v---project-use)
    - [Help](#help)
    - [Setup env](#setup-env)
    - [Start](#start)
    - [Worker](#worker)
    - [Stop](#stop)
    - [Restart](#restart)
    - [Display logs](#display-logs)
    - [Lint](#lint)
  - [V - License](#v---license)

## I - Intro

### 1 - Goal

Image Resizer is a school project made for @ecolehetic, with the aim to automate the resize of uploaded avatar from a user on a forum.

### 2 - Stack

| Service    | Type                       |
| ---------- | -------------------------- |
| Reactjs    | Front-end framework        |
| Golang     | Back-end language          |
| PostgresQL | Object-relational database |
| RabbitMQ   | Message broker             |

## III - Continuous integration

### 1 - On each push on every branch 

| FRONT | BACK |
| --- | --- |
| [![REACT](https://github.com/blyndusk/image-resizer/actions/workflows/react.yml/badge.svg)](https://github.com/blyndusk/image-resizer/actions/workflows/react.yml) | [![GO](https://github.com/blyndusk/image-resizer/actions/workflows/go.yml/badge.svg)](https://github.com/blyndusk/image-resizer/actions/workflows/go.yml) |

### 2 - On each push on `main` branch

[![DOCKER](https://github.com/blyndusk/image-resizer/actions/workflows/docker.yml/badge.svg)](https://github.com/blyndusk/image-resizer/actions/workflows/docker.yml)

### 3 - Manually triggered

[![RELEASE](https://github.com/blyndusk/image-resizer/actions/workflows/release.yml/badge.svg)](https://github.com/blyndusk/image-resizer/actions/workflows/release.yml)


## III - Schemas

### 1 - Local infrastructure

![local-infrastructure](./docs/infrastructure-local-solution.png)

### 2 - Cloud infrastructure

![cloud-infrastructure](./docs/infrastructure-cloud-solution.png)

### 3 - Database schema

![database-schema](./docs/database-schema.png)

## IV - Conventions, templates and guidelines

### A - Commit conventions

```
tag: #issue_id - message
```

See [commit_conventions.md](.github/commit_conventions.md) for more informations.

### B - Issue template

See [user-story.md](.github/ISSUE_TEMPLATE/user-story.md) for more informations.

### C - Branch naming convention

```
type_scope-of-the-work
```

See [branch_naming_convention.md](.github/branch_naming_convention.md) for more informations.

### D - Pull request template

See [pull_request_template.md](.github/pull_request_template.md) for more informations.

## V - Project use

### Help

```bash
$ make help
```

### Setup env

```bash
$ make setup-env
```

This command will copy env file. Please update `api/.env` with actual default environment variables, the same from the `docker-compose.yaml`.

### Start

```bash
$ make start
```

This command will start the projet via `docker-compose.yaml`, with full reset.

### Worker

```bash
$ make start-producer
```

This command will start the queue producer, which'll add the images to the queue

```bash
$ make start-consumer
```

This command will start the queue producer, which'll resize the images


### Stop

```bash
$ make start
```

This command will shutdown every container.

### Restart

```bash
$ make restart
```

Stop then start

### Display logs

```bash
$ make logs
```

### Lint

```bash
$ make lint-app
```

This command use `ESLint`.

```bash
$ make lint-api
```

This command use `gofmt`.

## V - License

Under [MIT](./LICENSE) license.
