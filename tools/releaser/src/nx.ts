import {getCommitsSinceTag} from "./git";
import {spawnSync} from "child_process";

export function getAffectedProjects(commitHash: string): string[] {
  try {
    const {
      stdout,
    } = spawnSync('yarn -s nx show projects --type app --affected', {
      shell: "/bin/bash",
      env: {
        ...process.env,
        NX_HEAD: commitHash,
        NX_BASE: `${commitHash}^1`,
      }
    })

    return stdout.toString()
      .split('\n')
      .map(s => s.trim())
      .filter(s => !!s)

  } catch (e: any) {
    console.error('[ERROR] ', e)
    return []
  }
}

/**
 * Return if a core release should be extracted from a commit list.
 *
 * A core release is a release that involves at least one NX project with the tag "scope:core"
 * @param commits
 */
export async function checkCoreRelease(commits: Awaited<ReturnType<typeof getCommitsSinceTag>>) {
  // The following command list affected projects between the first and the last commit of the list,
  // then, for each project, details it to extract the tags
  // then, for each json documents representing projects, the tags are extracted and output to the console
  // It just was efficient enough to do it in bash but of you don't like it, feel free to recode it using NX sdk ;)
  const {
    stdout,
    stderr
  } = spawnSync('yarn -s nx show projects --affected | xargs -I % bash -c "yarn -s nx show project % --json" | jq -n \'[ inputs.tags ]\'', {
    shell: "/bin/bash",
    env: {
      ...process.env,
      NX_HEAD: commits[0].hash,
      NX_BASE: commits[commits.length - 1].hash,
    }
  })

  const stderrAsString = stderr.toString()

  if (stderrAsString.length) {
    console.error(stderrAsString)
  }

  const tags = JSON.parse(stdout.toString()).flat()

  return tags.includes("scope:core")
}
