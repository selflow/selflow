import {getCommitsSinceTag} from "./git";
import {spawnSync} from "child_process";
import * as process from "process";

export function getAffectedProjects(commitHash: string): string[] {
  return getAffectedProjectsBetweenCommitHashes(commitHash, `${commitHash}^1`)
}

export function getAffectedProjectsBetweenCommitHashes(headCommitHash: string, baseCommitHash: string): string[] {
  try {
    const {
      stdout,
    } = spawnSync('yarn -s nx show projects --type app --affected', {
      shell: "/bin/bash",
      env: {
        ...process.env,
        NX_HEAD: headCommitHash,
        NX_BASE: baseCommitHash,
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

export function getProjectDetails(project: string): null |{tags: string[]} {
  try {
    const {
      stdout,
    } = spawnSync(`yarn -s nx show project ${project} --json`, {
      shell: "/bin/bash",
      env: {
        ...process.env,
      }
    })

    return JSON.parse(stdout.toString())
  } catch (e: any) {
    console.error('[ERROR] ', e)
    return null
  }
}

/**
 * Return if a core release should be extracted from a commit list.
 *
 * A core release is a release that involves at least one NX project with the tag "scope:core"
 * @param commits
 */
export async function checkCoreRelease(commits: Awaited<ReturnType<typeof getCommitsSinceTag>>): Promise<boolean> {
  // The following command list affected projects between the first and the last commit of the list,
  // then, for each project, details it to extract the tags
  // then, for each json documents representing projects, the tags are extracted and output to the console
  // It just was efficient enough to do it in bash but of you don't like it, feel free to recode it using NX sdk ;)

  const projectDetails = await Promise.all(getAffectedProjectsBetweenCommitHashes(commits[commits.length - 1].hash, commits[0].hash)
    .map(project => new Promise<ReturnType<typeof getProjectDetails>>(resolve => resolve(getProjectDetails(project)))))

  return projectDetails.some(p => p.tags.includes("scope:core"))
}
