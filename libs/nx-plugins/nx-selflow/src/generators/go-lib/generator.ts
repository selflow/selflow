import {
  addProjectConfiguration,
  formatFiles,
  generateFiles,
  names,
  offsetFromRoot,
  Tree,
} from '@nrwl/devkit';
import * as path from 'path';
import { NxPluginsNxSelflowGeneratorSchema } from './schema';
import { camelCase, snakeCase } from 'change-case';

interface NormalizedSchema extends NxPluginsNxSelflowGeneratorSchema {
  projectName: string;
  projectRoot: string;
  projectDirectory: string;
  parsedTags: string[];
  packageName: string;
  packageNameCamelCase: string;
  packageNameUpperCamelCase: string;
}

function normalizeOptions(
  tree: Tree,
  options: NxPluginsNxSelflowGeneratorSchema
): NormalizedSchema {
  const name = names(options.name).fileName;
  const projectDirectory = `${names(options.directory).fileName}/${name}`;
  const projectName = name.replace(new RegExp('/', 'g'), '-');
  const projectRoot = `${projectDirectory}`;
  const parsedTags = options.tags
    ? options.tags.split(',').map((s) => s.trim())
    : [];

  const packageNameCamelCase = camelCase(projectName);

  return {
    ...options,
    projectName,
    packageName: snakeCase(projectName),
    packageNameCamelCase,
    packageNameUpperCamelCase:
      packageNameCamelCase[0].toUpperCase() + packageNameCamelCase.slice(1),
    projectRoot,
    projectDirectory,
    parsedTags: [...parsedTags, ''],
  };
}

function addFiles(tree: Tree, options: NormalizedSchema) {
  const templateOptions = {
    ...options,
    ...names(options.name),
    offsetFromRoot: offsetFromRoot(options.projectRoot),
    template: '',
  };

  generateFiles(
    tree,
    path.join(__dirname, 'files'),
    options.projectRoot,
    templateOptions
  );
}

export default async function (
  tree: Tree,
  options: NxPluginsNxSelflowGeneratorSchema
) {
  const normalizedOptions = normalizeOptions(tree, options);

  addProjectConfiguration(tree, normalizedOptions.projectName, {
    root: normalizedOptions.projectRoot,
    projectType: 'library',
    sourceRoot: normalizedOptions.projectRoot,
    targets: {
      test: {
        executor: '@nx-go/nx-go:test',
      },
      lint: {
        executor: '@nx-go/nx-go:lint',
      },
    },
    tags: normalizedOptions.parsedTags,
  });
  addFiles(tree, normalizedOptions);

  await formatFiles(tree);
}
