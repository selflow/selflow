---
title: 🐋 Docker Steps
---

Docker steps are used to execute code inside a docker container.

## Options

### `image` (**Required**) {#image}

_Supports Go Template_

Docker image to use as a base.

:::warning
There is currently a limitation and the user of the image must be a root user, otherwise the commands won't be executed.
the image also need to have a shell at `/bin/sh`
:::

### `commands` (**Required**) {#commands}

_Supports Go Template_

Commands to run on the container. For now, that will always be executed by a shell located at `/bin/sh`.
you can use a multiline string fot that field

### `persistence`

The persistence attribute allows step to share files between them using docker volumes.
This way, if a file is added or change in a step, it can also be available in another step **as long as it is a Docker Step**.

The volumes are scoped to a specific run and can't be used in other runs.

```yaml
workflow:
  timeout: 5m
  steps:
    step-a:
      kind: docker
      with:
        persistence:
          my-volume: /workdir
        image: node:lts
        commands: |
          echo "Hello!" > /workdir/someFile.txt

    step-b:
      kind: docker
      needs:
        - step-a
      with:
        persistence:
          my-volume: /workdir
        image: node:lts
        commands: |
          cat /workdir/someFile.txt
```

In this example, `step-a` and `step-b`depends on the same persistence volume called `my-volume`. In both case, it will
be mapped to the `/workdir` path on the container so if a step add a file in this directory, it will be available for the other step.
Using the [`needs`](./#needs) attribute, we specify that `step-a` will be executed before `step-b`.
After `step-a` wrote in `/workdir/someFile.txt` and finished, the file will be available to be read by `step-b`.

Here, `step-b` will log `Hello!`
