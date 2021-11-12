# concourse-cnb

This concourse task is to leverage [cloud native buildpack](https://buildpacks.io/) to build image from source to image

It references the [teckton cloud native buildpack](https://github.com/tektoncd/catalog/tree/main/task/buildpacks) implementation

## input params

    ```
        APP_IMAGE: # Application image registry
        RUN_IMAGE: # build pack run image. default: gcr.io/paketo-buildpacks/run:base-cnb
        IMAGE_REPO: # default: index.docker.io
        IMAGE_REPO_USERNAME: # e.g. dockerhub username
        IMAGE_REPO_PASSWORD: # e.g. dockerhub password
        BUILD_ENV_BP_KEEP_FILES: # keep cetain artifacts in the run workspace
    ```

## Optional Input

* env

Buildpack may need environment variable during build time. The docker container **datianshi/cnb-config** provides a utility to generate environment variables.


```
- task: prepare_build_env_variable
  image: buildpack-config
  config:
    outputs:
    - name: env
    platform: linux
    params:
      BUILD_ENV_BP_VAR1: VALUE1
      BUILD_ENV_BP_VAR2: VALUE2
    run:
      path: /bin/sh        
      args:
      - -eec
      - |
        buildpack env
```

This above example reads env with prefix BUILD_ENV_ and generate output inside env folder:

* file BP_VAR1 with content VALUE1
* file BP_VAR2 with content VALUE2

The buildpack task then copy tne output env folder to /platform/bindings

* bindings

Buildpack may need secrets e.g. maven repo secret during build time. The docker container **datianshi/cnb-config** provides a utility to generate [bindings](https://paketo.io/docs/howto/configuration/#bindings).

```
- task: prepare_maven_binding
  image: buildpack-config
  config:
    outputs:
    - name: bindings
    platform: linux
    params:
      MAVEN_REPO_USERNAME: ((maven_repo_username))
      MAVEN_REPO_PASSWORD: ((maven_repo_password))
      MAVEN_REPO_ID: ((maven_repo_id))
    run:
      path: /bin/sh        
      args:
      - -eec
      - |
        buildpack bindings -n maven -t maven maven-settings -u ${MAVEN_REPO_USERNAME} -p ${MAVEN_REPO_PASSWORD} -r ${MAVEN_REPO_ID}
```

The above example generate a bindings output folder with

```
bindings
--binding_name
  --settings.xml
  --type
```

The buildpack task then copy the output bindings folder to /platform/bindings

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

