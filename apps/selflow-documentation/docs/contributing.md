---
title: Contributing
---

Pull Requests are welcome !
If you have an idea or want to start a major feature, please [open a discussion](https://github.com/selflow/selflow/discussions/new/choose) on GitHub.

## Set up the project locally

### Requirements

- Node.js
- Yarn
- Go

Depending on what you are doing, you might want to have Docker installed and set up locally.

Setting up the project locally is fairly easy. Start by cloning the repository :

```bash
git clone https://github.com/selflow/selflow.git
cd selflow
```

Now you can install NX dependencies using `yarn`

```bash
yarn
```

If needed, you can also install go dependencies using

```bash
go mod download
```

:::note
If you want to change something to the selflow core, you do not have to set up NX on your repository but it will make your life easier
:::

## Repository Structure

The Selflow repository is a mono-repository that is meant to manage the majority of the Selflow ecosystem.
It allows us to have a global vision of the impact pof any new feature.

For that we are using [NX](https://nx.dev/) which is an amazing tool !

The main directories are :

- Go executable : `/apps`
- Go public libraries : `/pkg`
- Go private libraries : `/internal`
- Applications : `/apps`
- E2E projects : `/apps`
- Other libraries : `/libs`

:::note
In the future, the go executables will be moved to the `/apps` directory and the packages to the `/libs` directories
:::
