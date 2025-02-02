import { CreateNodesV2 } from 'nx/src/project-graph/plugins';
import { readFile } from 'fs/promises';
import * as path from 'node:path';
import { CreateNodesResultV2 } from 'nx/src/project-graph/plugins/public-api';

const SELFLOW_MAIN_PACKAGE_NAME = 'github.com/selflow/selflow';
const GO_MAIN_PACKAGE = 'main';


function groupGoFileDirectories(files: string[]) {
  return files.reduce((acc, file) => {
    const dir = file.split('/').slice(0, -1).join('/');
    if (!acc.has(dir)) {
      acc.set(dir, []);
    }
    acc.get(dir)?.push(file);
    return acc;
  }, new Map<string, string[]>());
}

export const createNodesV2: CreateNodesV2 = [
  '**/!(*_test).go',
  async (goFiles) => {

    console.log(goFiles);

    const dirs = groupGoFileDirectories([...goFiles]);

    const goProjects = await Promise.all(
      [...dirs.entries()]
        .map(async ([dir, files]) => {

          files.sort();
          const firstFile = files[0];
          const fileContent = await readFile(firstFile);
          const packageName = RegExp(/package\s+(\w+)/).exec(fileContent.toString())?.[1];
          return {
            packageName,
            dir,
            sourceFile: firstFile
          };
        })
    );


    const goExecutableProjects = goProjects.filter(project => project.packageName === GO_MAIN_PACKAGE);

    const appProjects: CreateNodesResultV2 = goExecutableProjects.map(project => {
      const appSourceRoot = path.dirname(project.sourceFile);
      const projectName = project.packageName.replace(/\//g, '-');

      const main = SELFLOW_MAIN_PACKAGE_NAME + '/' + appSourceRoot;

      return ([
        project.sourceFile,
        {
          projects: {
            [appSourceRoot]: {
              name: projectName,
              sourceRoot: appSourceRoot,
              projectType: 'application',
              targets: {
                build: {
                  executor: '@nx-go/nx-go:build',
                  options: {
                    outputPath: path.join('dist', appSourceRoot),
                    main
                  },
                  outputs: ['{options.outputPath}']
                },
                serve: {
                  executor: '@nx-go/nx-go:serve',
                  options: {
                    main
                  }
                },
                test: {
                  executor: '@nx-go/nx-go:test'
                },
                lint: {
                  executor: '@nx-go/nx-go:lint'
                }
              },
              tags: ['lang:go']
            }
          }
        }
      ]);
    });

    const goProjectsWithoutMain = goProjects.filter(project => project.packageName !== GO_MAIN_PACKAGE);

    const libProjects: CreateNodesResultV2 = goProjectsWithoutMain.map(project => {
      const librarySourceRoot = path.dirname(project.sourceFile);
      const projectName = project.packageName.replace(/\//g, '-');

      return ([
        project.sourceFile,
        {
          projects: {
            [librarySourceRoot]: {
              name: projectName,
              sourceRoot: librarySourceRoot,
              projectType: 'library',
              targets: {
                test: {
                  executor: '@nx-go/nx-go:test'
                },
                lint: {
                  executor: '@nx-go/nx-go:lint'
                }
              },
              tags: ['lang:go']
            }
          }
        }
      ]);
    });

    return [...libProjects, ...appProjects];
  }
];
