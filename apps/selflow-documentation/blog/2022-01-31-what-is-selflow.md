---
slug: what-is-selflow
title: What is Selflow ?
authors: anthony-quere
tags: []
---

Selflow is a simple workflow orchestrator.

Workflows in Selflow are a set of tasks with requirements or none that will run asynchronously in a Docker container

It is currently a school project but I have the ambition to make it grow even after.

It was designed to be runnable on any linux machine in your servers, homelab or in the cloud.

## Easy to use

- Build workflows using Yaml
- Write your own plugins with [the language you want](https://grpc.io/docs/languages/)
- Supports Docker and SSH steps
- See your workflow execution using a simple ui ([wip](./state-of-selflow))
