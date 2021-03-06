# News-API
[![Build Status](https://travis-ci.org/IIC2173-2015-2-Grupo2/news-api.svg)](https://travis-ci.org/IIC2173-2015-2-Grupo2/news-api) [![Code Climate](https://codeclimate.com/github/IIC2173-2015-2-Grupo2/news-api/badges/gpa.svg)](https://codeclimate.com/github/IIC2173-2015-2-Grupo2/news-api)

> Format & style: [`gofmt`](https://golang.org/cmd/gofmt/), [`golint`](https://github.com/golang/lint), [`govet`](https://golang.org/cmd/vet/)

## API Documentation

### Objects

#### `NewsItem`

Describes a NewsItem.

| Field | Type | Description |
|-------|------|-------------|
| `id` | `number` | Unique ID |
| `title` | `string` | - |
| `url` | `string` | - |
| `source` | `string` | Provider |
| `language` | `string` | Language: `en`, `es`. |
| `body` | `string` | Short description |
| `image` | `string` | Representative image URL |

#### `User`

Describes a User.

| Field | Type | Description |
|-------|------|-------------|
| `id` | `number` | Unique ID |
| `name` | `string` | - |
| `username` | `string` | Primary key of `users` |
| `email` | `string` | - |

#### `Tag`

Describes a Tag. News have many tags.

| Field | Type | Description |
|-------|------|-------------|
| `id` | `number` | Unique ID |
| `name` | `string` | Tag name |

#### `Category`

Describes a Category. News belongs to one category.

| Field | Type | Description |
|-------|------|-------------|
| `id` | `number` | Unique ID |
| `name` | `string` | Category name |

#### `Person`

Describes a Person. News are related with many people.

| Field | Type | Description |
|-------|------|-------------|
| `id` | `number` | Unique ID |
| `name` | `string` | Person name |

#### `Location`

Describes a Location. News belongs to one location.

| Field | Type | Description |
|-------|------|-------------|
| `id` | `number` | Unique ID |
| `name` | `string` | Location name |

#### `NewsProvider`

Describes `NewsItem` source. NewsItem has one NewsProvider

| Field | Type | Description |
|-------|------|-------------|
| `id` | `number` | Unique ID |
| `name` | `string` | News Provider name |

---

### API Usage

#### `POST` `api/v1/auth/signup`

Creates an account.

| Argument | Type | Description |
|----------|------|-------------|
| `name`  | `string`| - |
| `username`  | `string`| - |
| `email` | `string`| - |
| `password`  | `string`| - |

This returns the account session token

| Argument | Type | Description |
|----------|------|-------------|
| `token`  | `string`| Access token |

#### `POST` `api/v1/auth/token`

Re-new expired token

| Argument | Type | Description |
|----------|------|-------------|
| `username`  | `string`| - |
| `password`  | `string`| - |

This returns the account session token

| Argument | Type | Description |
|----------|------|-------------|
| `token`  | `string`| Access token |

### Authenticated API Usage
> Each requests must include a valid non-expired `Bearer <token>`  as `Authentication` header.
> Otherwise will return a `401 Unauthorized` status code.

#### `GET` `api/v1/private/news`

Returns a `NewsItem`'s list

| Argument | Type | Default | Description |
|----------|------|---------|-------------|
| `page` | `uint` | `0` | Page number |

#### `GET` `api/v1/private/news/:id`

Returns the `NewsItem` associated with that `id`

#### `GET` `api/v1/private/search`

Search `NewsItem` with:

| Argument | Type | Default | Description |
|----------|------|---------|-------------|
| `page` | `uint` | `0` | Page number |
| `tags` | `[]string` | | Filter by `Tag`'s name |
| `providers` | `[]string` | | Filter by `NewsProvider`'s name|
| `locations` | `[]string` | | Search news in this locations|
| `people` | `[]string` | | Search news related with|
| `categories` | `[]string` | | Filter by Categories name|

##### Example
```sh
.../api/v1/private/news?tags=love&tags=kill&providers=cnn&providers=emol&people=name&categories=national
```
#### `GET` `api/v1/private/tags`

Returns a `Tag`'s list

#### `GET` `api/v1/private/news_providers`

Returns a `NewsProvider`'s list

#### `GET` `api/v1/private/categories`

Returns a `Category`'s list

#### `GET` `api/v1/private/people`

Returns a `Person`'s list

#### `GET` `api/v1/private/locations`

Returns a `Location`'s list


## Development

Install [Golang](https://golang.org/).

Make sure to configure `$GOPATH`. For example:
```sh
$ export GOPATH=$HOME/Repositories/go
$ export PATH=$PATH:$GOPATH/bin
```

Get this repository using `go`:
```sh
$ go get github.com/tools/godep
$ go get github.com/IIC2173-2015-2-Grupo2/news-api

# Project directory
$ cd $GOPATH/src/github.com/IIC2173-2015-2-Grupo2/news-api/
```

### Local

Setup database
```sh
# Example values
export NEO4J_HOST="192.168.99.100"
export NEO4J_PORT="7474"
export NEO4J_USER="neo4j"
export NEO4J_PASS="neo4j"

export ANALYTICS_TOKEN="GOOGLE_ANALYTICS_TOKEN"
```

Setup other environment variables
```sh
export ENVIRONMENT="PRODUCTION"
export SECRET_HASH="SECRET_HASH"
export AUTH="ENABLE"
export LOADER_IO_TOKEN="IO_TOKEN"
```

Build and run the project locally using:
```sh
$ make start
```

### [Docker](https://www.docker.com/)

Build and run:
```sh
$ docker-compose up -d
```

##### Setup Postgres

Log to postgres container:
```sh
$ docker exec -it newsapi_db_1 /bin/bash
```

Log to postgres service:
```sh
$ psql -U postgres
```

Create Database
```sh
CREATE DATABASE newsapi;
```
