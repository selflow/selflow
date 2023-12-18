import { getCommitsSinceTag } from './git';
import {
  createProjectGraphAsync,
  readProjectsConfigurationFromProjectGraph,
} from 'nx/src/project-graph/project-graph';
import { getAffectedGraphNodes } from 'nx/src/command-line/affected/affected';

export async function getAffectedProjects(
  commitHash: string
): Promise<string[]> {
  return getAffectedProjectsBetweenCommitHashes(commitHash, `${commitHash}^1`);
}

export async function getAffectedProjectsBetweenCommitHashes(
  headCommitHash: string,
  baseCommitHash: string
): Promise<string[]> {
  const projectGraph = await createProjectGraphAsync();

  const affectedGraph = await getAffectedGraphNodes(
    {
      base: baseCommitHash,
      head: headCommitHash,
    },
    projectGraph
  );

  return affectedGraph
    .filter((project) => project.type === 'app')
    .map((project) => project.name);
}

export async function getProjectDetails(project: string): Promise<string[]> {
  const projectGraph = await createProjectGraphAsync();
  const projects = readProjectsConfigurationFromProjectGraph(projectGraph);

  return projects.projects[project]?.tags ?? [];
}

/**
 * Return if a core release should be extracted from a commit list.
 *
 * A core release is a release that involves at least one NX project with the tag "scope:core"
 * @param commits
 */
export async function checkCoreRelease(
  commits: Awaited<ReturnType<typeof getCommitsSinceTag>>
): Promise<boolean> {
  const releaseAffectedProjects = await getAffectedProjectsBetweenCommitHashes(
    commits[commits.length - 1].hash,
    commits[0].hash
  );
  const projectDetails = await Promise.all(
    releaseAffectedProjects.map((p) => getProjectDetails(p))
  );

  return projectDetails.some((projectTags) =>
    projectTags.includes('scope:core')
  );
}
