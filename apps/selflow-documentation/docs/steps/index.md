---
title: Steps
---

:::info
For clarity, the steps configuration are showed in a yaml format but you can also use the same keys in a json file and
it will still run fine !
:::

## Write a step

Steps are written in a workflow file

```yaml
workflow:
  timeout: 5m
  steps:
    my-step:
      kind: docker
      with:
        image: 'alpine:3.10.0'
        commands: |
          echo "Selflow is love, Selflow is Life !"
```

In the above example, we are creating a step called `my-step`.
This step has a `kind`, this is a **mandatory field** that says the kind of step to use, here _"docker"_.
It also has an object passed with the `with` key with a few properties that the docker step needs.
This is also a **mandatory field**.

## Optional configuration

### `needs`

#### Syntax

```yaml
workflow:
  timeout: 5m
  steps:
    # ...
    # Other steps definitions
    # ...
    my-step:
      kind: docker
      needs:
        - some-step
        - some-other-step
      with:
        image: 'alpine:3.10.0'
        commands: |
          echo "Selflow is love, Selflow is Life !"
```

In this example, the step `my-step` depends on steps `some-step` and `some-other-step` so it will start only after both steps stops with a status of _SUCCESS_.

The nice thing about step dependencies, is that **they can share variables**.

#### Example

```yaml
workflow:
  timeout: 5m
  steps:
    step-a:
      kind: docker
      with:
        image: 'alpine:3.10.0'
        commands: |
          echo "::output::foo::bar"

    step-b:
      kind: docker
      needs:
        - step-a
      with:
        image: 'alpine:3.10.0'
        commands: |
          echo "{{ .Needs "step-a" "foo" }}"
```

A lot of things are happening there. First we have two steps, `step-a` and `step-b`, both docker steps.

The step `step-a` logs "_::output::foo::bar_".
This syntax is used to tell the docker steps that it needs to output a variable called "foo" with the value of "bar".

The in `step-b`, with are using this syntax `{{ .Needs "step-a" "foo" }}`.
Here we are using [Go Template](https://pkg.go.dev/text/template) to declare that the step needs the value of the variable "foo" output by `step-a`.

So in this example, `step-b` will output "bar".

You can find more details about how to use [Go Template](https://pkg.go.dev/text/template) on its great documentation !

### `if`

#### Syntax :

```yaml
workflow:
  timeout: 5m
  steps:
    my-step:
      kind: docker
      if: 'false'
      with:
        image: 'alpine:3.10.0'
        commands: |
          echo "This will not be executed"
```

The `if` command allows to execute a step only if the specified condition is truthy.
The condition is considered truthy if it is contains an empty string, "no" or "false".

It doesn't do much in the example above but you can combine it with step variables as explained in the [`Needs`](#needs) attribute to make it dynamic !

#### Example

```yaml
workflow:
  timeout: 5m
  steps:
    step-a:
      kind: docker
      with:
        image: "alpine:3.10.0"
        commands: |
          # a will be equal to either "true" of "false" at random
          a=$(if [ "$(($RANDOM % 2))" -eq "0" ]; then echo "true"; else echo "false"; fi)
          echo "::output::foo::$a"

    step-b:
      kind: docker
      needs:
        - step-a
      if: "{{ .Needs "step-a" "foo" }}"
      with:
        image: "alpine:3.10.0"
        commands: |
          echo "I might be executed"
```

It this example, `step-b` depends on `step-a`.
`step-a` will output a "foo" variable that will randomly equals "true" or "false".
The output is used in the `if` attribute of the second step to only execute it if `step-a` returned "true".
