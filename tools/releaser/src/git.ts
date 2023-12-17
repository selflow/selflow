import {Commit} from "./types";
import simpleGit from "simple-git";
import {parseCommit} from "./gitmoji";
import {getAffectedProjects} from "./nx";
import {gitConfig} from "./config";

function getGit() {
  return simpleGit({
    config: [
      `user.name="${gitConfig.username}"`,
      `user.email="${gitConfig.email}"`,
    ]
  })
}

export async function getLastTag() {
  const git = getGit();
  const {latest: lastTag} = await git.tags()
  return lastTag
}

export async function getCommitsSinceTag(tag: string): Promise<Commit[]> {
  const git = getGit();
  const {all: commits} = await git.log({
    from: tag,
    to: 'HEAD'
  })

  return Promise.all(commits.map(async c => {
    console.log("[DEBUG] Analyzing commit", c.message)
    let commit: Commit = {
      ...c,
      smallHash: c.hash.slice(0, 8),
    }

    commit = parseCommit(commit)
    commit.affectedProjects = getAffectedProjects(c.hash)

    return commit
  }))
}

export async function createReleaseCommit(releaseName: string, changelog: string, ...files: string[]) {
  const git = getGit()
  await git.commit(`:bookmark: Release ${releaseName}\n
[skip ci]

--- CHANGELOG ---
${changelog}
`, files)
}

export async function createReleaseTag(releaseName: string) {
  const git = getGit()
  await git.addTag(releaseName)
}

export async function pushChanges() {
  const git = getGit()
  await git.push(gitConfig.remoteName, gitConfig.releaseBranch, ["--tags"])
}
