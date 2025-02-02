// src/plugins/project-generators.ts
import { CreateNodesResultV2 } from 'nx/src/project-graph/plugins/public-api';
import * as path from 'node:path';
import { SELFLOW_MAIN_PACKAGE_NAME } from './constants';

import { GoProject } from './types';

const defaultGoProjectTargets = {
  test: { executor: '@nx-go/nx-go:test' },
  lint: { executor: '@nx-go/nx-go:lint' },
};

export function generateAppProject(
  project: GoProject
): CreateNodesResultV2[number] {
  const appSourceRoot = path.dirname(project.sourceFile);
  const projectName = project.packageName.replace(/\//g, '-');
  const main = `${SELFLOW_MAIN_PACKAGE_NAME}/${appSourceRoot}`;

  return [
    project.sourceFile,
    {
      projects: {
        [appSourceRoot]: {
          name: projectName,
          sourceRoot: appSourceRoot,
          projectType: 'application',
          targets: {
            ...defaultGoProjectTargets,
            build: {
              executor: '@nx-go/nx-go:build',
              options: {
                outputPath: path.join('dist', appSourceRoot),
                main,
              },
              outputs: ['{options.outputPath}'],
            },
            serve: {
              executor: '@nx-go/nx-go:serve',
              options: { main },
            },
          },
          tags: ['lang:go'],
        },
      },
    },
  ];
}

export function generateLibProject(
  project: GoProject
): CreateNodesResultV2[number] {
  const librarySourceRoot = path.dirname(project.sourceFile);
  const projectName = project.packageName.replace(/\//g, '-');

  return [
    project.sourceFile,
    {
      projects: {
        [librarySourceRoot]: {
          name: projectName,
          sourceRoot: librarySourceRoot,
          projectType: 'library',
          targets: {
            ...defaultGoProjectTargets,
          },
          tags: ['lang:go'],
        },
      },
    },
  ];
}
