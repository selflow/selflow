---
title: Getting Started
---

The easiest way to get started with Selflow is to use the [Selflow CLI](./ecosystem/cli/index.mdx) and run the [Selflow Daemon](./ecosystem/selflow-daemon). You can follow [this documentation](./ecosystem/cli#installation).

Then, you can start creating your first workflow ! Open a yaml file and write this :

```yaml
workflow:
  timeout: 5m
  steps:
    my-step:
      kind: docker
      with:
        image: alpine:3.10.0
        commands: |
          echo "Hello World !"
```

You can then run this with

```bash
selflow-cli run ./path/to/your/file.yaml
```

A run will now start, triggering a step in a docker container using the "alpine:3.10.0" image. This step will log "Hello World !".

For more informations, go to the [Workflow Syntax](./workflow-syntax) documentation, the [Docker step](./steps/docker) specifications or the [Selflow CLI](./ecosystem/cli) documentation.
