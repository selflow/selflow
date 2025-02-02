import { CreateNodesV2 } from 'nx/src/project-graph/plugins';
import { GO_MAIN_PACKAGE } from './constants';
import { generateAppProject, generateLibProject } from './projectConfigBuilder';
import { groupGoFileDirectories, parseGoFiles } from './goFileUtils';

export const createNodesV2: CreateNodesV2 = [
  '**/!(*_test).go',
  async (goFiles) => {
    console.error(goFiles);
    const dirs = groupGoFileDirectories([...goFiles]);
    const goProjects = await parseGoFiles(dirs);

    const goExecutableProjects = goProjects.filter(
      (project) => project.packageName === GO_MAIN_PACKAGE
    );
    const goLibraryProjects = goProjects.filter(
      (project) => project.packageName !== GO_MAIN_PACKAGE
    );

    const appProjects = goExecutableProjects.map(generateAppProject);
    const libProjects = goLibraryProjects.map(generateLibProject);

    return [...libProjects, ...appProjects];
  },
];
