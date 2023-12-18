import { Octokit } from 'octokit';
import semver from 'semver';
import { githubConfig } from './config';

export async function createRelease(tagName: string, changelog: string) {
  const octokit = new Octokit({ auth: githubConfig.token });

  const {
    data: { login },
  } = await octokit.rest.users.getAuthenticated();

  console.log(`[INFO]  Github - Logged in as ${login}`);

  const {
    data: { name: releaseName },
  } = await octokit.rest.repos.createRelease({
    name: tagName,
    body: changelog,
    prerelease: !!semver.prerelease(tagName),
    tag_name: tagName,

    owner: githubConfig.repoOwner,
    repo: githubConfig.repoName,
  });

  console.log(`[INFO]  Github - Created Release ${releaseName}`);
}
