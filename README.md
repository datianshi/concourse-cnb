# concourse-cnb

This concourse task is to leverage [cloud native buildpack](https://buildpacks.io/) to build image from source to image

It references the [teckton cloud native buildpack](https://github.com/tektoncd/catalog/tree/main/task/buildpacks) implementation

## input

    ```
        APP_IMAGE: # Application image registry
        RUN_IMAGE: # build pack run image. default: gcr.io/paketo-buildpacks/run:base-cnb
        IMAGE_REPO: # default: index.docker.io
        IMAGE_REPO_USERNAME: # e.g. dockerhub username
        IMAGE_REPO_PASSWORD: # e.g. dockerhub password
        BUILD_ENV_BP_KEEP_FILES: # keep cetain artifacts in the run workspace
    ```

## output

* image output has a file with image digest number

    ```
    cat image/digest 
    sha256:398d6977720147febe5ff15011ce0235187658a46a2b28c2013fe9465626cb27
    ```
## A Sample pipeline

```
resources:
- name: source
  type: git
  source:
    uri: https://github.com/spring-projects/spring-petclinic
    branch: main
    username: ((github_token))
    password: x-oauth-basic
- name: buildpack
  type: git
  source:
    uri: https://github.com/datianshi/concourse-cnb
    branch: main
jobs:
- name: build
  plan:
  - in_parallel:
    - get: source
      trigger: true
    - get: buildpack
  - task: build
    input_mapping:
      pipeline: buildpack
    params:
      APP_IMAGE: myorg/spring-petclinic
      IMAGE_REPO_USERNAME: ((docker_username))
      IMAGE_REPO_PASSWORD: ((docker_password))
    file: buildpack/buildpack/task.yaml
```

