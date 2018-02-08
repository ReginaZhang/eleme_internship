#### This is not a working project!
* Repo purpose: to save the work and the memory of the internship in Ele.me (饿了么) 
* Many thanks to: Tools team @Core Infrastructure Department, Ele.me
* Special thanks to: Ed Huang for helping me get the opportunity and guiding me through the project ^-^

##### Sensitive information of the company is removed from the project

# Pansible

> Ansible in parallel

## Introduction

- Provide dynamic inventory
- Split ansible playbook and run them on multiple workers
- Wraps app deployment as service

## Goals

- Simplify the deployment of infrastructure app
- Provides easy configuration management
- Speed up ansible, so that ansible jobs at large scale roughly feel like
  running the same job against 10-20 hosts

## Design

- Server starts in a docker container
- Ansible playbooks are also ran in docker containers after the runner fetches the
  playbook and ssh keys from the server.
- Task status are sent back to the runners by callback plugin through http(s). Runners
  save a copy of the contents in memory and then feedback to the server with 
  indefinite retry in case of server failure

## Development

Setup a dev environment:

```
docker-compose up -d
make setup
```

To build the project, run:

```
make build
```

To cross compile for other platform by adding `OS` and `ARCH` to set
`GOOS` and `GOARCH`, for example, to build for linux-amd64:

```
make build OS=linux ARCH=amd64
```

The built executable can be found at `build/${OS}_${ARCH}`.

To build the docker image, run:
```
docker build -t pansible:dev .
```

To push image to registry, add path to the registry as variable `IMAGE` in the Makefile,
for example, `IMAGE := docker.elenet.me/john.doe/pansible`, then run:
```
make -B docker push
```

## Deployment

### Configuration

To configure pansible using a config file, see `pansible.yml` for example.  
OR  
Every field in the config file can be set with environment variable,
by adding `PANSIBLE` as prefix.  
For example, `PANSIBLE_PORT=:5757`.

### Starting the server

Using binary:

```
pansible server --config ${PATH_TO_CONFIG}
```
> Config file defaults to `./pansible.yml`

Using docker image:

Pull from registry
```
docker pull ${PATH_TO_REGISTRY}/pansible:dev
```

To start the server using docker container
```
docker run -p 5757:5757 -p 5858:5858 --network pansible_default --mount type=bind,source=${PATH_TO_SSHKEY_ON_HOST},target=/opt/pansible/tmp/keydir,readonly ${PATH_TO_REGISTRY}/pansible:dev ./pansible server --config ${PATH_TO_CONFIG}
```

## Usage

Pansible uses restful API with JSON body.  
Valid `COFFEE_TOKEN` required.

### Create an inventory

Upload through front-end form

OR 

Create through http request call  
Method: POST  
Path: /api/v1/inventory  
Body:
```
{
  "name": "example_inventory",
  "env": "alpha",
  "version": "0.0.1",
  "hosts": "all:
              hosts:
                localhost
              children:
                webservers:
                  hosts:
                    foo.example[3:5].com:
                dbservers:
                  hosts:
                    one.example.com:
                    two.example.com:",
  "vars": "all:
             allkey1: allval1
             allkey2: allval2
           webservers:
             wskey1: wsval1
             wskey2: wsval2"
}
```
> Note: `hosts` and `vars` should be in Yaml format as shown above for both methods

Response body: 
>{  
   &nbsp;&nbsp;"id": 5,  
   &nbsp;&nbsp;"name": "example_inventory",  
   &nbsp;&nbsp;"env": "alpha",  
   &nbsp;&nbsp;"version": "0.0.1",  
   &nbsp;&nbsp;"hosts": "{SAME_AS_UPLOADED}",  
   &nbsp;&nbsp;"vars": "{SAME_AS_UPLOADED}"  
 }

### Create a playbook

Method: POST  
Path: /api/v1/playbook  
Body:
```json
{
  "name": "example_playbook",
  "git_repo": "git@some_repo.git",
  "entry": "some_playbook.yml"
}

```
Response body:
>{  
   &nbsp;&nbsp;"id": 3,  
   &nbsp;&nbsp;"name": "example_playbook",  
   &nbsp;&nbsp;"git_repo": "git@some_repo.git",  
   &nbsp;&nbsp;"entry": "some_playbook.yml"  
 }

### Create an app

##### Add a new app
Method: POST  
Path: /api/v1/app  
Body:
```json
{
  "appid": "example_app"
}

```
Response body:
>{  
  &nbsp;&nbsp;"id": 7,  
  &nbsp;&nbsp;"appid": "example_app"  
 }

##### Add a playbook to app
Method: POST  
Path: /api/v1/app/playbook  
Body:
```json
{
  "app_id": 7,
  "playbook_id": 3 
}

```

##### Add an inventory to app
Method: POST  
Path: /api/v1/app/inventory  
Body:
```json
{
  "app_id": 7,
  "inventory_id": 5 
}

```


### Run a job

Method: POST  
Path: /api/v1/job  
Body:
```json
{
  "env": "alpha",
  "playbook_id": 19,
  "inventory_id": 32
}
```

Response body:
>{  
  &nbsp;&nbsp;"id": 172,  
  &nbsp;&nbsp;"uuid": "6e63a20f-e1d9-47cc-84ae-cf5bd46d4161",  
  &nbsp;&nbsp;"env": "alpha",  
  &nbsp;&nbsp;"playbook_id": 19,  
  &nbsp;&nbsp;"inventory_id": 32  
}


### Get job results

Method: GET  
Path: /api/v1/job/result/{JOB_UUID}

Response body:
```json
{
  "localhost": {
    "failures": [
      {
        "id": 29,
        "result": "{ANSIBLE_FAILURE_INFO}",
        "stats_id": 316
      }
    ],
    "stats": {
      "failed": 1,
      "id": 316,
      "job_start_time": "2018-01-30T03:33:45Z",
      "run_id": "18c77843-954a-49e0-9a36-cd270af665bb",
      "target": "localhost",
      "unreachable": false
    }
  },
  "one.example.com": {
    "failures": null,
    "stats": {
      "failed": 0,
      "id": 317,
      "job_start_time": "2018-01-30T03:33:45Z",
      "run_id": "18c77843-954a-49e0-9a36-cd270af665bb",
      "target": "one.example.com",
      "unreachable": true
    }
  },
  "two.example.com": {
    "failures": null,
    "stats": {
      "failed": 0,
      "id": 318,
      "job_start_time": "2018-01-30T03:33:45Z",
      "run_id": "18c77843-954a-49e0-9a36-cd270af665bb",
      "target": "two.example.com",
      "unreachable": false
    }
  }
}
```

### Get job real-time progress

Protocol: Websocket  
Path: /api/v1/job/progress/{JOB_UUID}

Callback json will be sent through  
Example:
> {  
&nbsp;&nbsp;"_pansible_task": "18c77843-954a-49e0-9a36-cd270af665bb",  
&nbsp;&nbsp;"type": "unreachable",  
&nbsp;&nbsp;"target": "one.example.com",  
&nbsp;&nbsp;"_pansible_job": "3ea081a4-f312-4ff7-8ca1-b053e5ef1a7e",  
&nbsp;&nbsp;" result": {  
&nbsp;&nbsp;&nbsp;&nbsp;"msg": "Failed to connect to the host via ssh: ssh: Could not resolve hostname one.example.com: nodename nor servname provided, or not known",  
&nbsp;&nbsp;&nbsp;&nbsp;"unreachable": true,  
&nbsp;&nbsp;&nbsp;&nbsp;"changed": false  
&nbsp;&nbsp;}  
}

