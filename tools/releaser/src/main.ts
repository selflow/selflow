import {readFileSync, writeFileSync} from 'fs'
import {updateJsonFile} from "@nx/workspace";
import {ReleaseType} from "./types";
import {createReleaseCommit, createReleaseTag, getCommitsSinceTag, getLastTag, pushChanges} from "./git";
import {checkCoreRelease} from "./nx";
import {getReleaseType} from "./gitmoji";
import {buildChangelog} from "./changelog";
import {getNextRelease} from "./semver";
import {createRelease} from "./github";


async function main() {
  console.log("[DEBUG] Get last Git tag...")
  const lastTag = await getLastTag()

  console.log("[DEBUG] Get commits since last tag...")
  const commitsSinceLastTag = await getCommitsSinceTag(lastTag)
  if (commitsSinceLastTag.length === 0) {
    console.log("[INFO]  No commit found, exiting")
    return
  }

  console.log("[DEBUG] Checking for core release...")
  const isCoreRelease = await checkCoreRelease(commitsSinceLastTag)

  const oldestCommit = commitsSinceLastTag[commitsSinceLastTag.length - 1];
  const newestCommit = commitsSinceLastTag[0];

  console.log("[DEBUG] Get Release Type...")
  let releaseType: ReleaseType = await getReleaseType(commitsSinceLastTag)

  const shouldRelease = !!releaseType

  if (shouldRelease && !isCoreRelease) {
    releaseType = "prerelease"
  }

  let nextRelease = 'None'
  if (shouldRelease) {
    nextRelease = getNextRelease(lastTag, releaseType)
  }

  console.log("[INFO]  Commit count:\t", commitsSinceLastTag.length)
  console.log("[INFO]  Oldest Commit:\t", `${oldestCommit.message} (${oldestCommit.hash})`)
  console.log("[INFO]  Newest Commit:\t", `${newestCommit.message} (${newestCommit.hash})`)
  console.log("[INFO]  Core Release:\t", isCoreRelease ? 'yes' : 'no')
  console.log("[INFO]  Release Type:\t", releaseType)
  console.log("[INFO]  Should Release:\t", shouldRelease ? 'yes' : 'no')
  console.log("[INFO]  Previous tag:\t", lastTag || 'none')
  console.log("[INFO]  Next Release:\t", nextRelease)

  if (!shouldRelease) {
    console.log("[DEBUG]  Nothing to Release")
    return
  }

  console.log("[DEBUG] Generating changelog...")
  const releaseChangelog = await buildChangelog(commitsSinceLastTag, nextRelease)

  console.log(releaseChangelog)

  console.log("[DEBUG] Update CHANGELOG.md...")
  const changelog = readFileSync('./CHANGELOG.md').toString()
  writeFileSync('./CHANGELOG.md', releaseChangelog + '\n\n' + changelog)

  console.log("[DEBUG] Update package.json...")
  updateJsonFile('./package.json', pack => {
    pack.version = nextRelease
  })

  console.log("[DEBUG] Create release commit...")
  await createReleaseCommit(nextRelease, releaseChangelog, 'package.json', 'CHANGELOG.md')

  console.log("[DEBUG] Create release tag...")
  await createReleaseTag(nextRelease)

  console.log("[DEBUG] Push changes...")
  await pushChanges()

  console.log("[DEBUG] Publish GitHub release...")
  await createRelease(nextRelease, releaseChangelog)

  console.log("[DEBUG] Done.")

}

main().then(null)
