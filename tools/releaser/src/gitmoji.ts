import { Commit, MainReleaseType } from './types';
import { getCommitsSinceTag } from './git';
import { config, releasePower } from './config';

const gitmojiMessageRegex = /^:[a-z0-9_]+:/;

export function parseCommit(
  commit: Omit<Commit, 'category' | 'messageToDisplay'>
): Commit {
  if (!gitmojiMessageRegex.test(commit.message)) {
    console.log('[WARN]  Invalid commit message', commit.message);
    return { ...commit };
  }

  const [, emoji, message] = commit.message.split(':', 3);

  return {
    ...commit,
    category: emoji,
    messageToDisplay: message,
  };
}

export async function getReleaseType(
  commits: Awaited<ReturnType<typeof getCommitsSinceTag>>
) {
  const invertedConfig = new Map<string, MainReleaseType>();
  for (const releaseType in config) {
    config[releaseType as MainReleaseType].forEach((emoji: string) =>
      invertedConfig.set(emoji, releaseType as MainReleaseType)
    );
  }

  let finalRelease: null | MainReleaseType = null;

  for (const commit of commits) {
    if (!gitmojiMessageRegex.test(commit.message)) {
      console.log('[WARN]  Invalid commit message', commit.message);
      continue;
    }
    const emoji = commit.message.split(':', 3)[1];
    const gitEmoji = `:${emoji}:`;

    if (invertedConfig.has(gitEmoji)) {
      const challengedNextReleaseType = invertedConfig.get(gitEmoji);

      if (!finalRelease) {
        finalRelease = challengedNextReleaseType;
        continue;
      }

      if (
        releasePower[challengedNextReleaseType] > releasePower[finalRelease]
      ) {
        finalRelease = challengedNextReleaseType;
      }
    }
  }

  return finalRelease;
}
